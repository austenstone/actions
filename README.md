# 🚀 GitHub Actions Demo Repository

Welcome to the comprehensive GitHub Actions demonstration repository! 🚀

This repository showcases advanced GitHub Actions patterns, workflows, and best practices through practical, real-world examples. This comprehensive demo includes **8 sophisticated workflows**, **3 sample applications**, **3 custom actions**, and extensive documentation covering every aspect of GitHub Actions automation.

## 📊 Repository Overview

![GitHub Actions Demo](images/overview-actions-simple.png)

This repository serves as both a learning resource and a practical reference for implementing CI/CD pipelines, security scanning, performance monitoring, and more. Whether you're just starting with GitHub Actions or looking to implement advanced enterprise patterns, this repository has examples for every level.

### 🎯 What You'll Learn

- **CI/CD Pipeline Design**: From basic builds to complex deployment strategies
- **Security Integration**: Automated security scanning and compliance checks  
- **Performance Monitoring**: Lighthouse audits, load testing, and regression detection
- **Custom Actions**: Create reusable actions in JavaScript, Docker, and Composite formats
- **Advanced Patterns**: Matrix strategies, conditional execution, and workflow orchestration
- **Best Practices**: Security, performance, maintainability, and troubleshooting
- **Real-World Examples**: Production-ready workflows for immediate use

### 🔧 Repository Structure

```
├── .github/
│   ├── workflows/          # 8 comprehensive workflows
│   │   ├── 01-web-app-cicd.yml
│   │   ├── 02-multi-platform-release.yml
│   │   ├── 03-scheduled-maintenance.yml
│   │   ├── 04-security-deployment.yml
│   │   ├── 05-custom-actions-demo.yml
│   │   ├── 06-event-driven-workflows.yml
│   │   ├── 07-performance-optimization.yml
│   │   └── 08-interactive-dashboard.yml
│   └── actions/            # Custom actions
│       ├── smart-deploy/   # Composite action
│       ├── health-check/   # JavaScript action  
│       └── security-scanner/ # Docker action
├── src/                    # Sample applications
│   ├── web-app/           # Node.js Express application
│   ├── cli-tool/          # Go CLI application
│   └── microservice/      # Go microservice
├── docs/                   # Comprehensive documentation
├── examples/              # Additional examples and templates
└── images/               # Repository assets
```

## 📋 Overview

This demo repository contains 8 sophisticated workflows that demonstrate different aspects of GitHub Actions:

1. **Smart Web App CI/CD Pipeline** - Multi-platform testing, artifact sharing, security scanning, intelligent deployment
2. **Multi-Platform Release Automation** - Cross-platform builds, Docker images, GitHub releases, Homebrew formula
3. **Scheduled Maintenance & Monitoring** - Automated repository maintenance, dependency updates, security audits
4. **Security-First Deployment Pipeline** - CodeQL analysis, container scanning, OIDC authentication, policy validation
5. **Custom Action Ecosystem** - Composite, JavaScript, and Docker actions with reusable workflows
6. **Event-Driven Workflow Orchestra** - Complex workflow orchestration, artifact sharing, conditional execution
7. **Performance Optimization Pipeline** - Lighthouse audits, load testing, regression detection, dashboard updates
8. **Interactive Development Dashboard** - Real-time metrics, performance tracking, repository insights

## 🎯 Learning Paths

Choose your learning path based on your experience level:

### 👶 Beginner Path
Start with basic concepts and build your first workflow:
- Getting started with GitHub Actions
- Understanding workflows, jobs, and steps
- Basic CI/CD pipeline creation
- Simple deployment automation

### 🔧 Intermediate Path
Master advanced CI/CD patterns:
- Matrix strategies and job dependencies
- Artifact management and caching
- Multi-environment deployments
- Security integration

### 🚀 Advanced Path
Build enterprise-grade automation:
- Custom action development
- Complex workflow orchestration
- Security and compliance automation
- Performance optimization

### 🏗 Specialized Tracks
- **Security & Compliance**: Focus on security automation
- **Performance & Scale**: Optimize for speed and cost
- **Developer Experience**: Build user-friendly workflows

## 🛠 Workflows Included

### 01 - Smart Web App CI/CD Pipeline
**File**: `.github/workflows/01-web-app-cicd.yml`
**Demonstrates**:
- Multi-platform testing matrix
- Conditional execution and path filtering
- Artifact sharing between jobs
- Environment-specific deployments
- Security scanning integration

**Triggers**: Push to main/develop, PRs, manual dispatch

### 02 - Multi-Platform Release Automation
**File**: `.github/workflows/02-multi-platform-release.yml`
**Demonstrates**:
- Cross-platform binary compilation
- Automated GitHub releases
- Multi-architecture Docker images
- Package signing and SBOM generation
- Changelog automation

