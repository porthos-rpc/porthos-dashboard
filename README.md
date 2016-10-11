# Porthos Dashboard

A dashboard application for operating a porthos-rpc cluster.

## Status

[![Build Status](https://travis-ci.org/porthos-rpc/porthos-dashboard.svg?branch=master)](https://travis-ci.org/porthos-rpc/porthos-dashboard)

## Build and run

```sh
go build
./porthos-dashboard -bind 8080 -broker amqp:// -db dash.db
```

## Example

<img src="https://raw.githubusercontent.com/porthos-rpc/porthos-dashboard/master/screenshot.png">
