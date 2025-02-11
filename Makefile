lint:
	golangci-lint run --config .golangci.yml

test: test-code coverage

coverage:
	go tool cover -html=overalls.coverprofile -o coverage.html
	rm overalls.coverprofile

test-code:
	go test -covermode=count -coverprofile=overalls.coverprofile -p 2 github.com/kzozulya1/powapp...

gen:
	go generate ./...

build-server:
	go build -mod=readonly -o bin/pow-server ./server

build-client:
	go build -mod=readonly -o bin/pow-client ./client

build: build-server build-client

run-server: build-server
	bin/pow-server

run-client: build-client
	#bin/pow-client -addr 127.0.0.1:9876
	bin/pow-client

docker-build-push:
	docker build . -t kzozulya/pow-app:0.1.0
	docker push kzozulya/pow-app:0.1.0

docker-compose-up:
	sudo docker compose up
