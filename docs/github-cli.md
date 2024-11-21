---
layout: default
title: GitHub CLI
nav_order: 4
---

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
