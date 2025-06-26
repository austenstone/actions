#!/bin/bash

set -e

# Input parameters
SCAN_TYPE="${1:-all}"
TARGET_PATH="${2:-.}"
SEVERITY_THRESHOLD="${3:-medium}"
FAIL_ON_FINDINGS="${4:-true}"
OUTPUT_FORMAT="${5:-json}"

echo "🔍 Starting security scan..."
echo "Scan type: $SCAN_TYPE"
echo "Target path: $TARGET_PATH"
echo "Severity threshold: $SEVERITY_THRESHOLD"
echo "Output format: $OUTPUT_FORMAT"
echo "Fail on findings: $FAIL_ON_FINDINGS"

# Initialize counters
TOTAL_FINDINGS=0
CRITICAL_COUNT=0
HIGH_COUNT=0
MEDIUM_COUNT=0
LOW_COUNT=0

# Create output directory
mkdir -p /tmp/scan-results
REPORT_FILE="/tmp/scan-results/security-report.json"

# Initialize report
cat > "$REPORT_FILE" << EOF
{
  "scan_metadata": {
    "timestamp": "$(date -Iseconds)",
    "scan_type": "$SCAN_TYPE",
    "target_path": "$TARGET_PATH",
    "severity_threshold": "$SEVERITY_THRESHOLD"
  },
  "findings": [],
  "summary": {
    "total": 0,
    "critical": 0,
    "high": 0,
    "medium": 0,
    "low": 0
  }
}
EOF

# Function to add finding to report
add_finding() {
    local severity=$1
    local category=$2
    local description=$3
    local file_path=$4
    local line_number=${5:-""}
    
    # Update counters
    case $severity in
        "critical") ((CRITICAL_COUNT++)) ;;
        "high") ((HIGH_COUNT++)) ;;
        "medium") ((MEDIUM_COUNT++)) ;;
        "low") ((LOW_COUNT++)) ;;
    esac
    ((TOTAL_FINDINGS++))
    
    # Add to JSON report
    local finding=$(cat << EOF
{
  "severity": "$severity",
  "category": "$category",
  "description": "$description",
  "file_path": "$file_path",
  "line_number": "$line_number",
  "timestamp": "$(date -Iseconds)"
}
EOF
    )
    
    # Use jq to add the finding to the report
    local temp_file=$(mktemp)
    jq --argjson finding "$finding" '.findings += [$finding]' "$REPORT_FILE" > "$temp_file" && mv "$temp_file" "$REPORT_FILE"
}

# Function to scan for secrets
scan_secrets() {
    echo "🔐 Scanning for secrets and sensitive data..."
    
    # Simple regex-based secret detection
    local secret_patterns=(
        "password\s*=\s*['\"][^'\"]*['\"]"
        "api[_-]?key\s*=\s*['\"][^'\"]*['\"]"
        "secret\s*=\s*['\"][^'\"]*['\"]"
        "token\s*=\s*['\"][^'\"]*['\"]"
        "['\"][A-Za-z0-9+/]{40,}['\"]"  # Base64-like strings
        "-----BEGIN.*PRIVATE KEY-----"
    )
    
    for pattern in "${secret_patterns[@]}"; do
        while IFS= read -r line; do
            if [[ -n "$line" ]]; then
                local file_path=$(echo "$line" | cut -d: -f1)
                local line_num=$(echo "$line" | cut -d: -f2)
                local content=$(echo "$line" | cut -d: -f3-)
                
                add_finding "high" "secrets" "Potential secret or sensitive data detected: $content" "$file_path" "$line_num"
            fi
        done < <(grep -rn -E "$pattern" "$TARGET_PATH" 2>/dev/null || true)
    done
}

