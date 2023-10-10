# kyasshu: simple, fast, and secure s3 proxy

## Development

1. Clone the repository
2. Install dependencies: `go mod download`
3. Copy `config.example.yaml` to `config.yaml` and edit it
4. Run the server: `go run . serve`
5. Start caddy proxy: `caddy start` in the root of the repository