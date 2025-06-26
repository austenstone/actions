# GitHub Actions Best Practices

This document outlines best practices for GitHub Actions based on the patterns demonstrated in this repository.

## 🏗️ Workflow Design

### Structure and Organization
- **Use descriptive workflow names** with consistent numbering/prefixes
- **Group related steps** into logical jobs
- **Keep workflows focused** on a single purpose
- **Use meaningful job and step names** that clearly indicate their purpose

### Event Triggers
- **Be specific with triggers** to avoid unnecessary runs
- **Use path filters** to trigger only when relevant files change
- **Combine multiple trigger types** for comprehensive automation
- **Use workflow_dispatch** for manual testing and debugging

### Job Dependencies
- **Use `needs`** to create proper job dependencies
- **Minimize dependencies** to maximize parallelization
- **Use conditional execution** (`if`) to skip unnecessary jobs
- **Share data between jobs** using outputs and artifacts

## 🔧 Performance Optimization

### Caching Strategy
```yaml
- uses: actions/cache@v4
  with:
    path: |
      ~/.npm
      ~/.cache/pip
      ~/.gradle/caches
    key: ${{ runner.os }}-deps-${{ hashFiles('**/package-lock.json') }}
    restore-keys: |
      ${{ runner.os }}-deps-
```

### Matrix Builds
```yaml
strategy:
  matrix:
    os: [ubuntu-latest, windows-latest, macos-latest]
    node-version: [16, 18, 20]
  fail-fast: false  # Continue other jobs if one fails
```

### Conditional Execution
```yaml
- name: Deploy to production
  if: github.ref == 'refs/heads/main' && github.event_name == 'push'
```

### Resource Management
- **Use appropriate runner types** (ubuntu-latest is fastest/cheapest)
- **Cache dependencies** to speed up builds
- **Use build matrices** for parallel execution
- **Set timeouts** to prevent runaway jobs

## 🔒 Security Best Practices

### Secrets Management
- **Never hardcode secrets** in workflow files
- **Use GitHub Secrets** for sensitive data
- **Scope secrets appropriately** (repository vs environment)
- **Rotate secrets regularly**

### Action Security
- **Pin action versions** to specific commits or tags
- **Review third-party actions** before using
- **Use official actions** when available
- **Keep actions updated** but test before upgrading

### Permission Management
```yaml
permissions:
  contents: read
  security-events: write
  packages: write
  id-token: write  # For OIDC
```

### OIDC Authentication
```yaml
- name: Configure AWS credentials
  uses: aws-actions/configure-aws-credentials@v4
  with:
    role-to-assume: ${{ secrets.AWS_ROLE_ARN }}
    aws-region: us-east-1
```

## 🚀 Deployment Practices

### Environment Strategy
- **Use environment protection rules** for production
- **Implement approval workflows** for critical deployments
- **Use environment-specific secrets** and variables
- **Test in staging** before production deployment

### Rollback Capability
- **Plan for rollback scenarios** in deployment workflows
- **Keep previous versions** accessible
- **Implement health checks** post-deployment
- **Use blue-green or canary deployments** for zero-downtime

### Artifact Management
```yaml
- name: Upload artifacts
  uses: actions/upload-artifact@v4
  with:
    name: build-artifacts-${{ github.sha }}
    path: dist/
    retention-days: 30
```

## 🧪 Testing Integration

### Test Strategy
- **Run tests early** in the workflow
- **Use test matrices** for multiple environments
- **Separate unit and integration tests**
- **Generate test reports** and upload artifacts

### Quality Gates
- **Fail fast** on test failures
- **Require status checks** for merge protection
- **Use quality thresholds** (coverage, security scans)
- **Block deployments** on test failures

## 📊 Monitoring and Observability

### Logging Best Practices
- **Use descriptive log messages** with consistent formatting
- **Include context** (commit SHA, environment, version)
- **Use emojis** for visual scanning (as shown in demos)
- **Log important decisions** and their reasoning

