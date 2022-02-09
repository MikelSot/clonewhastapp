wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

CompileDaemon --build="go build cmd/main.go" --command=./main
