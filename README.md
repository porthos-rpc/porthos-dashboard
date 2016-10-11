# Porthos Dashboard

A dashboard application for operating a porthos-rpc cluster.

## Status

[![Build Status](https://travis-ci.org/porthos-rpc/porthos-dashboard.svg?branch=master)](https://travis-ci.org/porthos-rpc/porthos-dashboard)

## Requirements

- RabbitMQ (any other AMQP 0.9.1 broker)
- SQLite

## Build

First of all, you need to download `govendor`:

```sh
go get -u github.com/kardianos/govendor
```

Fetch all go dependencies:

```sh
govendor sync
```

Then build the app:

```sh
go build
```

## Running the app

After building the executable, you may run the dashboard as following:

```sh
./porthos-dashboard -bind 8080 -broker amqp:// -db dash.db
```

## Local environment

To speed things up we provide a docker environment. Just run `docker-compose up` then open `http://localhost:8080/`.

## Example

<img src="https://raw.githubusercontent.com/porthos-rpc/porthos-dashboard/master/screenshot.png">
