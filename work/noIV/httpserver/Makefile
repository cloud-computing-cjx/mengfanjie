export tag=v1.0-metrics

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t jinyumantang/httpserver:${tag} .

push: release
	echo "pushing jinyumantang/httpserver"
	docker push jinyumantang/httpserver:${tag}