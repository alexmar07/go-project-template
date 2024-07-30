# Go Project Template

Based on [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)

## How to

### Init project
```bashgo install github.com/swaggo/swag/cmd/swag@latest
go mod init [repo]
```
Example:
```bash
go mod init github.com/go-project-template
```

### Update .gitignore

Remove go.sum and go.mod from `.gitignore`

## Packages

### Air

Live reload for Go apps. https://github.com/air-verse/air

### Swaggo

Swaggo is a tool that creates Swagger documentation for Go APIs. https://github.com/swaggo/swag

### Golangci-lint
Golangci-lint is a tool for checking Go code quality, finding issues, bugs, and style problems. https://github.com/golangci/golangci-lint