NAME := zae

deps:
	go mod download

install:
	go build -o ${HOME}/.local/bin/${NAME} ./cmd/${NAME}/main.go

lint:
	golint ./...

vet:
	go vet ./...

grab:
	go run cmd/${NAME}/main.go --grab

imports:
	goimports -l -w .
