module myproject

go 1.24.5

replace myproject/internal/handlers => ./internal/handlers

replace myproject/internal/server => ./internal/server

replace myproject/internal/service => ./internal/service

replace myproject/pkg/morse => ./pkg/morse
