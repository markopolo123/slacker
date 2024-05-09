---
title: Uninstall Slacker CLI
weight: 1
description: How to remove the Slacker CLI from your system
---

In this tutorial, you learn how to remove the Slacker CLI on MacOS.

> Slacker has no daemon or long running processes

* [ ] Remove the binary from your system
* [ ] Remove the configuration file

## Step 1: Remove the binary

1. Open a terminal window.
2. Run the following command to remove the binary:

```bash
rm $(which slacker)
```

## Step 2: Remove the configuration file

1. Run the following command to remove the configuration file:

> This command will remove the configuration file and all the data associated with it.

```bash
rm -rf ~/.config/slacker
```
