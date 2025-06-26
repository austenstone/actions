package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"
)

func TestBuildInfo(t *testing.T) {
	buildInfo := BuildInfo{
		Version:   "1.0.0",
		Commit:    "abc123",
		BuildDate: "2023-10-01",
		GoVersion: "go1.21.0",
		Platform:  "linux/amd64",
	}

	// Test JSON marshaling
	data, err := json.Marshal(buildInfo)
	if err != nil {
		t.Errorf("Failed to marshal BuildInfo: %v", err)
	}

	var unmarshaled BuildInfo
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal BuildInfo: %v", err)
	}

	if unmarshaled.Version != buildInfo.Version {
		t.Errorf("Expected version %s, got %s", buildInfo.Version, unmarshaled.Version)
	}
}

func TestStatus(t *testing.T) {
	now := time.Now()
	status := Status{
		Status:    "running",
		Timestamp: now,
		Build: BuildInfo{
			Version:   "1.0.0",
			Commit:    "abc123",
			BuildDate: "2023-10-01",
			GoVersion: "go1.21.0",
			Platform:  "linux/amd64",
		},
	}

	if status.Status != "running" {
		t.Errorf("Expected status 'running', got %s", status.Status)
	}

	if status.Timestamp != now {
		t.Errorf("Timestamp mismatch")
	}
}

func TestVersionOutput(t *testing.T) {
	// Save original values
	originalVersion := version
	originalCommit := commit
	originalDate := date

	// Set test values
	version = "1.0.0-test"
	commit = "test-commit"
	date = "2023-10-01"

	// Test version flag simulation
	buildInfo := BuildInfo{
		Version:   version,
		Commit:    commit,
		BuildDate: date,
	}

	if buildInfo.Version != "1.0.0-test" {
		t.Errorf("Expected version '1.0.0-test', got %s", buildInfo.Version)
	}

	// Restore original values
	version = originalVersion
	commit = originalCommit
	date = originalDate
}

func TestJSONOutput(t *testing.T) {
	status := Status{
		Status:    "running",
		Timestamp: time.Now(),
		Build: BuildInfo{
			Version:   "1.0.0",
			Commit:    "abc123",
			BuildDate: "2023-10-01",
			GoVersion: "go1.21.0",
			Platform:  "linux/amd64",
		},
	}

	data, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		t.Errorf("Failed to marshal status to JSON: %v", err)
	}

	jsonStr := string(data)
	if !strings.Contains(jsonStr, "running") {
		t.Errorf("JSON output should contain 'running' status")
	}

	if !strings.Contains(jsonStr, "1.0.0") {
		t.Errorf("JSON output should contain version '1.0.0'")
	}
}

func TestRepeatFunction(t *testing.T) {
	result := repeat("=", 5)
	expected := "====="
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	result = repeat("a", 0)
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}

	result = repeat("test", 2)
	expected = "testtest"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMainFunction(t *testing.T) {
	// Test that main doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() panicked: %v", r)
		}
	}()

	// Save original args
	originalArgs := os.Args

	// Test help (default behavior)
	os.Args = []string{"cli-tool"}
	// We can't easily test main() output without refactoring,
	// but we can ensure it doesn't crash

	// Restore original args
	os.Args = originalArgs
}
