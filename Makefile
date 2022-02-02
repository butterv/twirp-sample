.PHONY: init
init:
	GO111MODULE=on go mod download

.PHONY: build
build:
	GO111MODULE=on go build

.PHONY: test
test:
	go test ./...

.PHONY: test-v
test-v:
	go test -v ./...

.PHONY: benchmark
benchmark:
	go test -bench . -benchmem

.PHONY: lint
lint:
	GO111MODULE=on golint ./...

generate-pb:
	$(eval OUTPUT_PATH := "./app/gen/go/v1")
	@for file in `\find proto/v1 -type f -name '*.proto'`; do \
		protoc \
			-I ./proto/v1/ \
			-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
			-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
			--go_out $(OUTPUT_PATH) \
			--go_opt paths=source_relative \
			--go-grpc_out $(OUTPUT_PATH) \
			--go-grpc_opt require_unimplemented_servers=false,paths=source_relative \
			--grpc-gateway_out $(OUTPUT_PATH) \
			--grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt paths=source_relative \
			--twirp_out $(OUTPUT_PATH) \
			--twirp_opt paths=source_relative \
			$$file; \
	done

app-build:
ifeq ($(tag),)
	@echo "Please execute this command with the docker image tag."
	@echo "Usage:"
	@echo "	$$ make app-build tag=<version>"
else
	docker build -f ./Dockerfile -t butterv/kubernetes-sample1:${tag} ./
endif

app-push:
ifeq ($(tag),)
	@echo "Please execute this command with the docker image tag."
	@echo "Usage:"
	@echo "	$$ make app-push tag=<version>"
else
	docker push butterv/kubernetes-sample1:${tag}
endif

skeema-init:
	@skeema init \
		-h 127.0.0.1 -p \
		--schema sample \
		--dir schemas \
		common # environment name

skeema-lint:
	@(cd schemas && skeema lint -p common)

skeema-diff:
	@(cd schemas && skeema diff -p common)

skeema-dry-run:
	@(cd schemas && skeema push --dry-run -p common)

skeema-push:
	@(cd schemas && skeema push -p common)
