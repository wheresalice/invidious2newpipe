# invidious2newpipe

Takes an opml export from Invidious and converts it into a JSON file for use with NewPipe

This lets you export your Invidious subscriptions into NewPipe

## Usage

Download the appropriate binary for your system and **Export subscriptions as OPML (for NewPipe & FreeTube)**

Then run:

```shell
./invidious2newpipe cli ~/Downloads/subscription_manager
```

Alternatively you can run this as a web service to make it easier to migrate between devices

```shell
SUBS_DIR=/tmp ./invidious2newpipe web
```