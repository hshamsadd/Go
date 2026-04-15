# Go
- go mod init Go/my-go-app
- go mod tidy
- goimports -w .
- go fmt ./...
- git rm -r --cached .
- git add .
- git commit -m "chore: apply production-level gitignore"
- docker build -t my-go-app:latest .
- docker run -p 8080:8080 \
  --env DATABASE_URL="postgres://hyfuser:hyfpassword@host.docker.internal:5432/postgres?sslmode=disable" \
  my-go-app:latest