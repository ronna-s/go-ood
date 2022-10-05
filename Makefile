build:
	@docker build . -t go-ood
download:
	docker run -v $(shell pwd):/root go-ood --rm go mod download
tidy:
	docker run -v $(shell pwd):/root --rm go-ood go mod tidy
lint:
	docker run -v $(shell pwd):/root --rm go-ood go vet ./...
gen:
	docker run -v $(shell pwd):/root --rm go-ood go generate ./...
godoc:
	docker run --rm -p 8080:8080 go-ood godoc -http=:8080
test:
	docker run go-ood go test ./...
test-maze:
	docker run -v $(shell pwd):/root --rm -it go-ood go test github.com/ronna-s/go-ood/cmd/maze
run-maze:
	@docker run -v $(shell pwd):/root --rm -it go-ood go run cmd/maze/maze.go
test-pnp:
	@docker run -v $(shell pwd):/root --rm go-ood go test github.com/ronna-s/go-ood/pkg/pnpdev
run-pnp:
	@docker run -v $(shell pwd):/root --rm -it go-ood go run cmd/pnp/pnp.go
test-heap:
	@docker run go-ood go test github.com/ronna-s/go-ood/pkg/heap
run-top:
	@docker run -v $(shell pwd):/root --rm -it go-ood go run cmd/top/top.go
