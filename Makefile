VERSION = 0.0.1
TAG = $(VERSION)
PREFIX = nginx/nginx_exporter

.PHONY: build
build:
	CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X main.version=$(VERSION)" -o nginx_exporter

.PHONY: test
test:
	go test ./... -race -shuffle=on

.PHONY: container
container:
	docker build --build-arg VERSION=$(VERSION) --target container -f build/Dockerfile -t $(PREFIX):$(TAG) .

.PHONY: push
push: container
	docker push $(PREFIX):$(TAG)

.PHONY: deps
deps:
	@go mod tidy && go mod verify && go mod download

.PHONY: clean
clean:
	-rm -r dist
	-rm nginx_exporter
