package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type HealthStats struct {
	Uptime   string `json:"uptime"`
	MemAlloc string `json:"mem_alloc"`
	TotAlloc string `json:"tot_alloc"`
	SysMem   string `json:"sys_mem"`
	NumGC    string `json:"num_gc"`
}

var startTime time.Time

func main() {
	port := 8080
	mux := http.NewServeMux()
	startTime = time.Now()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/metrics", metricsHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	go func() {
		fmt.Println("🚀 Healthcheck API running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()

	<-stop
	fmt.Println("\nExiting server")
}

const memFormat = "%.2vMiB"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	health := &HealthStats{
		MemAlloc: fmt.Sprintf(memFormat, float64(mem.Alloc)/1024/1024),
		TotAlloc: fmt.Sprintf(memFormat, float64(mem.TotalAlloc)/1024/1024),
		SysMem:   fmt.Sprintf(memFormat, float64(mem.Sys)/1024/1024),
		NumGC:    fmt.Sprintf("%v", runtime.NumGoroutine()),
		Uptime:   fmt.Sprintf("%ds", int(time.Since(startTime).Seconds())),
	}

	jsonData, err := json.Marshal(*health)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v\n", err.Error()), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Fprintf(w, "app_uptime_seconds %.0f\n", uptime.Seconds())
	fmt.Fprintf(w, "go_goroutines %d\n", runtime.NumGoroutine())
	fmt.Fprintf(w, "go_mem_alloc_bytes %d\n", m.Alloc)
	fmt.Fprintf(w, "go_mem_total_alloc_bytes %d\n", m.TotalAlloc)
	fmt.Fprintf(w, "go_mem_sys_bytes %d\n", m.Sys)
}