**Triggers**: Git tags, manual dispatch

### 03 - Scheduled Maintenance & Monitoring  
**File**: `.github/workflows/03-scheduled-maintenance.yml`
**Demonstrates**:
- Multiple cron schedules
- Repository maintenance automation
- Dependency auditing and updates
- Issue and PR cleanup
- Automated reporting

**Triggers**: Scheduled (daily/weekly/monthly), manual dispatch

### 04 - Security-First Deployment Pipeline
**File**: `.github/workflows/04-security-deployment.yml`
**Demonstrates**:
- CodeQL security scanning
- Container vulnerability assessment
- OIDC authentication
- Policy validation with OPA
- Secure deployment practices

**Triggers**: Push to main, manual dispatch with environment selection

### 05 - Custom Action Ecosystem
**File**: `.github/workflows/05-custom-actions-demo.yml`
**Demonstrates**:
- Composite actions (Smart Deploy)
- JavaScript actions (simulated)
- Docker actions (simulated)
- Reusable workflow patterns
- Action marketplace integration

**Triggers**: Manual dispatch with environment options

### 06 - Event-Driven Workflow Orchestra
**File**: `.github/workflows/06-event-driven-workflows.yml`
**Demonstrates**:
- Workflow orchestration
- Event-driven triggering
- Cross-workflow communication
- Artifact sharing between workflows
- Complex conditional logic

**Triggers**: Push to main, manual dispatch with deployment targets

### 07 - Performance Optimization Pipeline
**File**: `.github/workflows/07-performance-optimization.yml`
**Demonstrates**:
- Lighthouse performance audits
- Node.js and Go benchmarking
- Load testing with artillery and hey
- Performance regression detection
- Database performance testing
- Multi-format reporting (JSON, SARIF, Markdown)

**Triggers**: Push to main/develop (path-filtered), pull requests, scheduled daily, manual dispatch

### 08 - Interactive Development Dashboard
**File**: `.github/workflows/08-interactive-dashboard.yml`
**Demonstrates**:
- Dynamic metrics collection
- Real-time dashboard generation
- Interactive charts and visualizations
- Repository overview and insights
- Workflow status monitoring
- GitHub Pages deployment

**Triggers**: Push to main/develop, pull requests, scheduled hourly, manual dispatch with options

## 🏁 Quick Start

1. **Fork this repository** to your GitHub account

2. **Enable GitHub Actions** in your forked repository

3. **Configure secrets** (optional, for full functionality):
   ```
   SNYK_TOKEN - For security scanning
   AWS_ROLE_ARN - For OIDC authentication demos
   ```

4. **Run your first workflow**:
   - Go to Actions tab
   - Select "05 - Custom Action Ecosystem"
   - Click "Run workflow"
   - Choose your options and run

5. **Explore the results**:
   - Check the workflow summary
   - Review the logs
   - Examine the artifacts

## 📁 Repository Structure

```
├── .github/
│   ├── workflows/          # 8 comprehensive workflows
│   │   ├── 01-web-app-cicd.yml
│   │   ├── 02-multi-platform-release.yml
│   │   ├── 03-scheduled-maintenance.yml
│   │   ├── 04-security-deployment.yml
│   │   ├── 05-custom-actions-demo.yml
│   │   ├── 06-event-driven-workflows.yml
│   │   ├── 07-performance-optimization.yml
│   │   └── 08-interactive-dashboard.yml
│   └── actions/            # Custom actions
│       ├── smart-deploy/   # Composite action
│       ├── health-check/   # JavaScript action  
│       └── security-scanner/ # Docker action
├── src/                    # Sample applications
│   ├── web-app/           # Node.js Express application
│   ├── cli-tool/          # Go CLI application
│   └── microservice/      # Go microservice
├── docs/                   # Comprehensive documentation
├── examples/              # Additional examples and templates
└── images/               # Repository assets
```

## 🎓 What You'll Learn

### GitHub Actions Fundamentals
- Workflow syntax and structure
- Event triggers and conditions
- Job dependencies and matrix strategies
- Environment variables and secrets

### Advanced Patterns
- Custom action development
- Workflow orchestration
- Security automation
- Performance optimization
- Cost management

### Best Practices
- Security-first development
- Artifact management
- Error handling and rollback
- Monitoring and observability
- Documentation and maintenance

### Enterprise Features
- OIDC authentication
- Environment protection rules
- Compliance automation
- Approval workflows
- Audit and reporting

## 🔧 Prerequisites

- GitHub account with Actions enabled
- Basic understanding of YAML
- Familiarity with Git workflows
- Optional: Docker, Node.js, Go (for local development)

## 🚀 Running the Demos

