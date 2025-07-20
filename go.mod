module myproject

require (
	myproject/internal/handlers v1.2.3 // indirect
	myproject/internal/service v0.0.0-00010101000000-000000000000 // indirect
)

require myproject/internal/server v1.2.3

require myproject/pkg/morse v1.2.3 // indirect

go 1.24.5

replace myproject/internal/handlers => ./internal/handlers

replace myproject/internal/server => ./internal/server

replace myproject/internal/service => ./internal/service

replace myproject/pkg/morse v1.2.3 => ./pkg/morse
