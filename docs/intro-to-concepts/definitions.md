---
layout: default
title: Definitions
parent: Intro to Concepts
has_children: false
nav_order: 1
---

# Definitions

Some basic definitions to get us started...

![overview-actions-simple](/images/overview-actions-simple.png)

## [Workflow](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#workflows)

A workflow is a configurable automated process that will run one or more jobs. Workflows are defined by a YAML file checked in to your repository in the `.github/workflows` directory. A repository can have multiple workflows, each of which can perform a different set of tasks.

## [Events](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#events)

An event is a specific activity in a repository that triggers a workflow run. It could be triggered by an event in your repository, or they can be triggered manually, or at a defined schedule.

## [Jobs](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#jobs)

A job is a set of steps in a workflow that is executed on the same runner. Each step is either a shell script that will be executed, or an action that will be run. Steps are executed in order and are dependent on each other. Since each step is executed on the same runner, you can share data from one step to another.

## [Steps / Actions](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#actions)

A step can be a script that will be executed or a GitHub action.

## [Runners](https://docs.github.com/en/actions/about-github-actions/understanding-github-actions#runners)

A runner is a server that runs your workflows when they're triggered. Each runner can run a single job at a time.

* GHR: GitHub-Hosted Runner
* SHR: Self-Hosted Runner
