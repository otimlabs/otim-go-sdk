# Contributing to Otim Go SDK

Thank you for your interest in contributing to the Otim Go SDK! This document provides guidelines and information for contributors.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/otim-go-sdk.git`
3. Create a branch for your changes: `git checkout -b feature/your-feature-name`

## Development Setup

### Prerequisites

- Go 1.25 or later
- Make

### Building

```bash
go build ./...
```

### Running Tests

```bash
go test ./...
```

### Linting

```bash
go vet ./...
```

## Making Changes

1. Ensure your code follows Go conventions and passes `go vet`
2. Add tests for new functionality
3. Update documentation as needed
4. Keep commits focused and write clear commit messages

## Pull Request Process

1. Ensure all tests pass
2. Update the README.md if your changes affect usage
3. Fill out the pull request template
4. Request review from maintainers

## Code Style

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions focused and reasonably sized

## Reporting Issues

When reporting issues, please include:

- A clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Go version and OS
- Relevant code snippets or error messages

## License

By contributing, you agree that your contributions will be licensed under the same dual license as the project (Apache 2.0 and MIT).
