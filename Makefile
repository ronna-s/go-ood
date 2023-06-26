CONTAINER_ENGINE ?= podman

build:
	@$(CONTAINER_ENGINE) build . -t go-ood
download:
	$(CONTAINER_ENGINE) run -v $(shell pwd):/root go-ood --rm go mod download
tidy:
	$(CONTAINER_ENGINE) run -v $(shell pwd):/root --rm go-ood go mod tidy
lint:
	$(CONTAINER_ENGINE) run -v $(shell pwd):/root --rm go-ood go vet ./...
gen:
	$(CONTAINER_ENGINE) run -v $(shell pwd):/root --rm go-ood go generate ./...
godoc:
	$(CONTAINER_ENGINE) run --rm -p 8080:8080 go-ood godoc -http=:8080
test:
	$(CONTAINER_ENGINE) run go-ood go test ./...
test-maze:
	$(CONTAINER_ENGINE) run -v $(shell pwd):/root --rm -it go-ood go test github.com/ronna-s/go-ood/cmd/maze
run-maze:
	@$(CONTAINER_ENGINE) run -v $(shell pwd):/root --rm -it go-ood go run cmd/maze/maze.go
test-pnp:
	@$(CONTAINER_ENGINE) run -v $(shell pwd):/root --rm go-ood go test github.com/ronna-s/go-ood/pkg/pnpdev
run-pnp:
	@$(CONTAINER_ENGINE) run -v $(shell pwd):/root --rm -it go-ood go run cmd/pnp/pnp.go
test-heap:
	@$(CONTAINER_ENGINE) run go-ood go test github.com/ronna-s/go-ood/pkg/heap
run-top:
	@$(CONTAINER_ENGINE) run -v $(shell pwd):/root --rm -it go-ood go run cmd/top/top.go
