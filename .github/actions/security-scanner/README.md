# Security Scanner Action

A comprehensive Docker-based GitHub Action for security scanning of codebases.

## Features

- 🔐 **Secret Detection**: Finds hardcoded secrets, API keys, and sensitive data
- 📦 **Dependency Scanning**: Checks for vulnerable dependencies in Node.js, Go, and Python projects
- 🔍 **Code Analysis**: Identifies potentially unsafe code patterns
- 📊 **Multiple Output Formats**: JSON, SARIF, and Markdown table formats
- ⚙️ **Configurable Thresholds**: Set minimum severity levels for reporting
- 🎯 **Flexible Scanning**: Choose specific scan types or run comprehensive scans

## Usage

### Basic Usage

```yaml
- name: Security Scan
  uses: ./.github/actions/security-scanner
  with:
    target-path: './src'
```

### Advanced Configuration

```yaml
- name: Comprehensive Security Scan
  uses: ./.github/actions/security-scanner
  with:
    scan-type: 'all'
    target-path: '.'
    severity-threshold: 'high'
    fail-on-findings: 'true'
    output-format: 'sarif'
  id: security-scan

- name: Upload SARIF results
  uses: github/codeql-action/upload-sarif@v3
  if: always()
  with:
    sarif_file: ${{ steps.security-scan.outputs.scan-report }}
```

## Inputs

| Input | Description | Required | Default | Options |
|-------|-------------|----------|---------|---------|
| `scan-type` | Type of scan to perform | No | `all` | `secrets`, `dependencies`, `code`, `all` |
| `target-path` | Path to scan (relative to workspace) | No | `.` | Any valid path |
| `severity-threshold` | Minimum severity level to report | No | `medium` | `low`, `medium`, `high`, `critical` |
| `fail-on-findings` | Whether to fail the action if findings are discovered | No | `true` | `true`, `false` |
| `output-format` | Output format for results | No | `json` | `json`, `sarif`, `table` |

## Outputs

| Output | Description |
|--------|-------------|
| `findings-count` | Total number of security findings |
| `critical-count` | Number of critical findings |
| `high-count` | Number of high severity findings |
| `scan-report` | Path to the detailed scan report |

## Scan Types

### Secrets Scanning
Detects:
- Hardcoded passwords
- API keys and tokens
- Private keys
- Base64-encoded secrets
- Common secret patterns

### Dependency Scanning
Supports:
- **Node.js**: Uses `npm audit` to check package.json dependencies
- **Go**: Uses `govulncheck` for Go module vulnerabilities
- **Python**: Uses `safety` to check requirements.txt

### Code Analysis
Identifies:
- **JavaScript**: Dangerous patterns like `eval()`, `innerHTML`, `document.write()`
- **Go**: Potential SQL injection, command injection, file access issues
- Common security anti-patterns

## Examples

### Secrets-Only Scan

```yaml
- name: Scan for Secrets
  uses: ./.github/actions/security-scanner
  with:
    scan-type: 'secrets'
    target-path: './src'
    fail-on-findings: 'true'
```

### Dependency Vulnerability Check

```yaml
- name: Check Dependencies
  uses: ./.github/actions/security-scanner
  with:
    scan-type: 'dependencies'
    severity-threshold: 'high'
    output-format: 'json'
```

### Full Security Audit

```yaml
- name: Full Security Audit
  uses: ./.github/actions/security-scanner
  with:
    scan-type: 'all'
    target-path: '.'
    severity-threshold: 'medium'
    output-format: 'sarif'
    fail-on-findings: 'true'
```

### Matrix Scanning

```yaml
jobs:
  security-scan:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        scan-type: [secrets, dependencies, code]
    
    steps:
      - uses: actions/checkout@v4
      
      - name: Security Scan - ${{ matrix.scan-type }}
        uses: ./.github/actions/security-scanner
        with:
          scan-type: ${{ matrix.scan-type }}
          target-path: '.'
          fail-on-findings: 'false'
```

## Output Formats

### JSON Format
Structured JSON report with detailed findings:
```json
{
  "scan_metadata": {
    "timestamp": "2024-01-01T12:00:00Z",
    "scan_type": "all",
    "target_path": "."
  },
  "findings": [
    {
      "severity": "high",
      "category": "secrets",
      "description": "Potential API key detected",
      "file_path": "./config.js",
      "line_number": "15"
    }
  ],
  "summary": {
    "total": 1,
    "critical": 0,
    "high": 1,
    "medium": 0,
    "low": 0
  }
}
```

### SARIF Format
Standard SARIF format for integration with GitHub Security tab:
```json
{
  "version": "2.1.0",
  "runs": [{
    "tool": {
      "driver": {
        "name": "Security Scanner",
        "version": "1.0.0"
      }
    },
    "results": [...]
  }]
}
```

### Table Format
Markdown table for easy reading:
```markdown
| Severity | Category | File | Description |
|----------|----------|------|-------------|
| high | secrets | ./config.js | Potential API key detected |
```

## Integration with GitHub Security

### Upload SARIF Results

```yaml
- name: Upload SARIF
  uses: github/codeql-action/upload-sarif@v3
  if: always()
  with:
    sarif_file: ${{ steps.security-scan.outputs.scan-report }}
```

### Create Issues for Findings

```yaml
- name: Create Security Issue
  if: steps.security-scan.outputs.critical-count > 0
  uses: actions/github-script@v7
  with:
    script: |
      github.rest.issues.create({
        owner: context.repo.owner,
        repo: context.repo.repo,
        title: '🚨 Critical Security Findings Detected',
        body: 'Critical security vulnerabilities found in the latest scan.',
        labels: ['security', 'bug']
      });
```

## Limitations

- This is a demo action with simplified scanning logic
- For production use, consider dedicated security tools like:
  - GitHub Advanced Security
  - Snyk
  - OWASP ZAP
  - Semgrep
  - Bandit

## License

MIT License. See [LICENSE](../../../LICENSE) for details.
