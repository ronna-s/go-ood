build:
	@docker build . -t go-ood
download:
	docker run -v $(shell pwd):/root go-ood go mod download
tidy:
	docker run -v $(shell pwd):/root go-ood go mod tidy
lint:
	docker run -v $(shell pwd):/root go-ood golint ./...
gen:
	docker run -v $(shell pwd):/root go-ood go generate ./...
godoc:
	docker run -p 8080:8080 go-ood godoc -http=:8080
test:
	docker run go-ood go test ./...
test-maze:
	docker run go-ood go test github.com/ronnas/go-ood/cmd/maze
run-maze:
	@docker run go-ood
test-pnp:
	docker run go-ood go test github.com/ronnas/go-ood/pkg/pnp
run-pnp:
	@docker run -it go-ood pnp
test-heap:
	docker run go-ood go test github.com/ronnas/go-ood/pkg/heap
run-top:
	@docker run go-ood top
