
# URL Health Monitor 🚀

A real-time URL health monitoring system written in Go, capable of monitoring multiple URLs in parallel with a modern web interface.

## 🌟 Features

- ✅ **Parallel Monitoring**: Check multiple URLs simultaneously using goroutines
- ⚡ **Response Time Tracking**: Precise response time measurement for each URL
- 📊 **HTTP Status Codes**: Display status codes for each URL
- 🚨 **Notifications**: Alerts for downtime and recovery events
- 🌐 **Web Interface**: Modern and responsive interface to visualize results
- ⚙️ **YAML Configuration**: Simple configuration via YAML file
- 🔄 **REST API**: JSON endpoint for integration with other systems
- ⏱️ **Continuous Monitoring**: Automatic checks at regular intervals

## 🚀 Installation and Usage

### Prerequisites
- Go 1.21 or newer

### 1. Install Dependencies
```bash
go mod tidy
```

### 2. Configuration
Modify the `config.yaml` file according to your needs:

```yaml
urls:
  - https://google.com
  - https://github.com
  - https://example.com
  - https://your-website.com

timeout_seconds: 5
check_interval_seconds: 60
```

### 3. Running the Application
```bash
# Build
go build -o url-health-monitor.exe .

# Run
./url-health-monitor.exe
```

### 4. Access the Interface
Open your browser and navigate to: http://localhost:8080

## 📡 REST API

### GET /api/status
Returns the status of all URLs in JSON format:

```json
{
  "timestamp": "2025-06-26T10:30:00Z",
  "total": 3,
  "healthy": 2,
  "statuses": [
    {
      "url": "https://google.com",
      "status_code": 200,
      "response_time_ms": 45,
      "is_healthy": true,
      "last_check": "2025-06-26T10:30:00Z"
    }
  ]
}
```

### GET /health
Health endpoint for the service itself:

```json
{
  "status": "healthy",
  "timestamp": "2025-06-26T10:30:00Z",
  "service": "url-health-monitor"
}
```

## 🏗️ Architecture

### Go Concepts Used

1. **Goroutines**: Parallel URL checking for better performance
2. **Channels**: Communication between goroutines to collect results
3. **Context**: Proper handling of cancellation and timeouts
4. **Mutex**: Protection of shared data during concurrent access
5. **HTTP Client**: HTTP client configured with timeout
6. **HTML Templates**: Web interface rendering
7. **YAML Parsing**: Configuration file reading

### Code Structure

```
main.go
├── Config          # Configuration structure
├── URLStatus       # URL status structure
├── Monitor         # Main manager
│   ├── checkURL()           # Single URL check
│   ├── checkAllURLs()       # Parallel checking
│   ├── startMonitoring()    # Monitoring loop
└── HTTP Handlers    # Web and API endpoints
```

## 🔧 Advanced Configuration

### Customizing Timeouts
```yaml
timeout_seconds: 10        # Timeout for each HTTP request
check_interval_seconds: 30 # Interval between checks
```

### Adding URLs
Simply add new URLs to the list:
```yaml
urls:
  - https://my-website.com
  - https://api.my-service.com/health
  - https://cdn.example.com
```

## 📝 Logs

The program displays detailed logs:

```
2025/06/26 10:30:00 🚀 Starting URL monitoring...
2025/06/26 10:30:00 Configuration loaded successfully: 3 URLs to monitor
2025/06/26 10:30:01 🔄 Performing new health check
2025/06/26 10:30:01 🌐 Starting web server on http://localhost:8080...
2025/06/26 10:31:00 🔄 Performing new health check
```

## 🤝 Contributing

1. Fork the project
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request
