# GitHub Actions Overview

## Intro to Concepts

### Workflow, Job, Step

### Action: Marketplace Action, Custom Actions (Composite Action)

### Runner: GitHub-Hosted Runner (GHR) and Self-Hosted Runner (SHR)

### ARC: Actions Runtime Configuration

### Workflow Trigger

## GitHub Actions Ethos

### Essence/Opinionation

### Traceability

### Reusability (DRY)

### Ephemerality

### Extensibility: Do Anything You Want

### Easy to Get Started

## How to Author Workflow Files

### The Developer Loop: Writing, Testing, Debugging

Writing a workflow file is as simple as creating a `.yml` file in the `.github/workflows` directory of your repository.

To test your workflow file you will push it to your repository and navigate to the Actions tab to see the status of your workflow run.

When the workflow run is complete you can view the logs of each step to see what happened.

### GitHub CLI

The GitHub CLI brings GitHub to the terminal. It's also preinstalled on all GitHub runners!

If you need to quickly perform a GitHub task this is the easiest way to do it!

<details>
  <summary>Comment on an issue</summary>
  
```yml
on:
  issues:
    types:
      - opened
jobs:
  comment:
    runs-on: ubuntu-latest
    steps:
      - run: gh issue comment $ISSUE --body "Thank you for opening this issue!"
        env:
          GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          ISSUE: ${{ github.event.issue.html_url }}
```
</details>

