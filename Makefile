build_app:
	go build ./cmd/app/main.go -c app


run_app:
	@go run ./cmd/app/main.go


test:
	@go build ./cmd/test/main.go
