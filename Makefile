.PHONY: migrate run tests

migrate:
	@migrate -database 'mysql://root:mauFJcuf5dhRMQrjj@tcp(localhost:3306)/articles?query' -path ./migrations up

run:
	@ go run .

tests:
	@go test -v ./...