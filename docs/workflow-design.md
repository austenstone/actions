# Workflow Design Guide

This guide provides comprehensive guidance for designing effective GitHub Actions workflows.

## 📋 Table of Contents

- [Design Principles](#design-principles)
- [Workflow Structure](#workflow-structure)
- [Job Design](#job-design)
- [Advanced Patterns](#advanced-patterns)
- [Testing Strategies](#testing-strategies)
- [Deployment Patterns](#deployment-patterns)
- [Monitoring and Observability](#monitoring-and-observability)
- [Examples](#examples)

## 🎯 Design Principles

### 1. Single Responsibility
Each workflow should have a clear, single purpose.

```yaml
# ✅ Good - Single purpose
name: Test Suite
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: npm test

# ❌ Avoid - Multiple unrelated purposes
name: Everything
jobs:
  test-and-deploy-and-notify:
    # Too many responsibilities
```

### 2. Fail Fast
Design workflows to fail quickly when issues are detected.

```yaml
jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Validate syntax
        run: yamllint .github/workflows/
      - name: Lint code
        run: npm run lint
  
  test:
    needs: validate  # Only run if validation passes
    runs-on: ubuntu-latest
    steps:
      - run: npm test
```

### 3. Modularity
Break complex workflows into smaller, reusable components.

```yaml
# Reusable workflow
name: Build and Test
on:
  workflow_call:
    inputs:
      node-version:
        required: true
        type: string

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: ${{ inputs.node-version }}
      - run: npm ci
      - run: npm test
```

### 4. Idempotency
Workflows should produce the same result when run multiple times.

```yaml
# ✅ Idempotent deployment
- name: Deploy
  run: |
    if ! kubectl get deployment app-deployment; then
      kubectl create deployment app-deployment --image=app:${{ github.sha }}
    else
      kubectl set image deployment/app-deployment app=app:${{ github.sha }}
    fi
```

## 🏗️ Workflow Structure

### Basic Structure
```yaml
name: Descriptive Workflow Name
on:
  # Triggers
env:
  # Global environment variables
jobs:
  job-name:
    name: Human Readable Job Name
    runs-on: ubuntu-latest
    env:
      # Job-specific environment variables
    outputs:
      # Job outputs
    steps:
      # Job steps
```

### Comprehensive Example
```yaml
name: 🚀 Production Deployment Pipeline

on:
  push:
    branches: [main]
    tags: ['v*']
  workflow_dispatch:
    inputs:
      environment:
        description: 'Deployment environment'
        required: true
        default: 'staging'
        type: choice
        options: [staging, production]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  # Pre-flight checks
  preflight:
    name: 🔍 Pre-flight Checks
    runs-on: ubuntu-latest
    outputs:
      version: ${{ steps.version.outputs.version }}
      deploy-env: ${{ steps.env.outputs.environment }}
    
    steps:
      - name: 📥 Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
      
      - name: 🏷️ Generate Version
        id: version
        run: |
          if [[ ${{ github.ref }} == refs/tags/* ]]; then
            VERSION=${GITHUB_REF#refs/tags/}
          else
            VERSION=$(date +%Y%m%d)-${GITHUB_SHA::8}
          fi
          echo "version=$VERSION" >> $GITHUB_OUTPUT
      
      - name: 🎯 Determine Environment
        id: env
        run: |
          if [[ "${{ github.event_name }}" == "workflow_dispatch" ]]; then
            ENV="${{ github.event.inputs.environment }}"
          elif [[ ${{ github.ref }} == refs/tags/* ]]; then
            ENV="production"
          else
            ENV="staging"
          fi
          echo "environment=$ENV" >> $GITHUB_OUTPUT

  # Build and test
  build:
    name: 🏗️ Build & Test
    runs-on: ubuntu-latest
    needs: preflight
    
    steps:
      - name: 📥 Checkout
        uses: actions/checkout@v4
      
      - name: 🏗️ Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '18'
          cache: 'npm'
      
      - name: 📦 Install Dependencies
        run: npm ci
      
      - name: 🧪 Run Tests
        run: npm test
      
      - name: 🏗️ Build Application
        run: npm run build
      
      - name: 📊 Upload Artifacts
        uses: actions/upload-artifact@v4
        with:
          name: build-${{ needs.preflight.outputs.version }}
          path: dist/

  # Deploy
  deploy:
    name: 🚀 Deploy to ${{ needs.preflight.outputs.deploy-env }}
    runs-on: ubuntu-latest
    needs: [preflight, build]
    environment:
      name: ${{ needs.preflight.outputs.deploy-env }}
      url: ${{ steps.deploy.outputs.url }}
    
    steps:
      - name: 📥 Download Artifacts
        uses: actions/download-artifact@v4
        with:
          name: build-${{ needs.preflight.outputs.version }}
      
      - name: 🚀 Deploy
        id: deploy
        run: |
          echo "Deploying version ${{ needs.preflight.outputs.version }} to ${{ needs.preflight.outputs.deploy-env }}"
          # Deployment logic here
          echo "url=https://${{ needs.preflight.outputs.deploy-env }}.example.com" >> $GITHUB_OUTPUT
```

## 💼 Job Design

### Job Dependencies
```yaml
jobs:
  lint:
    runs-on: ubuntu-latest
    steps: [...]
  
  test:
    runs-on: ubuntu-latest
    steps: [...]
  
  security:
    runs-on: ubuntu-latest
    steps: [...]
  
  build:
    needs: [lint, test, security]  # Wait for all checks
    runs-on: ubuntu-latest
    steps: [...]
  
  deploy:
    needs: build
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps: [...]
```

### Matrix Strategies
```yaml
jobs:
  test:
    runs-on: ${{ matrix.os }}
    strategy:
      fail-fast: false
      matrix:
        os: [ubuntu-latest, windows-latest, macos-latest]
        node: [16, 18, 20]
        include:
          - os: ubuntu-latest
            node: 20
            experimental: true
        exclude:
          - os: windows-latest
            node: 16
    
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: ${{ matrix.node }}
      - run: npm test
```

### Conditional Execution
```yaml
jobs:
  deploy:
    if: |
      github.event_name == 'push' && 
      github.ref == 'refs/heads/main' && 
      !contains(github.event.head_commit.message, '[skip deploy]')
    
    steps:
      - name: Deploy to Production
        if: github.ref == 'refs/heads/main'
        run: echo "Deploying to production"
      
      - name: Deploy to Staging
        if: github.ref != 'refs/heads/main'
        run: echo "Deploying to staging"
```

## 🔄 Advanced Patterns

### Reusable Workflows
```yaml
# .github/workflows/reusable-build.yml
name: Reusable Build
on:
  workflow_call:
    inputs:
      node-version:
        required: true
        type: string
      environment:
        required: false
        type: string
        default: 'development'
    outputs:
      build-version:
        description: "The build version"
        value: ${{ jobs.build.outputs.version }}
    secrets:
      NPM_TOKEN:
        required: true

jobs:
  build:
    runs-on: ubuntu-latest
    outputs:
      version: ${{ steps.version.outputs.version }}
    
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: ${{ inputs.node-version }}
      - run: npm ci
        env:
          NPM_TOKEN: ${{ secrets.NPM_TOKEN }}
      - run: npm run build:${{ inputs.environment }}
```

### Dynamic Matrix Generation
```yaml
jobs:
  generate-matrix:
    runs-on: ubuntu-latest
    outputs:
      matrix: ${{ steps.generate.outputs.matrix }}
    
    steps:
      - uses: actions/checkout@v4
      - id: generate
        run: |
          # Generate matrix based on changed files
          MATRIX=$(find src/ -name "package.json" -exec dirname {} \; | jq -R -s -c 'split("\n")[:-1]')
          echo "matrix={\"package\":$MATRIX}" >> $GITHUB_OUTPUT
  
  test:
    needs: generate-matrix
    runs-on: ubuntu-latest
    strategy:
      matrix: ${{ fromJson(needs.generate-matrix.outputs.matrix) }}
    
    steps:
      - uses: actions/checkout@v4
      - run: cd ${{ matrix.package }} && npm test
```

### Workflow Orchestration
```yaml
jobs:
  trigger-downstream:
    runs-on: ubuntu-latest
    steps:
      - name: Trigger Integration Tests
        uses: actions/github-script@v7
        with:
          script: |
            await github.rest.actions.createWorkflowDispatch({
              owner: context.repo.owner,
              repo: 'integration-tests-repo',
              workflow_id: 'integration.yml',
              ref: 'main',
              inputs: {
                version: '${{ github.sha }}',
                environment: 'staging'
              }
            });
```

## 🧪 Testing Strategies

### Unit Testing Pattern
```yaml
jobs:
  unit-tests:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        test-group: [utils, components, services, controllers]
    
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '18'
          cache: 'npm'
      
      - run: npm ci
      - run: npm test -- --testNamePattern="${{ matrix.test-group }}"
      
      - name: Upload Coverage
        uses: codecov/codecov-action@v3
        with:
          files: coverage/lcov.info
          flags: ${{ matrix.test-group }}
```

### Integration Testing Pattern
```yaml
jobs:
  integration-tests:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
    
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '18'
      
      - run: npm ci
      - run: npm run test:integration
        env:
          DATABASE_URL: postgres://postgres:test@localhost:5432/test
```

### E2E Testing Pattern
```yaml
jobs:
  e2e-tests:
    runs-on: ubuntu-latest
    
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '18'
      
      - name: Start Application
        run: |
          npm start &
          npx wait-on http://localhost:3000
      
      - name: Run E2E Tests
        uses: cypress-io/github-action@v6
        with:
          wait-on: 'http://localhost:3000'
          browser: chrome
      
      - name: Upload Screenshots
        uses: actions/upload-artifact@v4
        if: failure()
        with:
          name: cypress-screenshots
          path: cypress/screenshots
```

## 🚀 Deployment Patterns

### Blue-Green Deployment
```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    
    steps:
      - name: Deploy to Green Environment
        run: |
          # Deploy to inactive environment
          kubectl apply -f k8s/deployment-green.yml
          kubectl wait --for=condition=ready pod -l app=myapp,env=green
      
      - name: Run Smoke Tests
        run: |
          # Test green environment
          curl -f https://green.example.com/health
      
      - name: Switch Traffic
        run: |
          # Switch load balancer to green
          kubectl patch service myapp-service -p '{"spec":{"selector":{"env":"green"}}}'
      
      - name: Cleanup Blue Environment
        run: |
          # Remove old blue deployment
          kubectl delete deployment myapp-blue || true
```

### Canary Deployment
```yaml
jobs:
  canary-deploy:
    runs-on: ubuntu-latest
    
    steps:
      - name: Deploy Canary (10%)
        run: |
          kubectl apply -f k8s/canary-deployment.yml
          kubectl scale deployment myapp-canary --replicas=1
          kubectl scale deployment myapp-stable --replicas=9
      
      - name: Monitor Metrics
        run: |
          # Monitor error rates, response times
          sleep 300  # Wait 5 minutes
          ERROR_RATE=$(curl -s "http://prometheus:9090/api/v1/query?query=error_rate" | jq '.data.result[0].value[1]')
          if (( $(echo "$ERROR_RATE > 0.05" | bc -l) )); then
            echo "Error rate too high, rolling back"
            exit 1
          fi
      
      - name: Promote Canary
        run: |
          kubectl scale deployment myapp-canary --replicas=10
          kubectl scale deployment myapp-stable --replicas=0
```

### Rolling Deployment
```yaml
jobs:
  rolling-deploy:
    runs-on: ubuntu-latest
    
    strategy:
      matrix:
        region: [us-east-1, us-west-2, eu-west-1]
    
    steps:
      - name: Deploy to ${{ matrix.region }}
        run: |
          aws ecs update-service \
            --cluster myapp-${{ matrix.region }} \
            --service myapp \
            --task-definition myapp:${{ github.sha }} \
            --region ${{ matrix.region }}
      
      - name: Wait for Deployment
        run: |
          aws ecs wait services-stable \
            --cluster myapp-${{ matrix.region }} \
            --services myapp \
            --region ${{ matrix.region }}
```

## 📊 Monitoring and Observability

### Performance Monitoring
```yaml
jobs:
  performance:
    runs-on: ubuntu-latest
    
    steps:
      - name: Lighthouse CI
        uses: treosh/lighthouse-ci-action@v10
        with:
          configPath: './.lighthouserc.json'
          temporaryPublicStorage: true
      
      - name: Performance Budget Check
        run: |
          SCORE=$(jq '.[] | select(.audits) | .audits.performance.score * 100' lhci_reports/manifest.json)
          if (( $(echo "$SCORE < 90" | bc -l) )); then
            echo "Performance score ($SCORE) below threshold (90)"
            exit 1
          fi
```

### Health Checks
```yaml
jobs:
  health-check:
    runs-on: ubuntu-latest
    
    steps:
      - name: Application Health Check
        uses: ./.github/actions/health-check
        with:
          url: 'https://app.example.com/health'
          timeout: '30'
          retry-count: '3'
      
      - name: Database Health Check
        run: |
          pg_isready -h db.example.com -p 5432 -U app
      
      - name: External Dependencies Check
        run: |
          curl -f https://api.external-service.com/health
          curl -f https://cdn.example.com/health
```

### Alerting Integration
```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    
    steps:
      - name: Deploy Application
        id: deploy
        run: |
          # Deployment logic
          echo "status=success" >> $GITHUB_OUTPUT
      
      - name: Notify Success
        if: steps.deploy.outputs.status == 'success'
        uses: 8398a7/action-slack@v3
        with:
          status: success
          channel: '#deployments'
          message: '✅ Deployment successful'
        env:
          SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK }}
      
      - name: Notify Failure
        if: failure()
        uses: 8398a7/action-slack@v3
        with:
          status: failure
          channel: '#alerts'
          message: '🚨 Deployment failed'
        env:
          SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK }}
```

## 💡 Best Practices Summary

1. **Keep workflows focused and single-purpose**
2. **Use descriptive names and comments**
3. **Implement proper error handling**
4. **Cache dependencies for faster execution**
5. **Use matrix strategies for parallel execution**
6. **Implement proper secret management**
7. **Add comprehensive logging and monitoring**
8. **Use conditional execution to save resources**
9. **Implement proper rollback strategies**
10. **Test workflows thoroughly before production use**

## 📚 Additional Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Workflow Syntax Reference](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)
- [Actions Marketplace](https://github.com/marketplace?type=actions)
- [Community Best Practices](https://github.com/actions/toolkit)

Remember: Good workflow design is iterative. Start simple, measure performance, and gradually add complexity as needed.
