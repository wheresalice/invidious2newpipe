# invidious2newpipe

Takes an opml export from Invidious and converts it into a JSON file for use with NewPipe

This lets you export your Invidious subscriptions into NewPipe

## Usage

```shell
# From source code:
go run invidious2newpipe.go > export.json

# Specifying a custom file
go run invidious2newpipe.go ~/Downloads/subscription_manager > export.json

# Building and running a binary file
go build invidious2newpipe.go
./invidious2newpipe.go
```