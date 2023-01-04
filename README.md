## GoToko POS API

## Development Normal

- Copy dulu `.env.example` ke `.env`. Ubah dan sesuaikan value yang kurang sesuai.
- `go mod download` kalau belum.
- `go install github.com/swaggo/swag/cmd/swag@latest` untuk install swagger generator
- Jalankan `make swag-init`
- Jalankan `make run`

## Development Docker

- Copy dulu `.env.example` ke `.env`. Ubah dan sesuaikan value yang kurang sesuai.
- `docker build -t gotoko-pos-api .`
- Tunggu sampai selesai, kalau belum selesai ya gak bisa dijalankan.
- `docker run -p 8080:8080 -t gotoko-pos-api`

## Test

- `go test -v ./...`

## Specs

- [https://jsonapi.org/](https://jsonapi.org/)
- Auth Bearer token

## Tech Stack
- Golang
- Gin
- MySQL
- JWT