### Method 1: GitHub Actions Tab
1. Navigate to the Actions tab in your repository
2. Select the workflow you want to run
3. Click "Run workflow"
4. Fill in any required inputs
5. Click "Run workflow" to execute

### Method 2: Git Events
Some workflows trigger automatically on:
- Push to main/develop branches
- Creating pull requests
- Creating git tags
- Scheduled times

### Method 3: API/CLI
Use GitHub CLI or API to trigger workflows:
```bash
gh workflow run "01-web-app-cicd.yml" --ref main
```

## 📊 Monitoring and Observability

Each workflow includes:
- **Comprehensive logging** with emojis for easy scanning
- **Step summaries** with rich markdown formatting
- **Artifact uploads** for build outputs and reports
- **Status checks** for deployment verification
- **Notification integration** (simulated)

## 🔒 Security Features

- **Code scanning** with CodeQL
- **Dependency scanning** with npm audit and Go vulnerability check
- **Container scanning** with Trivy
- **Secret scanning** with TruffleHog
- **Policy validation** with Open Policy Agent
- **OIDC authentication** for cloud providers
- **Image signing** with Cosign
- **SBOM generation** with Syft

## 💡 Tips for Learning

1. **Start small** - Begin with the beginner workflows
2. **Read the code** - Each workflow is heavily commented
3. **Experiment** - Modify workflows and see what happens
4. **Check the logs** - Learn from the detailed output
5. **Use the summaries** - Rich step summaries explain what happened
6. **Build incrementally** - Start with basic patterns and add complexity

## 📚 Documentation

### Core Documentation
- **[Getting Started](docs/getting-started.md)** - Step-by-step guide for new users
- **[Best Practices](docs/best-practices.md)** - Comprehensive best practices for workflows
- **[Workflow Design](docs/workflow-design.md)** - Advanced workflow design patterns and strategies
- **[Troubleshooting](docs/troubleshooting.md)** - Common issues and solutions

### Examples and Templates
- **[Examples](examples/)** - Additional workflow examples and templates
- **[Basic CI/CD](examples/workflows/basic-ci-cd.yml)** - Simple CI/CD pipeline template
- **[Advanced Patterns](examples/workflows/advanced-patterns.yml)** - Complex workflow patterns and techniques

### Action Documentation
- **[Smart Deploy](/.github/actions/smart-deploy/README.md)** - Composite action documentation
- **[Health Check](/.github/actions/health-check/README.md)** - JavaScript action documentation
- **[Security Scanner](/.github/actions/security-scanner/README.md)** - Docker action documentation

## 🎓 Learning Paths

### Beginner Path (⭐)
**Goal**: Learn GitHub Actions fundamentals

1. **Start Here**: Read [Getting Started](docs/getting-started.md)
2. **Basic Workflow**: Study `01-web-app-cicd.yml`
3. **Run Your First Workflow**: Try the web app CI/CD pipeline
4. **Simple Example**: Explore [Basic CI/CD](examples/workflows/basic-ci-cd.yml)
5. **First Custom Action**: Use the `smart-deploy` composite action

**Estimated Time**: 2-4 hours

### Intermediate Path (⭐⭐)
**Goal**: Master CI/CD automation and security

1. **Multi-Platform Builds**: Study `02-multi-platform-release.yml`
2. **Security Integration**: Explore `04-security-deployment.yml`
3. **Scheduled Automation**: Understand `03-scheduled-maintenance.yml`
4. **Custom Actions**: Build with `health-check` JavaScript action
5. **Advanced Patterns**: Review [Advanced Patterns](examples/workflows/advanced-patterns.yml)

**Estimated Time**: 1-2 days

### Advanced Path (⭐⭐⭐)
**Goal**: Implement enterprise-grade automation

1. **Performance Monitoring**: Deep dive into `07-performance-optimization.yml`
2. **Event-Driven Architecture**: Master `06-event-driven-workflows.yml`
3. **Interactive Dashboards**: Implement `08-interactive-dashboard.yml`
4. **Custom Docker Actions**: Create with `security-scanner` action
5. **Workflow Orchestration**: Design complex multi-workflow systems

**Estimated Time**: 3-5 days

### Expert Path (⭐⭐⭐⭐)
**Goal**: Design organization-wide automation strategies

1. **Enterprise Patterns**: Implement organization-wide standards
2. **Custom Marketplace Actions**: Publish reusable actions
3. **Advanced Security**: Implement comprehensive security automation
4. **Performance Optimization**: Design high-performance workflows
5. **Compliance Automation**: Build regulatory compliance systems

**Estimated Time**: 1-2 weeks

### Specialized Learning Tracks

#### 🔒 Security & Compliance Track
Focus on security automation and compliance:
- Security scanning workflows (`04-security-deployment.yml`)
- Custom security actions (`security-scanner`)
- OIDC authentication patterns
- Compliance reporting automation
- Secret management best practices