# Function to scan dependencies
scan_dependencies() {
    echo "📦 Scanning dependencies for vulnerabilities..."
    
    # Scan Node.js dependencies
    if [[ -f "$TARGET_PATH/package.json" ]]; then
        echo "Scanning Node.js dependencies..."
        cd "$TARGET_PATH"
        
        # Use npm audit
        if npm audit --json > /tmp/npm-audit.json 2>/dev/null; then
            local vulnerabilities=$(jq -r '.vulnerabilities | to_entries[] | select(.value.severity == "high" or .value.severity == "critical") | "\(.key):\(.value.severity):\(.value.title)"' /tmp/npm-audit.json 2>/dev/null || echo "")
            
            while IFS=: read -r package severity title; do
                if [[ -n "$package" ]]; then
                    add_finding "$severity" "dependency" "Vulnerable dependency: $package - $title" "package.json" ""
                fi
            done <<< "$vulnerabilities"
        fi
        
        cd - > /dev/null
    fi
    
    # Scan Go dependencies
    if [[ -f "$TARGET_PATH/go.mod" ]]; then
        echo "Scanning Go dependencies..."
        cd "$TARGET_PATH"
        
        # Use govulncheck if available
        if command -v govulncheck &> /dev/null; then
            if govulncheck -json ./... > /tmp/go-vuln.json 2>/dev/null; then
                # Parse govulncheck output (simplified)
                local vuln_count=$(jq -r '.Vulns // [] | length' /tmp/go-vuln.json 2>/dev/null || echo "0")
                if [[ "$vuln_count" -gt 0 ]]; then
                    add_finding "high" "dependency" "Go vulnerabilities detected: $vuln_count issues found" "go.mod" ""
                fi
            fi
        fi
        
        cd - > /dev/null
    fi
    
    # Scan Python dependencies
    if [[ -f "$TARGET_PATH/requirements.txt" ]]; then
        echo "Scanning Python dependencies..."
        cd "$TARGET_PATH"
        
        # Use safety
        if safety check --json > /tmp/safety.json 2>/dev/null; then
            local vulnerabilities=$(jq -r '.[] | "\(.vulnerability_id):\(.advisory):\(.package_name)"' /tmp/safety.json 2>/dev/null || echo "")
            
            while IFS=: read -r vuln_id advisory package; do
                if [[ -n "$vuln_id" ]]; then
                    add_finding "high" "dependency" "Vulnerable Python package: $package - $advisory" "requirements.txt" ""
                fi
            done <<< "$vulnerabilities"
        fi
        
        cd - > /dev/null
    fi
}

# Function to scan code for vulnerabilities
scan_code() {
    echo "🔍 Scanning code for security vulnerabilities..."
    
    # Scan JavaScript/Node.js files
    find "$TARGET_PATH" -name "*.js" -not -path "*/node_modules/*" | while read -r file; do
        # Simple vulnerability patterns for demo
        local patterns=(
            "eval\s*\(" 
            "innerHTML\s*="
            "document\.write\s*\("
            "\.exec\s*\("
        )
        
        for pattern in "${patterns[@]}"; do
            while IFS=: read -r line_num content; do
                if [[ -n "$content" ]]; then
                    add_finding "medium" "code-vulnerability" "Potentially unsafe JavaScript pattern: $pattern" "$file" "$line_num"
                fi
            done < <(grep -n -E "$pattern" "$file" 2>/dev/null || true)
        done
    done
    
    # Scan Go files
    find "$TARGET_PATH" -name "*.go" | while read -r file; do
        # Simple Go vulnerability patterns
        local patterns=(
            "sql\.Query\s*\("
            "exec\.Command\s*\("
            "os\.OpenFile\s*\("
        )
        
        for pattern in "${patterns[@]}"; do
            while IFS=: read -r line_num content; do
                if [[ -n "$content" ]]; then
                    add_finding "medium" "code-vulnerability" "Potentially unsafe Go pattern: $pattern" "$file" "$line_num"
                fi
            done < <(grep -n -E "$pattern" "$file" 2>/dev/null || true)
        done
    done
}

# Function to update report summary
update_summary() {
    local temp_file=$(mktemp)
    jq --arg total "$TOTAL_FINDINGS" \
       --arg critical "$CRITICAL_COUNT" \
       --arg high "$HIGH_COUNT" \
       --arg medium "$MEDIUM_COUNT" \
       --arg low "$LOW_COUNT" \
       '.summary.total = ($total | tonumber) |
        .summary.critical = ($critical | tonumber) |
        .summary.high = ($high | tonumber) |
        .summary.medium = ($medium | tonumber) |
        .summary.low = ($low | tonumber)' \
       "$REPORT_FILE" > "$temp_file" && mv "$temp_file" "$REPORT_FILE"
}

# Run scans based on scan type
case $SCAN_TYPE in
    "secrets")
        scan_secrets
        ;;
    "dependencies")
        scan_dependencies
        ;;
    "code")
        scan_code
        ;;
    "all")
        scan_secrets
        scan_dependencies
        scan_code
        ;;
    *)
        echo "❌ Invalid scan type: $SCAN_TYPE"
        exit 1
        ;;
