# Contributing to crema

Thanks for your interest! 🎉

## Getting Started

1. Fork this repository
2. Clone: `git clone https://github.com/YOUR_USERNAME/crema`
3. Run tests: `go test ./...`

## Development

### Requirements
- Go 1.22 or later
- golangci-lint

### Build & Test
```bash
go generate              # Generate code
go test ./...            # Run all tests
go test -cover ./...     # Check coverage
golangci-lint run        # Lint
go test -race ./...      # Race detector
```

### Adding Extensions
New providers/codecs go in `ext/`:
- Create `ext/yourprovider/` with separate `go.mod`
- Add `README.md` with usage example
- Write comprehensive tests

## Pull Request Process

1. Create a branch: `git checkout -b feature/your-feature`
2. Make changes with tests
3. Ensure `golangci-lint run` passes
4. Commit with clear message
5. Push and open PR

### PR Checklist
- [ ] Tests added/updated
- [ ] Test coverage remains high
- [ ] golangci-lint passes
- [ ] Documentation updated

## Reporting Bugs

Use GitHub Issues with:
- Go version (`go version`)
- OS/Architecture
- Minimal reproduction code
- Expected vs actual behavior

## Code of Conduct

Please follow our [Code of Conduct](CODE_OF_CONDUCT.md).
