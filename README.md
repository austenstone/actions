# GitHub Actions Overview

This document is a high-level overview of GitHub Actions. It is not intended to be a comprehensive guide to the platform, but rather a starting point for understanding the basics.

## Intro to Concepts

There are a few concepts that are important to understand when working with GitHub Actions.

### Definitions

Some basic definitions to get us started...

![overview-actions-simple](images/overview-actions-simple.png)

#### [Workflow](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#workflows)

A workflow is a configurable automated process that will run one or more jobs. Workflows are defined by a YAML file checked in to your repository in the `.github/workflows` directory. A repository can have multiple workflows, each of which can perform a different set of tasks.

#### [Events](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#events)

An event is a specific activity in a repository that triggers a workflow run. It could be triggered by an event in your repository, or they can be triggered manually, or at a defined schedule.

#### [Jobs](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#jobs)

A job is a set of steps in a workflow that is executed on the same runner. Each step is either a shell script that will be executed, or an action that will be run. Steps are executed in order and are dependent on each other. Since each step is executed on the same runner, you can share data from one step to another.

#### [Steps / Actions](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#actions)

A step can be a script that will be executed or a GitHub action.

#### [Runners](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#runners)

A runner is a server that runs your workflows when they're triggered. Each runner can run a single job at a time.

* GHR: GitHub-Hosted Runner
* SHR: Self-Hosted Runner

### Action: Marketplace Action, Custom Actions (Composite Action)

An action is a custom application for the GitHub Actions platform that performs a complex but frequently repeated task. Use an action to help reduce the amount of repetitive code that you write in your workflow files. 

An action can pull your git repository from GitHub, set up the correct toolchain for your build environment, or set up the authentication to your cloud provider.

