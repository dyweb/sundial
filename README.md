
# Sundial

[![Build Status](https://travis-ci.org/dyweb/sundial.svg?branch=master)](https://travis-ci.org/dyweb/sundial)
[![Coverage Status](https://coveralls.io/repos/github/dyweb/sundial/badge.svg)](https://coveralls.io/github/dyweb/sundial)

Open Source WakaTime Server

## Getting Started

### Main Dependencies

- [caicloud/nirvana](https://github.com/caicloud/nirvana) Golang Restful API Framework for Productivity
- [jinzhu/gorm](https://github.com/jinzhu/gorm) The fantastic ORM library for Golang, aims to be developer friendly

### Building

```bash
dep ensure -s
make
```

### Running

```bash
./bin/sundial
```

### Debugging

First, we need to run the server, then we could use wakatime-cli to create a heartbeat:

```
python3 ./wakatime/wakatime/cli.py --file /home/ist/go/src/github.com/dyweb/sundial/pkg/apis/v1/descriptors/heartbeats.go  --alternate-project sundial --write --api-url http://localhost:8080/apis/v1/users/current/heartbeats --verbose
```

We could get the log from `$HOME/.wakatime.log`.

## Versioning

### [v1](./pkg/apis/v1)

The goal of v1 is to provide the MVP of Sundial, which supports projects, users and durations API.

## Development Document

[devel.md](./docs/devel.md)

## Authors

Please see [AUTHORS.md](./AUTHORS.md).
