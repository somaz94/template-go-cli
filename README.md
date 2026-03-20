# template-go-cli

A GitHub template repository for building Go CLI tools with [Cobra](https://github.com/spf13/cobra), [GoReleaser](https://goreleaser.com/), and automated CI/CD workflows.

<br/>

## What's Included

| Category | Files | Description |
|----------|-------|-------------|
| **CLI** | `cmd/main.go`, `cmd/cli/` | Cobra-based CLI with root, version, and example subcommand |
| **Business Logic** | `internal/example/` | Example package with tests (table-driven) |
| **Build** | `Makefile` | build, test, cover, fmt, vet, install, branch, pr |
| **Release** | `.goreleaser.yml` | Multi-platform builds + Homebrew tap + Scoop bucket |
| **CI/CD** | `.github/workflows/` | CI, release, changelog, contributors, dependabot, stale issues |
| **Scripts** | `scripts/` | PR auto-generator (`create-pr.sh`), curl installer (`install.sh`) |
| **Config** | `.github/dependabot.yml` | Weekly dependency updates (Actions + Go modules) |
| **Docs** | `CLAUDE.md` | Project guidelines for AI-assisted development |

<br/>

## Quick Start

<br/>

### 1. Create from Template

Click **"Use this template"** on GitHub, or:

```bash
gh repo create my-new-cli --template somaz94/template-go-cli --public --clone
cd my-new-cli
```

<br/>

### 2. Replace Placeholders

Find and replace the following placeholders across all files:

| Placeholder | Replace With | Example |
|-------------|-------------|---------|
| `YOUR_USERNAME` | Your GitHub username | `somaz94` |
| `YOUR_PROJECT` | Your repository name | `my-awesome-cli` |
| `YOUR_GITLAB_GROUP` | Your GitLab group (for mirror) | `backup6695808` |
| `mycli` | Your binary name | `my-awesome-cli` |

Quick replace:

```bash
# macOS
find . -type f -not -path './.git/*' -exec sed -i '' \
  -e 's/YOUR_USERNAME/somaz94/g' \
  -e 's/YOUR_PROJECT/my-awesome-cli/g' \
  -e 's/YOUR_GITLAB_GROUP/backup6695808/g' \
  -e 's/mycli/my-awesome-cli/g' {} +

# Linux
find . -type f -not -path './.git/*' -exec sed -i \
  -e 's/YOUR_USERNAME/somaz94/g' \
  -e 's/YOUR_PROJECT/my-awesome-cli/g' \
  -e 's/YOUR_GITLAB_GROUP/backup6695808/g' \
  -e 's/mycli/my-awesome-cli/g' {} +
```

<br/>

### 3. Initialize Module

```bash
rm go.sum
go mod init github.com/YOUR_USERNAME/YOUR_PROJECT
go mod tidy
```

<br/>

### 4. Build & Test

```bash
make build    # → ./bin/mycli
make test     # Run unit tests
```

<br/>

## Project Structure

```
.
├── cmd/
│   ├── main.go                  # Entry point
│   └── cli/
│       ├── root.go              # Root command with global flags
│       ├── version.go           # Version subcommand (ldflags)
│       └── greet.go             # Example subcommand (remove/replace)
├── internal/
│   └── example/
│       ├── example.go           # Example business logic (remove/replace)
│       └── example_test.go      # Example tests (remove/replace)
├── scripts/
│   ├── create-pr.sh             # Auto-generate PR body from commits
│   └── install.sh               # curl installer for releases
├── .github/
│   ├── workflows/
│   │   ├── ci.yml               # CI: test, vet, fmt, build
│   │   ├── release.yml          # GoReleaser on tag push
│   │   ├── changelog-generator.yml
│   │   ├── contributors.yml
│   │   ├── dependabot-auto-merge.yml
│   │   ├── stale-issues.yml
│   │   ├── issue-greeting.yml
│   │   └── gitlab-mirror.yml
│   ├── dependabot.yml
│   └── release.yml              # Release note categories
├── .goreleaser.yml
├── .gitignore
├── Makefile
├── CLAUDE.md
├── go.mod
└── README.md
```

<br/>

## Makefile Targets

```bash
make help            # Show all targets
make build           # Build binary → ./bin/mycli
make test            # Run unit tests with coverage
make cover           # Generate coverage report
make cover-html      # Open coverage in browser
make fmt             # Format code
make vet             # Run go vet
make install         # Install to /usr/local/bin
make clean           # Remove build artifacts
make branch name=x   # Create feature branch feat/x
make pr title="..."  # Test → push → create PR (auto-generated body)
```

<br/>

## CI/CD Workflows

| Workflow | Trigger | Description |
|----------|---------|-------------|
| `ci.yml` | push (main), PR, dispatch | Unit tests → Build → Version verify |
| `release.yml` | tag push `v*` | GoReleaser (binaries + Homebrew + Scoop) |
| `changelog-generator.yml` | after release, PR merge | Auto-generate CHANGELOG.md |
| `contributors.yml` | after changelog | Auto-generate CONTRIBUTORS.md |
| `dependabot-auto-merge.yml` | dependabot PR | Auto-merge minor/patch updates |
| `stale-issues.yml` | daily cron | Auto-close stale issues (30d + 7d) |
| `issue-greeting.yml` | issue opened | Welcome message |
| `gitlab-mirror.yml` | push to main | Mirror to GitLab |

<br/>

### Workflow Chain

```
tag push v* → Create release (GoReleaser)
                └→ Generate changelog
                      └→ Generate Contributors
```

<br/>

## GitHub Secrets Required

| Secret | Usage |
|--------|-------|
| `PAT_TOKEN` | Release, contributors (cross-repo access) |
| `GITLAB_TOKEN` | GitLab mirror (optional) |

> `GITHUB_TOKEN` is automatically provided by GitHub Actions.

<br/>

## Development Workflow

```bash
# 1. Create feature branch
make branch name=add-feature

# 2. Develop with tests
#    - Add code in internal/
#    - Add CLI command in cmd/cli/
#    - Write tests

# 3. Verify
make test
make build

# 4. Create PR
make pr title="feat: add new feature"
```

<br/>

## Release

```bash
git tag v0.1.0
git push origin v0.1.0
```

GoReleaser automatically:
- Builds binaries for linux/darwin/windows (amd64/arm64)
- Creates GitHub release with checksums
- Publishes to Homebrew tap and Scoop bucket

<br/>

## Conventions

- **Commits**: [Conventional Commits](https://www.conventionalcommits.org/) (`feat:`, `fix:`, `docs:`, `refactor:`, `test:`, `ci:`, `chore:`)
- **Branches**: `feat/name`, `fix/name`
- **paths-ignore**: CI skips `.github/workflows/**` and `**/*.md` changes

<br/>

## License

[Apache-2.0](LICENSE)