#### ⚡ Performance & Scale Track  
Optimize for speed and efficiency:
- Performance monitoring (`07-performance-optimization.yml`)
- Workflow optimization techniques
- Caching strategies
- Resource management
- Cost optimization

#### 🎯 Developer Experience Track
Build user-friendly automation:
- Interactive workflows (`08-interactive-dashboard.yml`)
- Rich reporting and summaries
- Developer-friendly notifications
- Self-service automation
- Documentation automation

## 🏃 Quick Start Guide

### 1. Fork and Setup
```bash
# Fork this repository to your GitHub account
# Clone your fork
git clone https://github.com/YOUR_USERNAME/actions-demo.git
cd actions-demo

# Explore the structure
tree -L 3
```

### 2. Configure Repository
1. **Enable GitHub Actions** in your repository settings
2. **Configure secrets** (optional, for full functionality):
   ```
   GITHUB_TOKEN - Automatically provided
   SNYK_TOKEN - For security scanning (get from snyk.io)
   ```
3. **Set up environments** (Settings → Environments):
   - `staging` - For staging deployments
   - `production` - For production deployments

### 3. Run Your First Workflow
1. Go to the **Actions** tab in your repository
2. Select **"05 - Custom Action Ecosystem"** 
3. Click **"Run workflow"**
4. Choose your options:
   - Environment: `staging`
   - Include performance tests: `true`
5. Click **"Run workflow"** to execute

### 4. Explore Sample Applications
```bash
# Web Application
cd src/web-app
npm install
npm start
# Visit http://localhost:3000

# CLI Tool
cd ../cli-tool
go build
./cli-tool version

# Microservice  
cd ../microservice
go build
./microservice
# Visit http://localhost:8080/health
```

### 5. Test Workflows Locally (Optional)
```bash
# Install act for local testing
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Test a workflow locally
act -j test-web-app -P ubuntu-latest=catthehacker/ubuntu:act-latest
```

## 🔧 Customization Guide

### Adapting for Your Technology Stack

#### Python Projects
```yaml
- name: Setup Python
  uses: actions/setup-python@v4
  with:
    python-version: '3.11'
    cache: 'pip'

- name: Install Dependencies
  run: |
    pip install -r requirements.txt
    pip install pytest pytest-cov

- name: Run Tests
  run: pytest --cov=. --cov-report=xml
```

#### Java Projects
```yaml
- name: Setup Java
  uses: actions/setup-java@v3
  with:
    java-version: '17'
    distribution: 'temurin'
    cache: 'maven'

- name: Run Tests
  run: ./mvnw clean test

- name: Build Application
  run: ./mvnw clean package -DskipTests
```

#### .NET Projects
```yaml
- name: Setup .NET
  uses: actions/setup-dotnet@v3
  with:
    dotnet-version: '8.0'

- name: Restore Dependencies
  run: dotnet restore

- name: Run Tests
  run: dotnet test --no-restore --logger trx
```

#### Docker Projects
```yaml
- name: Set up Docker Buildx
  uses: docker/setup-buildx-action@v3

- name: Build and Test
  uses: docker/build-push-action@v5
  with:
    context: .
    target: test
    load: true
    cache-from: type=gha
    cache-to: type=gha,mode=max
```

### Environment Configuration

#### Development Environment
- Automated testing on every push
- Code quality checks
- Security scanning
- Performance monitoring

#### Staging Environment  
- Pre-production testing
- Integration testing
- Performance testing
- Security validation

#### Production Environment
- Approval-gated deployments
- Blue-green or canary deployments
- Comprehensive monitoring
- Automated rollback

### Integration Examples

#### Cloud Providers
```yaml
# AWS Integration
- name: Configure AWS Credentials
  uses: aws-actions/configure-aws-credentials@v4
  with:
    role-to-assume: ${{ secrets.AWS_ROLE_ARN }}
    aws-region: us-east-1

# Azure Integration  
- name: Azure Login
  uses: azure/login@v1
  with:
    creds: ${{ secrets.AZURE_CREDENTIALS }}

# GCP Integration
- name: Authenticate to Google Cloud
  uses: google-github-actions/auth@v1
  with:
    workload_identity_provider: ${{ secrets.WIF_PROVIDER }}
    service_account: ${{ secrets.WIF_SERVICE_ACCOUNT }}
```

#### Monitoring Integration
```yaml
# Datadog
- name: Datadog Event
  uses: masci/datadog@v1
  with:
    api-key: ${{ secrets.DATADOG_API_KEY }}

# New Relic
- name: New Relic Deployment
  uses: newrelic/deployment-marker-action@v2
  with:
    apiKey: ${{ secrets.NEW_RELIC_API_KEY }}
```