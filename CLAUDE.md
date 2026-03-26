# CLAUDE.md — YOUR_PROJECT

A brief description of your CLI tool.

## Build & Test

```bash
make build           # Build binary → ./bin/mycli
make test            # Run unit tests
make test-unit       # go test ./... -v -race -cover
make cover           # Generate coverage report
make cover-html      # Open coverage in browser
make fmt             # go fmt
make vet             # go vet
make install         # Install to /usr/local/bin
```

## Project Structure

```
cmd/
  main.go              # Entry point
  cli/
    root.go            # Cobra root command + global flags
    version.go         # Version subcommand (ldflags)
    greet.go           # Example subcommand
internal/
  example/
    example.go         # Example business logic
    example_test.go    # Example tests
scripts/
  create-pr.sh         # Auto-generate PR body from commits
  install.sh           # curl installer
```

## Workflow After Code Changes

1. **Tests first** — Write or update tests. Run `make test` and ensure all tests pass.
2. **Documentation second** — Update README.md and CLAUDE.md if needed.
