air:
	air

run:
	go run ./main.go

tidy:
	go mod tidy

up:
	docker-compose up -d --build

down:
	docker-compose down

test:
	go test ./... -v -cover -timeout 30m
	go test ./... -v -coverprofile=cov.out -timeout 30m
	go tool cover -html cov.out -o cover.html

test-no-i:
	go test .\... -v -short -cover

#go install github.com/vektra/mockery/v2@v2.43.2
mock:
	mockery --dir core/ --output=mocks/core --outpkg=mocks --all --with-expecter=true
	mockery --dir infra/ --output=mocks/infra --outpkg=mocks --all --with-expecter=true