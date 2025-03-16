build:
	go fmt && go vet && go build

install:
	go fmt && go vet && go build && go install

new:
	go run main.go new github.com/mahbubzulkarnain/example

init:
	go run main.go init github.com/mahbubzulkarnain/example

entity:
	go run main.go entity --name User --table_name Users

dto:
	go run main.go dto --name User --table_name Users

repository:
	go run main.go repository --name User --table_name Users

service:
	go run main.go service --name User
