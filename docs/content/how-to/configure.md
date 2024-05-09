---
title: Configure Slacker
weight: 2
description: A tutorial for configuring Slack API Authentication
---

In this tutorial, you learn how to configure `Slacker` to work with the Slack API.

## Prerequisites

* [ ] `Slacker` CLI installed
* [ ] A Slack account

## Steps

1. [ ] Create a new Slack App
2. [ ] Generate a Slack API token
3. [ ] Configure `Slacker` with the Slack API token
4. [ ] Set the Slack User ID as an environment variable
5. [ ] Verify the configuration

## Step 1: Create a new Slack App

1. Open the [Slack API](https://api.slack.com/apps) page.
2. Click the `Create New App` button.
3. Enter a name for your app and select the workspace where you want to install the app.
4. Click `Create App`.
5. In the left-hand navigation, click `OAuth & Permissions`.
6. Under `Scopes`, add the following scopes:
   * users:read
   * users:write
   * users.profile:read

> **Note**: If you get stuck check the quick start guide on the [Slack Docs](https://api.slack.com/tutorials/tracks/getting-a-token)

## Step 2: Generate a Slack API token

1. In the left-hand navigation, click `OAuth & Permissions`.
2. Grab the User OAuth Token
3. in your terminal set the token as an environment variable:

```bash
   export SLACK_TOKEN=xoxp-rest-of-the-token
   ```

{{<img src="how-to/images/token.png" alt="tokenisation" lightbox="true">}}

## Step 3: Set your User ID as an environment variable

1. Navigate to the bottom left corner and click on your profile picture. 
2. From there, select 'Profile' and then click on the three hamburger dots on the right. 
3. Your user ID is available in this menu labelled as 'Member ID'.
4. Export that ID as an environment variable:

```bash
export SLACK_USER_ID=U12345678
```

{{<img src="how-to/images/user-id.png" alt="user ID" lightbox="true">}}

That's it! You've successfully configured `Slacker` to work with the Slack API. You can now use `Slacker` to interact with the Slack API:

```bash
slacker get
```