esac

# Update summary
update_summary

# Output results
echo ""
echo "📊 Security Scan Results:"
echo "========================="
echo "Total findings: $TOTAL_FINDINGS"
echo "Critical: $CRITICAL_COUNT"
echo "High: $HIGH_COUNT"  
echo "Medium: $MEDIUM_COUNT"
echo "Low: $LOW_COUNT"

# Set GitHub Actions outputs
echo "findings-count=$TOTAL_FINDINGS" >> $GITHUB_OUTPUT
echo "critical-count=$CRITICAL_COUNT" >> $GITHUB_OUTPUT
echo "high-count=$HIGH_COUNT" >> $GITHUB_OUTPUT
echo "scan-report=$REPORT_FILE" >> $GITHUB_OUTPUT

# Generate different output formats
case $OUTPUT_FORMAT in
    "json")
        echo "📄 JSON Report generated: $REPORT_FILE"
        ;;
    "sarif")
        echo "📄 Converting to SARIF format..."
        # Convert JSON to SARIF (simplified)
        cat > "${REPORT_FILE%.*}.sarif" << EOF
{
  "version": "2.1.0",
  "runs": [{
    "tool": {
      "driver": {
        "name": "Security Scanner",
        "version": "1.0.0"
      }
    },
    "results": $(jq '[.findings[] | {
      "ruleId": .category,
      "message": {"text": .description},
      "level": (if .severity == "critical" then "error" elif .severity == "high" then "error" elif .severity == "medium" then "warning" else "note" end),
      "locations": [{
        "physicalLocation": {
          "artifactLocation": {"uri": .file_path},
          "region": {"startLine": (.line_number | tonumber // 1)}
        }
      }]
    }]' "$REPORT_FILE")
  }]
}
EOF
        echo "scan-report=${REPORT_FILE%.*}.sarif" >> $GITHUB_OUTPUT
        ;;
    "table")
        echo "📄 Generating table format..."
        echo "| Severity | Category | File | Description |" > "${REPORT_FILE%.*}.md"
        echo "|----------|----------|------|-------------|" >> "${REPORT_FILE%.*}.md"
        jq -r '.findings[] | "| \(.severity) | \(.category) | \(.file_path) | \(.description) |"' "$REPORT_FILE" >> "${REPORT_FILE%.*}.md"
        echo "scan-report=${REPORT_FILE%.*}.md" >> $GITHUB_OUTPUT
        ;;
esac

# Create GitHub summary
{
    echo "## 🔒 Security Scan Results"
    echo ""
    echo "| Severity | Count |"
    echo "|----------|-------|"
    echo "| Critical | $CRITICAL_COUNT |"
    echo "| High | $HIGH_COUNT |"
    echo "| Medium | $MEDIUM_COUNT |" 
    echo "| Low | $LOW_COUNT |"
    echo "| **Total** | **$TOTAL_FINDINGS** |"
    echo ""
    
    if [[ $TOTAL_FINDINGS -gt 0 ]]; then
        echo "### 🚨 Top Findings"
        echo ""
        jq -r '.findings[:5] | .[] | "- **\(.severity | ascii_upcase)** (\(.category)): \(.description) in `\(.file_path)`"' "$REPORT_FILE"
    else
        echo "### ✅ No security issues found!"
    fi
} >> $GITHUB_STEP_SUMMARY

# Determine if we should fail
should_fail=false

case $SEVERITY_THRESHOLD in
    "low")
        [[ $TOTAL_FINDINGS -gt 0 ]] && should_fail=true
        ;;
    "medium")
        [[ $MEDIUM_COUNT -gt 0 || $HIGH_COUNT -gt 0 || $CRITICAL_COUNT -gt 0 ]] && should_fail=true
        ;;
    "high")
        [[ $HIGH_COUNT -gt 0 || $CRITICAL_COUNT -gt 0 ]] && should_fail=true
        ;;
    "critical")
        [[ $CRITICAL_COUNT -gt 0 ]] && should_fail=true
        ;;
esac

# Exit with error if configured to fail on findings
if [[ "$FAIL_ON_FINDINGS" == "true" && "$should_fail" == true ]]; then
    echo ""
    echo "❌ Security scan failed due to findings above threshold ($SEVERITY_THRESHOLD)"
    exit 1
fi

echo ""
echo "✅ Security scan completed successfully"
exit 0
