# User Service

This is the User Service

Generated with

```
micro new gitlab.com/bimoyong/go-user --namespace=go --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.srv.user
- Type: srv
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./bin/srv.user
```

Build a docker image
```
make docker
```