### Workflow Summaries
```yaml
- name: Create summary
  run: |
    echo "## 🚀 Deployment Summary" >> $GITHUB_STEP_SUMMARY
    echo "- **Version**: ${{ github.sha }}" >> $GITHUB_STEP_SUMMARY
    echo "- **Environment**: production" >> $GITHUB_STEP_SUMMARY
    echo "- **Status**: ✅ Success" >> $GITHUB_STEP_SUMMARY
```

### Status Checks
```yaml
- name: Update commit status
  uses: actions/github-script@v7
  with:
    script: |
      github.rest.repos.createCommitStatus({
        owner: context.repo.owner,
        repo: context.repo.repo,
        sha: context.sha,
        state: 'success',
        description: 'Deployment completed',
        context: 'deploy/production'
      })
```

## 🔄 Custom Actions

### Composite Actions
- **Group related steps** into reusable actions
- **Use clear input/output definitions**
- **Include proper documentation**
- **Version your actions** appropriately

### JavaScript Actions
- **Use TypeScript** for better maintainability
- **Include comprehensive tests**
- **Handle errors gracefully**
- **Follow semantic versioning**

### Docker Actions
- **Use minimal base images** (alpine, scratch)
- **Run as non-root user**
- **Include health checks**
- **Optimize for security**

## 📋 Maintenance

### Workflow Maintenance
- **Review workflows regularly** for optimization opportunities
- **Update actions** to latest versions
- **Clean up unused workflows** and artifacts
- **Monitor workflow performance** and costs

### Documentation
- **Document workflow purpose** and usage
- **Include examples** in README files
- **Maintain troubleshooting guides**
- **Keep documentation up-to-date**

### Dependency Management
- **Automate dependency updates** (Dependabot, Renovate)
- **Test updates** in development environments
- **Monitor security advisories**
- **Pin critical dependencies**

## 💡 Advanced Patterns

### Workflow Orchestration
- **Use repository_dispatch** for cross-repo workflows
- **Implement workflow_run** for sequential execution
- **Create reusable workflows** for common patterns
- **Use matrix strategies** for complex scenarios

### Dynamic Workflows
- **Generate job matrices** dynamically
- **Use workflow inputs** for flexibility
- **Implement conditional job creation**
- **Create self-configuring workflows**

### Enterprise Features
- **Implement compliance checks** and reporting
- **Use self-hosted runners** for specific requirements
- **Create organization-wide** action libraries
- **Implement workflow governance** and standards

## 🚫 Common Anti-Patterns to Avoid

### Performance Anti-Patterns
- ❌ Running unnecessary jobs for every change
- ❌ Not using caching for dependencies
- ❌ Using serial execution when parallel is possible
- ❌ Not setting appropriate timeouts

### Security Anti-Patterns
- ❌ Hardcoding secrets in workflow files
- ❌ Using `pull_request_target` without careful review
- ❌ Not pinning action versions
- ❌ Overly broad permissions

### Maintenance Anti-Patterns
- ❌ Duplicating logic across workflows
- ❌ Not documenting complex workflows
- ❌ Ignoring workflow failures
- ❌ Not cleaning up old artifacts

## 📈 Metrics and KPIs

### Workflow Performance
- **Execution time** trends
- **Success/failure rates**
- **Cost per workflow run**
- **Queue times** and concurrency

### Development Velocity
- **Time to deployment**
- **Deployment frequency**
- **Lead time for changes**
- **Mean time to recovery**

### Quality Metrics
- **Test coverage** trends
- **Security scan** results
- **Code quality** scores
- **Compliance** status

## 🎯 Conclusion

Following these best practices will help you create reliable, secure, and maintainable GitHub Actions workflows. Remember that GitHub Actions is constantly evolving, so stay updated with the latest features and security recommendations.

The workflows in this repository demonstrate many of these practices in action. Use them as a reference and starting point for your own automation journey!
