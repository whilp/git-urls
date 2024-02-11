default: test lint
LINT=golangci-lint run

test:
	go test -v ./...
	go test -covermode=count -coverprofile=profile.cov .

lint:
	$(LINT) ./...

install:
	go get -d -v ./... && go build -v ./...
	$(LINT) ./...

deps:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	go mod download -x
	go mod verify
	go mod tidy -v
