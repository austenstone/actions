package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"runtime"
	"time"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

type BuildInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

type Status struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Build     BuildInfo `json:"build"`
}

func main() {
	var (
		showVersion = flag.Bool("version", false, "Show version information")
		showStatus  = flag.Bool("status", false, "Show application status")
		outputJSON  = flag.Bool("json", false, "Output in JSON format")
		demo        = flag.Bool("demo", false, "Run demo command")
	)
	flag.Parse()

	if *showVersion {
		buildInfo := BuildInfo{
			Version:   version,
			Commit:    commit,
			BuildDate: date,
			GoVersion: runtime.Version(),
			Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		}

		if *outputJSON {
			output, _ := json.MarshalIndent(buildInfo, "", "  ")
			fmt.Println(string(output))
		} else {
			fmt.Printf("GitHub Actions Demo CLI Tool\n")
			fmt.Printf("Version:    %s\n", buildInfo.Version)
			fmt.Printf("Commit:     %s\n", buildInfo.Commit)
			fmt.Printf("Build Date: %s\n", buildInfo.BuildDate)
			fmt.Printf("Go Version: %s\n", buildInfo.GoVersion)
			fmt.Printf("Platform:   %s\n", buildInfo.Platform)
		}
		return
	}

	if *showStatus {
		status := Status{
			Status:    "running",
			Timestamp: time.Now(),
			Build: BuildInfo{
				Version:   version,
				Commit:    commit,
				BuildDate: date,
				GoVersion: runtime.Version(),
				Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			},
		}

		if *outputJSON {
			output, _ := json.MarshalIndent(status, "", "  ")
			fmt.Println(string(output))
		} else {
			fmt.Printf("Status: %s\n", status.Status)
			fmt.Printf("Timestamp: %s\n", status.Timestamp.Format(time.RFC3339))
			fmt.Printf("Version: %s\n", status.Build.Version)
			fmt.Printf("Platform: %s\n", status.Build.Platform)
		}
		return
	}

	if *demo {
		runDemo(*outputJSON)
		return
	}

	// Default behavior - show help
	fmt.Printf("GitHub Actions Demo CLI Tool v%s\n\n", version)
	fmt.Println("This tool demonstrates multi-platform release automation with GitHub Actions.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  cli-tool [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  cli-tool --version")
	fmt.Println("  cli-tool --status --json")
	fmt.Println("  cli-tool --demo")
}

func runDemo(jsonOutput bool) {
	features := []string{
		"Cross-platform compilation",
		"Automated releases",
		"GitHub CLI integration",
		"Docker multi-arch images",
		"Package signing",
		"SBOM generation",
	}

	platforms := []string{
		"linux/amd64",
		"linux/arm64",
		"darwin/amd64",
		"darwin/arm64",
		"windows/amd64",
	}

	demo := map[string]interface{}{
		"message":     "Welcome to the GitHub Actions Multi-Platform Release Demo! 🚀",
		"description": "This CLI tool showcases automated cross-platform builds and releases",
		"features":    features,
		"platforms":   platforms,
		"workflow":    "02-multi-platform-release.yml",
		"build_info": map[string]string{
			"version":    version,
			"commit":     commit,
			"platform":   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			"go_version": runtime.Version(),
		},
		"timestamp": time.Now().Format(time.RFC3339),
	}

	if jsonOutput {
		output, _ := json.MarshalIndent(demo, "", "  ")
		fmt.Println(string(output))
	} else {
		fmt.Println("🚀 GitHub Actions Multi-Platform Release Demo")
		fmt.Println("=" + repeat("=", 45))
		fmt.Println()
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Built with: %s\n", runtime.Version())
		fmt.Println()
		fmt.Println("✨ Features:")
		for _, feature := range features {
			fmt.Printf("  • %s\n", feature)
		}
		fmt.Println()
		fmt.Println("🖥️  Supported Platforms:")
		for _, platform := range platforms {
			fmt.Printf("  • %s\n", platform)
		}
		fmt.Println()
		fmt.Println("📋 Workflow: 02-multi-platform-release.yml")
		fmt.Println()
		fmt.Println("This tool demonstrates how GitHub Actions can:")
		fmt.Println("• Build binaries for multiple operating systems and architectures")
		fmt.Println("• Create automated releases with semantic versioning")
		fmt.Println("• Generate checksums and sign binaries")
		fmt.Println("• Build and push multi-architecture Docker images")
		fmt.Println("• Create Homebrew formulas and package distributions")
	}
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
