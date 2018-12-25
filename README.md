
# Sundial

Open Source WakaTime Server

## Getting Started

### Prerequisites

- golang
- dep

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
python3 ./wakatime/wakatime/cli.py --file /home/test/go/src/github.com/dyweb/sundial/pkg/apis/v1/descriptors/heartbeats.go --plugin "\"vscode/1.29.1 vscode-wakatime/1.2.5\"" --alternate-project sundial --write --api-url http://localhost:8080/apis/v1/users/current/heartbeats
```

We could get the log from `$HOME/.wakatime.log`.

## Versioning

<!-- Place versions of this project and write comments for every version -->

TODO

## Development Document

[devel.md](./docs/devel.md)

## Authors

Please see [AUTHORS.md](./AUTHORS.md).

## License

TODO
