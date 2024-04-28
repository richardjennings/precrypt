.PHONY: build-dev

build-deps:
	docker run --rm -v $$PWD:/app --workdir /app --entrypoint="" node npm install
	docker run --rm -v $$PWD:/app --workdir /app -entrypoint="" node npm run build

build: build-deps
	go build -o precrypt .