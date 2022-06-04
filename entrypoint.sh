wait-for "postgres:5432" -- "$@"
wait-for "crawler:7317" -- "$@"

# builds application binary for better perfomance
go build -o ./bin/main main.go

# runs binary
./bin/main