For the list of available extensions for the gh cli, see the topic [`gh-extension`](https://github.com/topics/gh-extension).

[Install](https://cli.github.com/)
[Manual](https://cli.github.com/manual/)
[Using GitHub CLI in workflows](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/using-github-cli-in-workflows)

### VS Code extension

There is a VS Code extension that provides syntax highlighting, intellisense, and more! This is a must have when authoring workflows.

[GitHub Actions Extension](https://marketplace.visualstudio.com/items?itemName=GitHub.vscode-github-actions)

![alt text](images/vscode-extension-1.png)

### Copilot

GitHub Copilot is an AI pair programmer that helps you write code faster and with less effort. It can be incredibly useful when writing GitHub Actions workflows. Leverage the completion or chat feature to get help with writing your workflows.

[GitHub Copilot](https://copilot.github.com/)

### Actions Loves JavaScript

One of the most popular languages for writing actions is JavaScript. This is because it is easy to get started with and has a lot of community support.

#### GitHub Actions ToolKit

The GitHub Actions ToolKit provides a set of packages to make creating actions easier.

* [Actions Toolkit](https://github.com/actions/toolkit?tab=readme-ov-file#readme)
* [Creating a JavaScript action](https://docs.github.com/en/actions/creating-actions/creating-a-javascript-action)

### Custom Actions

There are three types of custom actions:
* [JavaScript](https://docs.github.com/en/actions/creating-actions/creating-a-javascript-action)
* [Docker](https://docs.github.com/en/actions/creating-actions/creating-a-docker-container-action) (Not available on macOS or Windows runners)
* [Composite](https://docs.github.com/en/actions/creating-actions/creating-a-composite-run-steps-action)

* [About custom actions](https://docs.github.com/en/actions/creating-actions/about-custom-actions)

### Github-script

This action makes it easy to quickly write a script in your workflow that uses the GitHub API and the workflow run context. The GitHub Actions Toolkit is pre-installed and available for use in the script you write.

<details>
  <summary>Welcome a first-time contributor</summary>

```yml
on: pull_request_target

jobs:
  welcome:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/github-script@v7
        with:
          script: |
            // Get a list of all issues created by the PR opener
            // See: https://octokit.github.io/rest.js/#pagination
            const creator = context.payload.sender.login
            const opts = github.rest.issues.listForRepo.endpoint.merge({
              ...context.issue,
              creator,
              state: 'all'
            })
            const issues = await github.paginate(opts)

            for (const issue of issues) {
              if (issue.number === context.issue.number) {
                continue
              }

              if (issue.pull_request) {
                return // Creator is already a contributor.
              }
            }

            await github.rest.issues.createComment({
              issue_number: context.issue.number,
              owner: context.repo.owner,
              repo: context.repo.repo,
              body: `**Welcome**, new contributor!

                Please make sure you've read our [contributing guide](CONTRIBUTING.md) and we look forward to reviewing your Pull request shortly ✨`
            })
```
</details>

<details>
  <summary>Download data from a URL</summary>

```yml
on: pull_request

jobs:
  diff:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/github-script@v7
        with:
          script: |
            const diff_url = context.payload.pull_request.diff_url
            const result = await github.request(diff_url)
            console.log(result)
```
</details>

### Expressions

[Expressions](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions)

## How to Trigger/Initiate Workflow Runs

### Summary of Event Grid That Triggers Workflows

### Configuring Input (Activity Types): Conditionally Trigger

### Event: Workflow Dispatch + Inputs and Outputs

### Event: Workflow Run

### Event: Schedule

### Non-Core CI/CD Use Cases

### Concurrency: Order of Workflow Runs Based on When Trigger Happened

### Re-running Workflows

## How to Structure/Manage Jobs in the Workflow

### Understanding the 1:1 Job-to-Runner Mapping

### Parallelization of Jobs

### Linking Jobs

### Matrices

### Job Timeouts

### Sharing Artifacts Between Jobs

### Running Jobs in Containers / Service Containers

### Environments: Controls How/When a Job is Run Based on Protection Rules Set, Limits Branches, Scopes Secrets

### Conditional Jobs/Steps

### Permissions for Jobs

There is a default token called `GITHUB_TOKEN`.

[Assigning permissions to jobs](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/assigning-permissions-to-jobs)

## How to Use and Create Actions (Marketplace)

### What is an Action?

### Types of Actions

### Action Outputs

### Securing Usage of Actions

### Creating Your Own Actions

### Cool Actions to Look Out For: github-script, Anything by GitHub, Major Cloud Providers, Terraform, Docker

## How to Organize, Share, and Scale Workflows

### Reusable Workflows (and Outputs)

### Composite Actions

### Required Workflows (per Repo Rulesets)

### Starter Workflows

### Managing Updates to Workflows/Actions

### Monorepo vs Polyrepo

## How to Manage Artifacts / Caching

### Artifact Attestations

### Retention Period

### Sharing Artifacts Between Jobs

### Caching Dependencies

### Caching Limits

### Rotating Cache

## How to Create and Manage Secrets

### GitHub Secret Store (libsodium)

### Org-Level, Repo-Level, and Env-Level Secrets (Scope)

### Scoping/Limiting Secrets

### OIDC Access to Cloud Environments

### Reusable Workflows and Secrets

### Redacting Secrets

### Integrating with 3rd Party Key Vaults/HSMs

## How to Create and Manage Runners

### Based on Azure and Ephemeral

### Types and Sizes of Runners (and OSes)

### GHRs vs SHRs

### ARC: Actions Runtime Configuration

### Runner Groups

### Managed Runner GHR Images and Custom GHR Images

### GHR Networking

### Static IPs

### VNET Injection (Azure Private Networking)

### API Gateways

### Wireguard

### Runner Labels

### Auto-scaling and Scale Limits

## How to Govern Usage

### Repository Rulesets

### Push Rulesets: Control Who Can Push Changes

### Branch Rulesets

### Environment Protection Rules: Custom Gating

### Allow List for Marketplace Actions

### Spending Limits and Budgets/Cost Centers (BvN)

### Actions Policies

### Allow List for Marketplace Actions

### Enable/Disable Actions

### Audit Log

### Status Checks

## How to Observe What’s Going on with CI/CD

### Actions Usage Metrics

### Job Summaries

### Alerting/Notifications: Finished, Failed

### 3rd Party Tools: DataDog, Trunk, etc.

### Stats on Runner Utilization

### Workflow Run History

### Pinning and Searching Workflows

### Logging

## How to Manage Cost and Billing

### Pricing

### Entitlements

### Billing Page in GitHub (Soon to be BvN)

### CSV Usage Download and GitHub Usage Report Viewer

### Invoicing

### Paying Through GitHub vs MSFT Azure

## How to Migrate

### Importer

### VS Code Extension + Copilot

## Understanding Platform Limits

### Concurrency Limit

### API Rate Limits

### Reusable Workflow Limit

### Workflow Run Time

### Job Execution Time

### Matrix Size

### Cache Size per Repo

### Queue Limit

### Artifact and Log Retention

### Exception Process

[Usage Limits, Billing, and Administration Documentation](https://docs.github.com/en/actions/administering-github-actions/usage-limits-billing-and-administration)