You can write your own actions, or you can find actions to use in your workflows in the [GitHub Marketplace](https://github.com/marketplace?type=actions).

For more information, see [Creating actions](https://docs.github.com/en/actions/creating-actions).

### Runner: GitHub-Hosted Runner vs. Self-Hosted Runner

You can run your jobs on GitHub Hosted compute or you can host your own Self Hosted runners.

The standard runners GitHub offers are:
* `ubuntu-latest`
* `windows-latest`
* `macos-latest`

There are also [Larger runners](https://docs.github.com/en/actions/using-github-hosted-runners/about-larger-runners/about-larger-runners#specifications-for-general-larger-runners) for more demanding use cases.

* [About Larger Runners](https://docs.github.com/en/actions/using-github-hosted-runners/about-larger-runners)

#### Cost

Actions running on standard GitHub-hosted runners are free for public repositories and self-hosted runners are free for all repositories.

For private repositories, GitHub charges based on a [per-minute rate](https://docs.github.com/en/billing/managing-billing-for-github-actions/about-billing-for-github-actions#per-minute-rates). The cost is simply the number of minutes your job runs multiplied by the per-minute rate.

> [!TIP]
> GitHub always rounds up the time that a job runs to the nearest minute. For example, if your job runs for 61 seconds, GitHub will charge you for 2 minutes.

GitHub Plans get a certain number of included minutes per month:

* Free: 2,000
* Team: 3,000
* Enterprise: 50,000

> [!WARNING]
> These minutes are ONLY applicable to standard runners (not larger runners).
> The above values are for `ubuntu-latest` runners.
> `windows-latest` are 2x the cost (25k free)
> `macos-latest` are 10x the cost (5k free).

![alt text](images/Screenshot%202024-08-12%20at%2010.33.53 AM.png)

### Autoscaling with self-hosted runners (ARC)

You can automatically increase or decrease the number of self-hosted runners in your environment in response to the webhook events you receive with a particular label.

* [Autoscaling with self-hosted runners](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners/autoscaling-with-self-hosted-runners)
* [actions-runner-controller](https://github.com/actions/actions-runner-controller)

<!-- ## GitHub Actions Ethos

### Essence/Opinionation

### Traceability

### Reusability (DRY)

### Ephemerality

### Extensibility: Do Anything You Want

### Easy to Get Started

## How to Author Workflow Files -->

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

You can use expressions to programmatically set environment variables in workflow files and access contexts. An expression can be any combination of literal values, references to a context, or functions. You can combine literals, context references, and functions using operators. For more information about contexts, see "Contexts."

#### Literals

You can use boolean, null, number, or string data types.

<details>
  <summary>Example of literals</summary>

```yml
env:
  myNull: ${{ null }}
  myBoolean: ${{ false }}
  myIntegerNumber: ${{ 711 }}
  myFloatNumber: ${{ -9.2 }}
  myHexNumber: ${{ 0xff }}
  myExponentialNumber: ${{ -2.99e-2 }}
  myString: Mona the Octocat
  myStringInBraces: ${{ 'It''s open source!' }}
```
</details>

#### Operators

<details>
  <summary>Example of operators</summary>

```
Operator	Description
( )	Logical grouping
[ ]	Index
.	Property de-reference
!	Not
<	Less than
<=	Less than or equal
>	Greater than
>=	Greater than or equal
==	Equal
!=	Not equal
&&	And
||	Or
```
</details>

> [!TIP]
> You can use a ternary operator `condition ? true : false` as `${{ condition && true || false }}`.

[Expressions](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions)

#### Functions

You can use functions to transform data or to perform operations.

* [contains](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#contains)
* [startswith](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#startswith)
* [endsWith](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#endswith)
* [format](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#format)
* [join](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#join)
* [toJson](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#tojson)
* [fromJson](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#fromjson)
* [hashFiles](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#hashfiles)

#### Status Check functions

* [success](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#success)
* [always](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#always)
* [cancelled](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#cancelled)
* [failure](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#failure)

## How to Trigger/Initiate Workflow Runs

You can configure your workflows to run when specific activity on GitHub happens, at a scheduled time, or when an event outside of GitHub occurs.

* [Events that trigger workflows](https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/events-that-trigger-workflows)

<!-- ### Summary of Event Grid That Triggers Workflows -->

<!-- ### Configuring Input (Activity Types): Conditionally Trigger -->

### Event: Workflow Dispatch

Workflows triggered by `workflow_dispatch` and `workflow_call` access their inputs using the inputs context.

For workflows triggered by `workflow_dispatch` inputs are available in the `github.event.inputs`. 

<details>
  <summary>Example of on.workflow_dispatch.inputs</summary>

```yml
on:
  workflow_dispatch:
    inputs:
      logLevel:
        description: 'Log level'
        required: true
        default: 'warning'
        type: choice
        options:
        - info
        - warning
        - debug
      tags:
        description: 'Test scenario tags'
        required: false
        type: boolean
      environment:
        description: 'Environment to run tests against'
        type: environment
        required: true

jobs:
  log-the-inputs:
    runs-on: ubuntu-latest
    steps:
      - run: |
          echo "Log level: $LEVEL"
          echo "Tags: $TAGS"
          echo "Environment: $ENVIRONMENT"
        env:
          LEVEL: ${{ inputs.logLevel }}
          TAGS: ${{ inputs.tags }}
          ENVIRONMENT: ${{ inputs.environment }}
```
</details>

* [`workflow_dispatch` event](https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/events-that-trigger-workflows#workflow_dispatch)
* [`inputs` context](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/contexts#inputs-context)

### Event: Workflow Call
Workflows triggered by `workflow_call` access their inputs using the `inputs` context.

<details>
  <summary>Example of on.workflow_call.outputs</summary>

```yml
on:
  workflow_call:
    # Map the workflow outputs to job outputs
    outputs:
      workflow_output1:
        description: "The first job output"
        value: ${{ jobs.my_job.outputs.job_output1 }}
      workflow_output2:
        description: "The second job output"
        value: ${{ jobs.my_job.outputs.job_output2 }}
```
</details>

* [`workflow_call` event](https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/events-that-trigger-workflows#workflow_call)
* [`inputs` context](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/contexts#inputs-context)

### Event: Workflow Run

The `workflow_run` event allows you to execute a workflow based on execution or completion of another workflow.

<details>
  <summary>Running a workflow based on the conclusion of another workflow</summary>

```yml
on:
  workflow_run:
    workflows: [Build]
    types: [completed]

jobs:
  on-success:
    runs-on: ubuntu-latest
    if: ${{ github.event.workflow_run.conclusion == 'success' }}
    steps:
      - run: echo 'The triggering workflow passed'
  on-failure:
    runs-on: ubuntu-latest
    if: ${{ github.event.workflow_run.conclusion == 'failure' }}
    steps:
      - run: echo 'The triggering workflow failed'
```
</details>

### Event: Schedule

The `schedule` event allows you to trigger a workflow at a scheduled time.

<details>
  <summary>Running a workflow on a schedule</summary>

```yml
on:
  schedule:
            ┌───────────── minute (0 - 59)
            │ ┌───────────── hour (0 - 23)
            │ │ ┌───────────── day of the month (1 - 31)
            │ │ │ ┌───────────── month (1 - 12 or JAN-DEC)
            │ │ │ │ ┌───────────── day of the week (0 - 6 or SUN-SAT)
            │ │ │ │ │
            │ │ │ │ │
            │ │ │ │ │
    - cron: * * * * *
```
</details>

### Non-Core CI/CD Use Cases

Understand that there are many ways to use GitHub Actions beyond CI/CD. For example, you can use GitHub Actions to:
* 

### Concurrency: Order of Workflow Runs Based on When Trigger Happened

GitHub Actions also allows you to control the concurrency of workflow runs, so that you can ensure that only one run, one job, or one step runs at a time in a specific context.

> [!NOTE]
> This is NOT a queueing system. If you have a lot of workflow runs that are waiting to run, they will be run in the order that they were triggered.

<details>
  <summary>Example: Concurrency groups</summary>

```yml
on:
  push:
    branches:
      - main

concurrency:
  group: ci-${{ github.ref }}
  cancel-in-progress: true
```
</details>

<details>
  <summary>Example: Using concurrency to cancel any in-progress job or run</summary>

```yml
concurrency:
  group: ${{ github.ref }}
  cancel-in-progress: true
```
</details>

You can make the concurrency group as specific as you want. For example, you could use the branch name, the branch name and the event type, or the branch name and the event type and the workflow name.

### Re-running Workflows / Retries

You can re-run a workflow run from the Actions UI. This is useful if you want to re-run a failed workflow run, or if you want to re-run a successful workflow run.

Retrying a job programmatically is not officially supported but can be achieved using something like a [marketplace action](https://github.com/marketplace?query=retry)

## How to Structure/Manage Jobs in the Workflow

### Parallelization of Jobs

By default all jobs in a workflow run in parallel. You can control the order of jobs by specifying dependencies.

### Matrices

A matrix strategy is a great way to run the same job multiple times with different inputs. This is useful if you want to run your tests on multiple versions of a language, or if you want to run your tests on multiple operating systems.

> [!NOTE]
> The maximum number of jobs that can be used in a matrix strategy is 256.

<details>
  <summary>Example of a matrix strategy</summary>

```yml
jobs:
  example_matrix:
    strategy:
      matrix:
        version: [10, 12, 14]
        os: [ubuntu-latest, windows-latest]
```
</details>

* [Using a matrix for your jobs](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/using-a-matrix-for-your-jobs)

### Ordering Jobs

You can define the order of the jobs using the `needs` keyword. This is useful if you want to run a job that depends on the output of another job.

<details>
  <summary>Example of linking jobs</summary>

```yml
jobs:
  job1:
  job2:
    needs: job1
  job3:
    needs: [job1, job2]
    steps:
      - run: echo ${{ needs.job1.outputs.myOutput }}
```
</details>

* [Defining prerequisite jobs](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/using-jobs-in-a-workflow#defining-prerequisite-jobs)

### Job Timeouts

You can define a timeout for a job, and if the job takes longer than the timeout to run, the job will be cancelled.

The default timeout for a job is 6 hours or 360 minutes.

> [!NOTE]
> The `GITHUB_TOKEN` expires after the job finishes or 24 hours. This is a limiting factor for SHRs.

* [`jobs.<job_id>.steps[*].timeout-minutes`](https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#jobsjob_idstepstimeout-minutes)
* [`jobs.<job_id>.timeout-minutes`](https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#jobsjob_idtimeout-minutes)

### Sharing Artifacts Between Jobs

The [actions/upload-artifact](https://github.com/actions/upload-artifact) and [download-artifact](https://github.com/actions/download-artifact) actions let you share data between jobs. You do have to explicitly do this on a per-job basis.

<details>
  <summary>Example of sharing artifacts between jobs</summary>

```yml
name: Share data between jobs

on: [push]

jobs:
  job_1:
    name: Add 3 and 7
    runs-on: ubuntu-latest
    steps:
      - shell: bash
        run: |
          expr 3 + 7 > math-homework.txt
      - name: Upload math result for job 1
        uses: actions/upload-artifact@v4
        with:
          name: homework_pre
          path: math-homework.txt

  job_2:
    name: Multiply by 9
    needs: job_1
    runs-on: windows-latest
    steps:
      - name: Download math result for job 1
        uses: actions/download-artifact@v4
        with:
          name: homework_pre
      - shell: bash
        run: |
          value=`cat math-homework.txt`
          expr $value \* 9 > math-homework.txt
      - name: Upload math result for job 2
        uses: actions/upload-artifact@v4
        with:
          name: homework_final
          path: math-homework.txt

  job_3:
    name: Display results
    needs: job_2
    runs-on: macOS-latest
    steps:
      - name: Download math result for job 2
        uses: actions/download-artifact@v4
        with:
          name: homework_final
      - name: Print the final result
        shell: bash
        run: |
          value=`cat math-homework.txt`
          echo The result is $value
```
</details>

* [Passing data between jobs in a workflow](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/storing-workflow-data-as-artifacts#passing-data-between-jobs-in-a-workflow)

### Running Jobs in Containers / Service Containers

Running in a container will not always be faster than running on a GHR. The time it takes to download the container image and start the container can be longer than the time it takes to start a job on a GHR.

### Containers

Use `jobs.<job_id>.container` to create a container to run any steps in a job that don't already specify a container.

<details>
  <summary>Example of running a job within a container</summary>

```yml
name: CI
on:
  push:
    branches: [ main ]
jobs:
  container-test-job:
    runs-on: ubuntu-latest
    container: 
      image: node:18
      env:
        NODE_ENV: development
      ports:
        - 80
      volumes:
        - my_docker_volume:/volume_mount
      options: --cpus 1
    steps:
      - name: Check for dockerenv file
        run: (ls /.dockerenv && echo Found dockerenv) || (echo No dockerenv)
```
</details>

> [!TIP]
> You can omit the `image` keyword and use the short version `container: node:18` if you don't need to specify parameters.

* [Running jobs in a container](https://docs.github.com/en/actions/writing-workflows/choosing-where-your-workflow-runs/running-jobs-in-a-container)

#### Service Containers

Service containers let you run a container parallel to your job. This can be helpful if your job needs to talk to a database, for example.

* [About service containers](https://docs.github.com/en/actions/use-cases-and-examples/using-containerized-services/about-service-containers)

<details>
  <summary>Example of using a service container</summary>

```yml
name: Redis container example
on: push

jobs:
  # Label of the container job
  container-job:
    # Containers must run in Linux based operating systems
    runs-on: ubuntu-latest
    # Docker Hub image that `container-job` executes in
    container: node:16-bullseye

    # Service containers to run with `container-job`
    services:
      # Label used to access the service container
      redis:
        # Docker Hub image
        image: redis
```
</details>

#### Authenticating with a Container Registry

Sometimes you will need to authenticate with a container registry to pull an image. You can use the `credentials` keyword to do this.

<details>
  <summary>Example of authenticating with a container registry</summary>

```yml
jobs:
  build:
    services:
      redis:
        # Docker Hub image
        image: redis
        ports:
          - 6379:6379
        credentials:
          username: ${{ secrets.dockerhub_username }}
          password: ${{ secrets.dockerhub_password }}
      db:
        # Private registry image
        image:  ghcr.io/octocat/testdb:latest
        credentials:
          username: ${{ github.repository_owner }}
          password: ${{ secrets.ghcr_password }}
```
</details>

* [Authenticating with image registries](https://docs.github.com/en/actions/use-cases-and-examples/using-containerized-services/about-service-containers#authenticating-with-image-registries)

### Environments: Controls How/When a Job is Run Based on Protection Rules Set, Limits Branches, Scopes Secrets

You can create environments and secure those environments with deployment protection rules. A job that references an environment must follow any protection rules for the environment before running or accessing the environment's secrets.

Scoping secrets to an environment is very powerful because of the controls it gives you. You can limit which branches can access the secrets, and you can leverage the environment protection rules to control when a job can access the secrets.

#### Environment Protection Rules

Deployment protection rules require specific conditions to pass before a job referencing the environment can proceed.

##### Required Reviewers

You can require that specific individuals or teams review a pull request before a job can proceed.

##### Wait timer

You can delay a job for a specific amount of time before it can proceed.

##### Branch restrictions

You can restrict which branches or tags can access the environment.

##### Admin bypass

You can allow or disallow repository administrators to bypass the protection rules.

##### Custom deployment protection rules

You can create custom deployment protection rules to gate deployments with third-party services.

* [Deployment Protection Rules](https://docs.github.com/en/actions/managing-workflow-runs-and-deployments/managing-deployments/managing-environments-for-deployment#deployment-protection-rules)
* [Configuring custom deployment protection rules](https://docs.github.com/en/actions/managing-workflow-runs-and-deployments/managing-deployments/configuring-custom-deployment-protection-rules)
* [Environment Secrets](https://docs.github.com/en/actions/managing-workflow-runs-and-deployments/managing-deployments/managing-environments-for-deployment#environment-secrets)

### Conditional Jobs/Steps

You can use the `if` keyword to conditionally run a job or step.

```yml
if: ${{ ! startsWith(github.ref, 'refs/tags/') }}
```

<details>
  <summary>Example of conditional jobs</summary>

```yml
name: example-workflow
on: [push]
jobs:
  production-deploy:
    if: github.repository == 'octo-org/octo-repo-prod'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '14'
      - run: npm install -g bats
```
</details>

* [Using conditions to control job execution](https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/using-conditions-to-control-job-execution)

### Permissions for Jobs

There is a default token called `GITHUB_TOKEN` which by default has the permissions defined in your repositories Actions settings.

It's a good idea to limit permissions as much as possible by being explicit.

<details>
  <summary>Example of limiting permissions</summary>

```yml
jobs:
  stale:
    runs-on: ubuntu-latest

    permissions:
      issues: write
      pull-requests: write

    steps:
      - uses: actions/stale@v5
```
</details>

#### GitHub Apps

Using [actions/create-github-app-token](https://github.com/actions/create-github-app-token) you can get a token for a GitHub App. This is better than using a PAT because you get more control and you don't need to consume a license.

<details>
  <summary>Example of using a GitHub App token</summary>

```yml
name: Run tests on staging
on:
  push:
    branches:
      - main

jobs:
  hello-world:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/create-github-app-token@v1
        id: app-token
        with:
          app-id: ${{ vars.APP_ID }}
          private-key: ${{ secrets.PRIVATE_KEY }}
      - uses: ./actions/staging-tests
        with:
          token: ${{ steps.app-token.outputs.token }}
```
</details>

* [Assigning permissions to jobs](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/assigning-permissions-to-jobs)
* [Automatic token authentication](https://docs.github.com/en/actions/security-for-github-actions/security-guides/automatic-token-authentication)

## How to Use and Create Actions (Marketplace)

### What is an Action?

Actions are the building blocks that power your workflow. A workflow can contain one or more actions, either as individual steps or as part of an action group. An action is a reusable unit of code that can be used in multiple workflows. You can create your own actions, use actions created by the GitHub community, or use a combination of both.

* [GitHub Actions Marketplace](https://github.com/marketplace?type=actions)

### Types of Actions

Javascript actions are the most popular and easiest to get started with, Docker containers package the environment with the GitHub Actions code, and Composite actions are a way to reuse actions in a more modular way.

There are three types of custom actions:
* [JavaScript](https://docs.github.com/en/actions/creating-actions/creating-a-javascript-action)
* [Docker](https://docs.github.com/en/actions/creating-actions/creating-a-docker-container-action) (Not available on macOS or Windows runners)
* [Composite](https://docs.github.com/en/actions/creating-actions/creating-a-composite-run-steps-action)
* [About custom actions](https://docs.github.com/en/actions/creating-actions/about-custom-actions)

### Securing Usage of Actions

* [Security hardening for GitHub Actions](https://docs.github.com/en/actions/security-for-github-actions/security-guides/security-hardening-for-github-actions)

### Creating Your Own Actions

You can create your own actions to use in your workflows. This is a great way to encapsulate logic that you want to reuse across multiple workflows.

* [Creating actions](https://docs.github.com/en/actions/creating-actions)

### Cool Actions to Look Out For: github-script, Anything by GitHub, Major Cloud Providers, Terraform, Docker

Here are some popular actions to get you started:

* [GitHub Script](https://github.com/actions/github-script)
* [Awesome Actions](https://github.com/sdras/awesome-actions#readme)
* [GitHub Authored Actions](https://github.com/marketplace?query=publisher%3Aactions)
* [Azure Actions](https://github.com/marketplace?query=publisher%3Aazure)
* [AWS Actions](https://github.com/marketplace?query=publisher%3Aaws-actions)
* [GCP Actions](https://github.com/marketplace?query=publisher%3Agoogle-github-actions)
* [Build and Push Docker Images](https://github.com/marketplace/actions/build-and-push-docker-images)

## How to Organize, Share, and Scale Workflows

One of the most powerful features of GitHub Actions is the ability to share workflows across repositories. This is useful if you have a common workflow that you want to use in multiple repositories.

### Reusable Workflows

These are reusable jobs. They are a great way to share common logic across multiple workflows or just to organize your workflow into smaller, more manageable pieces.

#### Why?

* Easier to maintain
* Create workflows more quickly
* Avoid duplication. DRY(don't repeat yourself).
* Build consistently across multiple, dozens, or even hundreds of repositories
* Require specific workflows for specific deployments
* Promotes best practices
* Abstract away complexity

#### What can they do

* Can have inputs and outputs
* Can be nested 4 levels deep
* Only 20 unique reusable workflows can be in a single workflow
* Environment variables are not propagated to the reusable workflow
* Secrets are scoped to the caller workflow
* Secrets need to be passed to the reusable workflow

<details>
  <summary>Example of a reusable workflow</summary>

##### Defining the workflow (reusable-called.yml)

```yml
on:
  workflow_call:
    inputs:
      username:
        default: ${{ github.actor }}
        required: false
        type: string

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Run a one-line script
        run: echo Hello, ${{ inputs.username }}!
```

##### Using the workflow (caller.yml)

```yml
jobs:
  build:
    uses: ./.github/workflows/reusable-called.yml
    with:
      username: ${{ github.actor }}
```

</details>

* [Reusing workflows](https://docs.github.com/en/actions/sharing-automations/reusing-workflows)
* [Limitations](https://docs.github.com/en/actions/sharing-automations/reusing-workflows#limitations)

### Composite Actions

These are reusable steps. Use a composite action to combine(re-use) multiple steps.

> [!TIP]
> These are far less limited than reusable workflows. Consider using composite actions over reusable workflows to start.

<details>
  <summary>Example of a composite action</summary>

##### Defining the action (hello-world-composite-action.yml)

```yml
name: 'Hello World'
description: 'Greet someone'
inputs:
  who-to-greet:  # id of input
    description: 'Who to greet'
    required: true
    default: 'World'
outputs:
  random-number:
    description: "Random number"
    value: ${{ steps.random-number-generator.outputs.random-number }}
runs:
  using: "composite"
  steps:
    - name: Set Greeting
      run: echo "Hello $INPUT_WHO_TO_GREET."
      shell: bash
      env:
        INPUT_WHO_TO_GREET: ${{ inputs.who-to-greet }}

    - name: Random Number Generator
      id: random-number-generator
      run: echo "random-number=$(echo $RANDOM)" >> $GITHUB_OUTPUT
      shell: bash

    - name: Set GitHub Path
      run: echo "$GITHUB_ACTION_PATH" >> $GITHUB_PATH
      shell: bash
      env:
        GITHUB_ACTION_PATH: ${{ github.action_path }}

    - name: Run goodbye.sh
      run: goodbye.sh
      shell: bash
```

##### Using the action (caller.yml)

```yml
on: [push]

jobs:
  hello_world_job:
    runs-on: ubuntu-latest
    name: A job to say hello
    steps:
      - uses: actions/checkout@v4
      - id: foo
        uses: OWNER/hello-world-composite-action@TAG
        with:
          who-to-greet: 'Mona the Octocat'
      - run: echo random-number "$RANDOM_NUMBER"
        shell: bash
        env:
          RANDOM_NUMBER: ${{ steps.foo.outputs.random-number }}
```

</details>

* [Creating a composite action](https://docs.github.com/en/actions/sharing-automations/creating-actions/creating-a-composite-action)

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
