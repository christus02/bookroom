IMG ?= raghulc/bookroom-frontend:latest

all: build

build:
	go build -o bin/bookroom main.go

docker-build:
	docker build -t ${IMG} .

docker-push:
	docker push ${IMG}

docker: build docker-build docker-push
