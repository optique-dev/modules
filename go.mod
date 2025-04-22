module github.com/Courtcircuits/optique-modules

go 1.24.2

require github.com/Courtcircuits/optique-modules/http v0.0.0-20230831135626

require (
	github.com/golang-migrate/migrate/v4 v4.18.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

require (
	github.com/Courtcircuits/optique-modules/sql v0.0.0-00010101000000-000000000000
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/gofiber/fiber/v2 v2.52.6 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
)

replace github.com/Courtcircuits/optique-modules/http => ./http

replace github.com/Courtcircuits/optique-modules/sql => ./sql
