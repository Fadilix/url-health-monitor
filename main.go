package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	URLs                 []string `yaml:"urls"`
	TimeoutSeconds       int      `yaml:"timeout_seconds"`
	CheckIntervalSeconds int      `yaml:"check_interval_seconds"`
}

type URLStatus struct {
	URL          string
	StatusCode   int
	ResponseTime int64
	IsHealthy    bool
	LastChecked  time.Time
	Error        string
}

var (
	statuses = make(map[string]*URLStatus)

	mutex sync.RWMutex
)

func loadConfig(filename string) (Config, error) {
	var config Config

	data, err := os.ReadFile(filename)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func checkURL(url string, timeoutSeconds int) {
	client := http.Client{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(url)
	responseTime := time.Since(start).Milliseconds()

	status := &URLStatus{
		URL:          url,
		ResponseTime: responseTime,
		LastChecked:  time.Now(),
	}

	if err != nil {
		status.StatusCode = 0
		status.IsHealthy = false
		status.Error = err.Error()
	} else {
		defer resp.Body.Close()
		status.StatusCode = resp.StatusCode
		status.IsHealthy = resp.StatusCode >= 200 && resp.StatusCode < 400

		mutex.Lock()
		statuses[url] = status
		mutex.Unlock()
	}

	fmt.Printf("🟢 %s - Status : %d - Time : %d ms\n", url, resp.StatusCode, responseTime)
}

func startMonitoring(config Config) {
	ticker := time.NewTicker(time.Duration(config.CheckIntervalSeconds) * time.Second)
	defer ticker.Stop()

	fmt.Println("🚀 Lauching the monitoring...")

	for {
		fmt.Println("🔍 New verification")

		checkAllURLs(config)

		<-ticker.C
	}

}

func checkAllURLs(config Config) {
	var wg sync.WaitGroup

	for _, url := range config.URLs {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			checkURL(u, config.TimeoutSeconds)
		}(url)
	}

	wg.Wait()
}

func handleAPIStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	mutex.RLock()
	defer mutex.RUnlock()

	result := make([]URLStatus, 0, len(statuses))

	for _, status := range statuses {
		result = append(result, *status)
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	config, err := loadConfig("config.yaml")

	if err != nil {
		fmt.Println("Error while loading the configuration...")
		return
	}

	fmt.Printf("Config loaded successfully : %v\n", config)

	go startMonitoring(config)
	fmt.Println("🌐 Server listening on http://localhost:8080...")
	http.HandleFunc("/api/status", handleAPIStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
