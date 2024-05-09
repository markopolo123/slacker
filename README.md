# slacker

Set Slack statuses quicker than a cheetah on a caffeine rush

# Purpose

This is a CLI utility to set/clear your status in slack. It comes with 3 default statuses and of course you can make your own.

# Building

1. clone this respository
1. `flox activate`
1. `go mod tidy`
1. `go build .`

:warning: Flox is an enviornment and package manager in one. It can bring dependencies and allow for reproducable development enviornments. Get it at https://flox.dev. 

Alternatively, if you have a go environment already working, just clone and build.

# Usage

    $ slacker init
    $ slacker help


## Examples

List available statuses.

    $ slacker show
    
    ┏━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┓
    ┃   Shortcut    ┃      Text      ┃     Emoji     ┃   Duration    ┃
    ┣━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━┫
    ┃     desk      ┃   At my Desk   ┃      🖥️       ┃      1h       ┃
    ┃     lunch     ┃ Gone for lunch ┃      🍕       ┃      60m      ┃
    ┃      dnd      ┃ Do Not Disturb ┃      ⚠️        ┃      60m      ┃
    ┃               ┃                ┃               ┃               ┃
    ┗━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━┛


Set a status

    $ slacker dnd
    
    Setting status: dnd
    Slack status updated successfully!


# Token

You'll need a personal slack token (set as `SLACK_TOKEN`). A token can be created using
https://api.slack.com and creating an application. This is focused on "OAuth &
Scopes." The other paramters for a slack application are unimportant. 

If you want to create it from a manifest, it will be something like this.

```json
{
    "display_information": {
        "name": "slacker applications for $user"
    },
    "oauth_config": {
        "scopes": {
            "user": [
                "users.profile:read",
                "users.profile:write",
                "users:read"
            ]
        }
    },
    "settings": {
        "org_deploy_enabled": false,
        "socket_mode_enabled": false,
        "token_rotation_enabled": false
    }
}
```

# Configuration

Configuration/statuses are set in `$HOME/.config/slacker/config.toml` after running `slacker init`. 
