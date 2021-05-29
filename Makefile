docker-compose-build-api: 
	@docker-compose build

docker-compose-up-api: 
	@docker-compose up

docker-compose-stop-api: 
	@docker-compose stop

run-as-docker: docker-compose-build-api docker-compose-up-api 

run:
	go run main.go

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o comparison main.go

setup:
	GO111MODULE=on go mod download

benchmark-test:
	go test -bench=.

compile-proto-files: ## Generate the Go files based on .proto files specification
	@for x in ./proto/*.proto; do \
		protoc --go_out=. $$x; \
	done

wrk-get:
	./wrk/get.sh ${endpoint}

wrk-post:
	./wrk/post.sh ${endpoint}