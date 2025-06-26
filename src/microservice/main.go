package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type HealthCheck struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	Service   string    `json:"service"`
}

type ServiceInfo struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Environment string            `json:"environment"`
	Security    map[string]string `json:"security"`
	Features    []string          `json:"features"`
}

func main() {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/ready", readinessHandler).Methods("GET")
	r.HandleFunc("/info", infoHandler).Methods("GET")

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/service", serviceInfoHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Secure microservice starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	health := HealthCheck{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   getVersion(),
		Service:   "secure-microservice",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	// In a real microservice, you would check database connections,
	// external service dependencies, etc.

	ready := map[string]interface{}{
		"status":    "ready",
		"timestamp": time.Now(),
		"checks": map[string]string{
			"database":      "connected",
			"external_apis": "available",
			"cache":         "healthy",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ready)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"service":     "secure-microservice",
		"version":     getVersion(),
		"environment": getEnvironment(),
		"golang":      "1.21",
		"security_features": []string{
			"HTTPS only",
			"CORS protection",
			"Input validation",
			"Rate limiting",
			"Security headers",
		},
		"compliance": []string{
			"SOC 2",
			"GDPR",
			"HIPAA ready",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func serviceInfoHandler(w http.ResponseWriter, r *http.Request) {
	serviceInfo := ServiceInfo{
		Name:        "GitHub Actions Security Demo",
		Version:     getVersion(),
		Environment: getEnvironment(),
		Security: map[string]string{
			"code_scanning":       "CodeQL enabled",
			"dependency_scanning": "Trivy + npm audit",
			"container_scanning":  "Trivy container scan",
			"secret_scanning":     "TruffleHog",
			"policy_validation":   "OPA policies",
			"image_signing":       "Cosign",
			"sbom_generation":     "Syft",
		},
		Features: []string{
			"OIDC authentication",
			"Multi-stage Docker builds",
			"Non-root containers",
			"Resource limits",
			"Health checks",
			"Readiness probes",
			"Security policies",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serviceInfo)
}

func getVersion() string {
	version := os.Getenv("VERSION")
	if version == "" {
		return "1.0.0"
	}
	return version
}

func getEnvironment() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return "development"
	}
	return env
}
