# Troubleshooting Guide

This guide helps you resolve common issues when working with GitHub Actions.

## 📋 Table of Contents

- [Common Issues](#common-issues)
- [Workflow Errors](#workflow-errors)
- [Action-Specific Issues](#action-specific-issues)
- [Performance Issues](#performance-issues)
- [Security Issues](#security-issues)
- [Debugging Tips](#debugging-tips)
- [Getting Help](#getting-help)

## 🚨 Common Issues

### Workflow Not Triggering

**Problem**: Workflow doesn't run when expected.

**Common Causes & Solutions**:

1. **Branch name mismatch**
   ```yaml
   # ❌ Wrong - using 'master' when branch is 'main'
   on:
     push:
       branches: [master]
   
   # ✅ Correct
   on:
     push:
       branches: [main]
   ```

2. **Path filters too restrictive**
   ```yaml
   # ❌ Too specific
   on:
     push:
       paths: ['src/app.js']
   
   # ✅ Better
   on:
     push:
       paths: ['src/**']
   ```

3. **Workflow file in wrong location**
   - Must be in `.github/workflows/` directory
   - Must have `.yml` or `.yaml` extension

### Permission Denied Errors

**Problem**: Jobs fail with permission errors.

**Solutions**:

1. **Add necessary permissions**
   ```yaml
   jobs:
     deploy:
       runs-on: ubuntu-latest
       permissions:
         contents: read
         packages: write
         id-token: write
   ```

2. **Check repository settings**
   - Settings → Actions → General
   - Ensure proper workflow permissions are set

### Environment Variables Not Working

**Problem**: Environment variables are undefined or empty.

**Solutions**:

1. **Check variable scope**
   ```yaml
   # ✅ Job-level environment variable
   jobs:
     build:
       env:
         NODE_VERSION: '18'
       steps:
         - run: echo $NODE_VERSION
   
   # ✅ Step-level environment variable
   steps:
     - run: echo $NODE_VERSION
       env:
         NODE_VERSION: '18'
   ```

2. **Use proper syntax for outputs**
   ```yaml
   # ❌ Wrong
   run: echo "::set-output name=version::1.0.0"
   
   # ✅ Correct (new syntax)
   run: echo "version=1.0.0" >> $GITHUB_OUTPUT
   ```

## 🔧 Workflow Errors

### YAML Syntax Errors

**Problem**: Workflow fails to parse due to YAML syntax issues.

**Common Issues**:

1. **Indentation problems**
   ```yaml
   # ❌ Wrong indentation
   jobs:
   build:
     runs-on: ubuntu-latest
   
   # ✅ Correct indentation
   jobs:
     build:
       runs-on: ubuntu-latest
   ```

2. **Missing quotes for special characters**
   ```yaml
   # ❌ Wrong - contains special character
   run: echo Hello: World
   
   # ✅ Correct
   run: echo "Hello: World"
   ```

3. **Incorrect multiline syntax**
   ```yaml
   # ❌ Wrong
   run: |
   echo "line 1"
   echo "line 2"
   
   # ✅ Correct
   run: |
     echo "line 1"
     echo "line 2"
   ```

### Matrix Strategy Issues

**Problem**: Matrix jobs fail or don't run as expected.

**Solutions**:

1. **Handle matrix exclusions properly**
   ```yaml
   strategy:
     matrix:
       os: [ubuntu-latest, windows-latest, macos-latest]
       node: [16, 18, 20]
       exclude:
         - os: macos-latest
           node: 16
   ```

2. **Use proper matrix variable references**
   ```yaml
   # ✅ Correct matrix variable usage
   steps:
     - uses: actions/setup-node@v4
       with:
         node-version: ${{ matrix.node }}
   ```

### Conditional Logic Problems

**Problem**: Jobs run when they shouldn't or don't run when they should.

**Solutions**:

1. **Use proper condition syntax**
   ```yaml
   # ✅ Multiple conditions
   if: github.ref == 'refs/heads/main' && github.event_name == 'push'
   
   # ✅ Checking job results
   if: needs.test.result == 'success'
   
   # ✅ Checking for specific values
   if: contains(github.event.head_commit.message, '[deploy]')
   ```

2. **Debug conditions**
   ```yaml
   - name: Debug context
     run: |
       echo "Event name: ${{ github.event_name }}"
       echo "Ref: ${{ github.ref }}"
       echo "Actor: ${{ github.actor }}"
   ```

## 🎯 Action-Specific Issues

### Checkout Action Issues

**Problem**: Code not available or wrong version checked out.

**Solutions**:

1. **Specify fetch depth for git operations**
   ```yaml
   - uses: actions/checkout@v4
     with:
       fetch-depth: 0  # Fetch full history
   ```

2. **Checkout specific ref**
   ```yaml
   - uses: actions/checkout@v4
     with:
       ref: ${{ github.event.pull_request.head.sha }}
   ```

### Setup Actions Failing

**Problem**: Setup actions fail to install or configure tools.

**Solutions**:

1. **Use caching for faster installs**
   ```yaml
   - uses: actions/setup-node@v4
     with:
       node-version: '18'
       cache: 'npm'
       cache-dependency-path: package-lock.json
   ```

2. **Specify exact versions**
   ```yaml
   # ✅ Better - specific version
   - uses: actions/setup-node@v4
     with:
       node-version: '18.17.0'
   ```

### Upload/Download Artifact Issues

**Problem**: Artifacts not found or corrupted.

**Solutions**:

1. **Check artifact names match exactly**
   ```yaml
   # Upload
   - uses: actions/upload-artifact@v4
     with:
       name: test-results
   
   # Download (same job or different job)
   - uses: actions/download-artifact@v4
     with:
       name: test-results  # Must match exactly
   ```

2. **Handle multiple artifacts**
   ```yaml
   - uses: actions/download-artifact@v4
     with:
       pattern: test-results-*
       merge-multiple: true
   ```

## ⚡ Performance Issues

### Slow Workflow Execution

**Problem**: Workflows take too long to complete.

**Solutions**:

1. **Use dependency caching**
   ```yaml
   - uses: actions/setup-node@v4
     with:
       node-version: '18'
       cache: 'npm'
   
   - uses: actions/cache@v3
     with:
       path: ~/.npm
       key: ${{ runner.os }}-node-${{ hashFiles('**/package-lock.json') }}
   ```

2. **Parallel job execution**
   ```yaml
   jobs:
     test:
       strategy:
         matrix:
           test-group: [unit, integration, e2e]
   ```

3. **Skip unnecessary jobs**
   ```yaml
   jobs:
     test:
       if: contains(github.event.head_commit.modified, 'src/')
   ```

### Runner Resource Issues

**Problem**: Jobs fail due to resource constraints.

**Solutions**:

1. **Use appropriate runner size**
   ```yaml
   runs-on: ubuntu-latest-4-cores  # For CPU-intensive tasks
   ```

2. **Clean up resources**
   ```yaml
   - name: Cleanup
     if: always()
     run: |
       docker system prune -f
       rm -rf node_modules
   ```

## 🔒 Security Issues

### Secret Management Problems

**Problem**: Secrets not working or exposed.

**Solutions**:

1. **Check secret availability in environment**
   ```yaml
   # Secrets not available in pull requests from forks
   if: github.event_name != 'pull_request' || github.event.pull_request.head.repo.full_name == github.repository
   ```

2. **Mask sensitive values**
   ```yaml
   run: |
     echo "::add-mask::${{ secrets.API_KEY }}"
     echo "API key is: ${{ secrets.API_KEY }}"
   ```

### Permission Issues

**Problem**: Actions fail due to insufficient permissions.

**Solutions**:

1. **Set minimal required permissions**
   ```yaml
   permissions:
     contents: read
     pull-requests: write
     checks: write
   ```

2. **Use GITHUB_TOKEN appropriately**
   ```yaml
   - uses: actions/github-script@v7
     with:
       github-token: ${{ secrets.GITHUB_TOKEN }}
   ```

## 🐛 Debugging Tips

### Add Debug Logging

```yaml
- name: Debug Information
  run: |
    echo "Event: ${{ github.event_name }}"
    echo "Ref: ${{ github.ref }}"
    echo "SHA: ${{ github.sha }}"
    echo "Actor: ${{ github.actor }}"
    env | sort
```

### Use tmate for Interactive Debugging

```yaml
- name: Setup tmate session
  if: failure()
  uses: mxschmitt/action-tmate@v3
  with:
    limit-access-to-actor: true
```

### Enable Debug Logging

Set repository secrets:
- `ACTIONS_STEP_DEBUG` = `true`
- `ACTIONS_RUNNER_DEBUG` = `true`

### Test Locally

Use tools like [act](https://github.com/nektos/act):
```bash
# Install act
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Run workflow locally
act -j test
```

## 📞 Getting Help

### Check GitHub Status

- [GitHub Status Page](https://www.githubstatus.com/)
- [GitHub Actions Status](https://www.githubstatus.com/incidents)

### Documentation Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Workflow Syntax](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)
- [Marketplace Actions](https://github.com/marketplace?type=actions)

### Community Support

- [GitHub Community Forum](https://github.community/)
- [Stack Overflow](https://stackoverflow.com/questions/tagged/github-actions)
- [GitHub Actions Discord](https://discord.gg/github)

### Create Support Ticket

For issues with GitHub Actions service:
1. Go to [GitHub Support](https://support.github.com/)
2. Select "Actions" as the topic
3. Provide detailed information about your issue

## 🔍 Diagnostic Checklist

When troubleshooting, check:

- [ ] Workflow file syntax is valid YAML
- [ ] File is in `.github/workflows/` directory
- [ ] Trigger conditions are correct
- [ ] Branch names match exactly
- [ ] Required secrets are available
- [ ] Runner has necessary permissions
- [ ] Dependencies are properly cached
- [ ] Resource limits are sufficient
- [ ] Environment variables are set correctly
- [ ] Action versions are compatible

## 📝 Common Error Messages

### "Resource not accessible by integration"
**Cause**: Insufficient permissions
**Solution**: Add required permissions to job

### "No space left on device"
**Cause**: Runner out of disk space
**Solution**: Clean up files or use larger runner

### "Process completed with exit code 1"
**Cause**: Command failed
**Solution**: Check command syntax and dependencies

### "Unable to resolve action"
**Cause**: Action reference is incorrect
**Solution**: Check action name and version

### "Required context variable was not provided"
**Cause**: Missing required input or environment variable
**Solution**: Provide required values

Remember: Most issues can be resolved by carefully reading error messages and checking the workflow configuration against the documentation.
