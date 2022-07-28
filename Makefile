build:
	docker build . -t go-ood
download:
	docker run -v $(shell pwd):/root go-ood go mod download
tidy: build
	docker run -v $(shell pwd):/root go-ood go mod tidy
lint:
	docker run -v $(shell pwd):/root go-ood golint ./...
gen:
	docker run -v $(shell pwd):/root go-ood go generate ./...
godoc:
	docker run -p 8080:8080 go-ood godoc -http=:8080
test: build
	docker run go-ood go test ./...
test-maze: build
	docker run go-ood go test github.com/ronnas/go-ood/cmd/maze
run-maze: build
	@docker run go-ood
test-pnp: build
	docker run go-ood go test github.com/ronnas/go-ood/cmd/pnp
run-pnp: build
	@docker run go-ood pnp
test-heap: build
	docker run go-ood go test github.com/ronnas/go-ood/cmd/heap
run-heap: build
	@docker run go-ood heap
