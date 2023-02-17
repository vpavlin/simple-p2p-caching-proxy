PORT=8888

generate:
	oapi-codegen --package api -o internal/api/api.gen.go openapi/openapi.yaml

server:
	go run cmd/server/main.go $(PORT)