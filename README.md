# Go
- go mod init Go/my-go-app
- go mod tidy
- goimports -w .
- go fmt ./...
- git rm -r --cached .
- git add .
- git commit -m "chore: apply production-level gitignore"