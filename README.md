# 🏥 Healthcheck REST API (Go) - **Intermediate Level**

A lightweight monitoring service exposing health metrics and runtime information through REST endpoints. Demonstrates HTTP server development, runtime introspection, and graceful shutdown patterns in Go.

---

## 🚀 What is this?

A production-ready health monitoring API that exposes system metrics in both JSON and Prometheus formats. Perfect for microservices, containerized environments, and demonstrating Go's runtime monitoring capabilities.

---

## ✨ Features

- **Health Endpoint:** JSON health reports with uptime and memory stats
- **Metrics Endpoint:** Prometheus-style metrics for monitoring systems
- **Runtime Introspection:** Memory usage, garbage collection, and goroutine counts
- **Graceful Shutdown:** Proper signal handling and resource cleanup
- **Zero Dependencies:** Pure Go standard library implementation

---

## 🦄 Go Skills Demonstrated

- **HTTP Servers:** `net/http` server setup and routing
- **Runtime Package:** Memory stats and goroutine introspection
- **Signal Handling:** Graceful shutdown with `os/signal`
- **JSON Serialization:** Custom health data structures
- **Goroutines:** Concurrent server and signal handling
- **Time Operations:** Uptime calculation and formatting

---

## 🛠️ Usage

```sh
# Start the health server
go run healthcheck.go

# Check health status (JSON)
curl http://localhost:8080/health

# Get Prometheus metrics
curl http://localhost:8080/metrics
```

**Example JSON Response:**
```json
{
  "uptime": "123s",
  "mem_alloc": "1.23MiB", 
  "tot_alloc": "2.34MiB",
  "sys_mem": "4.56MiB",
  "num_gc": "5"
}
```

---

## 🎯 Learning Objectives

This project demonstrates:
- **HTTP Server Development:** Building REST APIs with standard library
- **System Monitoring:** Runtime introspection and health metrics
- **Concurrent Programming:** Server and signal handling with goroutines
- **Production Patterns:** Graceful shutdown and monitoring endpoints

**Difficulty:** ⭐⭐⭐ Intermediate - HTTP servers and system monitoring

---

**Author:** IAmSotiris
