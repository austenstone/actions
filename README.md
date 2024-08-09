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

### GitHub CLI: Every Runner Pre-loaded with it

### VS Code Extension

### Copilot

### Actions Loves JavaScript

### Actions Toolkit

### Custom Actions

### Github-script

### Developer Loop: Writing, Testing, Debugging

### Expressions

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
