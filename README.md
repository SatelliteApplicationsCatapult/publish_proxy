# Publish proxy

A really simple proxy to handle routing of requests from cube query to the publish server.

This is required in our environment because the internal cluster subnet of our kubernetes cluster
is the same as the publishing servers ip address.

You almost certainly do not need to use this. 

## prerequisites

Go 1.18+ https://go.dev/dl/

## build

```bash
go build
```

## Running

```bash
./publish_proxy --host "<publish server address>" --listen ":8080"
```