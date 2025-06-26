# Health Check Action

A comprehensive GitHub Action for performing health checks on web applications and APIs.

## Features

- 🏥 HTTP/HTTPS health checks with configurable parameters
- ⏱️ Configurable timeout and retry logic with exponential backoff
- 📊 Detailed response time measurements
- 🎯 Flexible status code validation
- 📋 Comprehensive reporting and summaries
- 🔄 Built-in retry mechanism for reliability

## Usage

### Basic Usage

```yaml
- name: Health Check
  uses: ./.github/actions/health-check
  with:
    url: 'https://your-app.com/health'
```

### Advanced Usage

```yaml
- name: Comprehensive Health Check
  uses: ./.github/actions/health-check
  with:
    url: 'https://your-app.com/api/health'
    timeout: '10'
    retry-count: '5'
    expected-status: '200'
  id: health-check

- name: Check Results
  run: |
    echo "Status: ${{ steps.health-check.outputs.status }}"
    echo "Response Time: ${{ steps.health-check.outputs.response-time }}ms"
    echo "Details: ${{ steps.health-check.outputs.details }}"
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `url` | URL to perform health check on | Yes | - |
| `timeout` | Timeout in seconds | No | `30` |
| `retry-count` | Number of retry attempts | No | `3` |
| `expected-status` | Expected HTTP status code | No | `200` |

## Outputs

| Output | Description |
|--------|-------------|
| `status` | Health check status (`healthy` or `unhealthy`) |
| `response-time` | Response time in milliseconds |
| `details` | JSON string with detailed results |

## Examples

### Web Application Health Check

```yaml
name: Web App Health Check
on:
  deployment_status:
    types: [success]

jobs:
  health-check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Wait for deployment
        run: sleep 30
      
      - name: Health Check
        uses: ./.github/actions/health-check
        with:
          url: ${{ github.event.deployment.payload.web_url }}/health
          timeout: '15'
          retry-count: '5'
```

### API Endpoint Validation

```yaml
- name: API Health Check
  uses: ./.github/actions/health-check
  with:
    url: 'https://api.example.com/v1/status'
    expected-status: '200'
    timeout: '5'
```

### Multiple Endpoint Checks

```yaml
jobs:
  health-checks:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        endpoint:
          - { url: 'https://app.com/health', name: 'main-app' }
          - { url: 'https://api.app.com/health', name: 'api' }
          - { url: 'https://admin.app.com/health', name: 'admin' }
    
    steps:
      - uses: actions/checkout@v4
      
      - name: Health Check - ${{ matrix.endpoint.name }}
        uses: ./.github/actions/health-check
        with:
          url: ${{ matrix.endpoint.url }}
          timeout: '10'
          retry-count: '3'
```

## Development

### Building

```bash
npm install
npm run build
```

### Testing

```bash
npm test
```

### Linting

```bash
npm run lint
```

## License

MIT License. See [LICENSE](../../../LICENSE) for details.
