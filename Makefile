.PHONY: build release

build:
	@go build -o ./build/late ./cmd/late

release:
	@./release.sh

