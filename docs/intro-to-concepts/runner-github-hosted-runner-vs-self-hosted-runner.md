---
layout: default
title: 'Runner: GitHub-Hosted Runner vs. Self-Hosted Runner'
parent: Intro to Concepts
has_children: false
nav_order: 2
---

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
