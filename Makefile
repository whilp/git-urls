default: test lint

test:
	go test -v ./...
	go test -covermode=count -coverprofile=profile.cov .

lint:
	golangci-lint run

install:
	go get -d -v ./... && go build -v ./...
	gometalinter --install --update

deps:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3
