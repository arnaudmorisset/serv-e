# serv-e

**serv-e** is a lightweight HTTP server that will log every incoming request.
You can use it to ease the debug of webhooks and callbacks in other software projects.

## Install

### Mac

You can grab the latest available binary from [Github's releases page](https://github.com/arnaudmorisset/serv-e/releases/).

> Homebrew formulae will be available soon!

### Docker

```shell
docker pull arnaudmorisset/serv-e
```

### Build from source

**serv-e** is written using [Go](https://go.dev) and does not rely on any dependencies other than the language runtime, which can be installed using [asdf](https://asdf-vm.com/#/).

```shell
# Checkout the source
git clone git@github.com:arnaudmorisset/serv-e.git && cd serv-e

# Install Go runtime (asdf required)
asdf install

# Build the project
go build .
```

## Launch the tests

The tests are in the special folder `tests`, and you have to specify it running them:

```shell
go test -v ./tests
```

## Running serv-e

Running the command without any arguments will spawn a server listening on `localhost:80`.

`localhost:80/` should be used as a target for your webhooks/callbacks.

`localhost:80/records` should be used to display the list of all recorded requests.

**Keep in mind that records are saved in-memory. Stopping the server will erase all logs.**
