# go-loader

Boilerplate loaders for Go projects: environment variables, TLS credentials, Google Cloud credentials and tokens, and files. Requires Go 1.25 (per `go.mod`).

## Installation

```bash
go get github.com/ralvarezdev/go-loader
```

Direct dependencies: `golang.org/x/oauth2`, `google.golang.org/api`, `google.golang.org/grpc`.

## Packages

- **`env`** — `Loader` interface (`LoadVariable`, `LoadDurationVariable`, `LoadSecondsVariable`, `LoadIntVariable`) and `NewDefaultLoader(loadFn, logger)`. `loadFn` is an optional function run first (for example one that loads a `.env` file). Missing variables return an error naming the key; successful loads log the key only.
- **`filesystem`** — `ReadFile`, `OpenFile`, `CloseFile`, `ReadCSVFile(file, readHeaders)`, `GetCurrentDirectory`, `GetExecutableGoModPath`.
- **`http/tls`** — `LoadTLSCredentials(pemServerCAPath)` and `LoadSystemCredentials()`, returning gRPC `TransportCredentials`.
- **`cloud/gcloud`** — `LoadGoogleCredentials(ctx)` (application default credentials) and `LoadServiceAccountCredentials(ctx, url, credentials)`, returning an OAuth token source.

## Usage

```go
loader, err := goloaderenv.NewDefaultLoader(nil, slog.Default())

var host string
var timeout time.Duration
var port int
err = loader.LoadVariable("SERVER_HOST", &host)
err = loader.LoadDurationVariable("SERVER_TIMEOUT", &timeout)
err = loader.LoadIntVariable("SERVER_PORT", &port)
```

`SERVER_*` are example names, not variables defined by this library.

## Development

```bash
go build ./...
go vet ./...
```

There are no tests.

## License

GNU General Public License v3.0 (see [LICENSE](LICENSE)).
