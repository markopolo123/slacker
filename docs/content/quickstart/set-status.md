---
title: Set Your Slack Status
weight: 2
description: Quickstart Usage
---

## Set Your Slack Status

* [ ] Set some template statuses
* [ ] set one as your slack status using slacker

```bash
slacker init
```

This will create a config.toml in `$HOME/.config/slacker/` with some example statuses.

Once you've set your user ID and user auth token environment variables, you may set your status using the following command:

```bash
slacker lunch
```

To see available statuses run `slacker show` :

```bash
slacker show
┏━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┓
┃   Shortcut    ┃      Text      ┃     Emoji     ┃   Duration    ┃
┣━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━┫
┃     desk      ┃   At my Desk   ┃      🖥️       ┃      1h       ┃
┃     lunch     ┃ Gone for lunch ┃      🍕       ┃      60m      ┃
┃      dnd      ┃ Do Not Disturb ┃      ⚠️        ┃      60m      ┃
┃               ┃                ┃               ┃               ┃
┗━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━┛
```

You may clear your status using the following command:

```bash
slacker clear
```

Running `slacker -h` will show a full help for the CLI.
