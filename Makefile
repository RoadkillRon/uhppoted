DEBUG ?= --debug
VERSION = v0.6.3
DIST   ?= development
LDFLAGS = -ldflags "-X uhppote.VERSION=$(VERSION)" 

.PHONY: docker
.PHONY: simulator
.PHONY: uhppoted-rest
.PHONY: uhppoted-mqtt
.PHONY: uhppoted-app-s3
.PHONY: uhppoted-app-sheets

all: test      \
	 benchmark \
     coverage

clean:
	go clean
	rm -rf bin

format: 
	cd uhppote-core;        go fmt ./...
	cd uhppoted-api;        go fmt ./...
	cd uhppote-simulator;   go fmt ./...
	cd uhppote-cli;         go fmt ./...
	cd uhppoted-rest;       go fmt ./...
	cd uhppoted-mqtt;       go fmt ./...
	cd uhppoted-app-s3;     go fmt ./...
	cd uhppoted-app-sheets; go fmt ./...
	cd integration-tests;   go fmt ./...
	go fmt ./...

build: format
	mkdir -p bin
	cd uhppote-core;        go build            ./...
	cd uhppoted-api;        go build            ./...
	cd uhppote-simulator;   go build -o ../bin/ ./...
	cd uhppote-cli;         go build -o ../bin/ ./...
	cd uhppoted-rest;       go build -o ../bin/ ./...
	cd uhppoted-mqtt;       go build -o ../bin/ ./...
	cd uhppoted-app-s3;     go build -o ../bin/ ./...
	cd uhppoted-app-sheets; go build -o ../bin/ ./...

test: build
	cd uhppote-core;        go test ./...
	cd uhppoted-api;        go test ./...
	cd uhppote-simulator;   go test ./...
	cd uhppote-cli;         go test ./...
	cd uhppoted-rest;       go test ./...
	cd uhppoted-mqtt;       go test ./...
	cd uhppoted-app-s3;     go test ./...
	cd uhppoted-app-sheets; go test ./...
#	go test ./...

vet: build
	go vet ./...

lint: build
	golint ./...

benchmark: build
	go test -bench ./...

coverage: build
	go test -cover ./...

integration-tests: build
#	go test integration-tests/cli/*.go
#	go clean -testcache && go test -count=1 integration-tests/mqttd/*.go
	go clean -testcache && go test -count=1 integration-tests/simulator/*.go

build-all: test vet
	mkdir -p dist/linux/$(DIST)
	mkdir -p dist/arm7/$(DIST)
	mkdir -p dist/darwin/$(DIST)
	mkdir -p dist/windows/$(DIST)
	mkdir -p dist/openapi/$(DIST)

	cd uhppote-cli; env GOOS=linux   GOARCH=amd64       go build -o ../dist/linux/$(DIST)    ./...
	cd uhppote-cli; env GOOS=linux   GOARCH=arm GOARM=7 go build -o ../dist/arm7/$(DIST)     ./...
	cd uhppote-cli; env GOOS=darwin  GOARCH=amd64       go build -o ../dist/darwin/$(DIST)   ./...
	cd uhppote-cli; env GOOS=windows GOARCH=amd64       go build -o ../dist/windows/$(DIST)  ./...

	cd uhppoted-rest; env GOOS=linux   GOARCH=amd64       go build -o ../dist/linux/$(DIST)   ./...
	cd uhppoted-rest; env GOOS=linux   GOARCH=arm GOARM=7 go build -o ../dist/arm7/$(DIST)    ./...
	cd uhppoted-rest; env GOOS=darwin  GOARCH=amd64       go build -o ../dist/darwin/$(DIST)  ./...
	cd uhppoted-rest; env GOOS=windows GOARCH=amd64       go build -o ../dist/windows/$(DIST) ./...

	cd uhppoted-mqtt; env GOOS=linux   GOARCH=amd64       go build -o ../dist/linux/$(DIST)   ./...
	cd uhppoted-mqtt; env GOOS=linux   GOARCH=arm GOARM=7 go build -o ../dist/arm7/$(DIST)    ./...
	cd uhppoted-mqtt; env GOOS=darwin  GOARCH=amd64       go build -o ../dist/darwin/$(DIST)  ./...
	cd uhppoted-mqtt; env GOOS=windows GOARCH=amd64       go build -o ../dist/windows/$(DIST) ./...	

	cd uhppoted-app-s3; env GOOS=linux   GOARCH=amd64       go build -o ../dist/linux/$(DIST)   ./...
	cd uhppoted-app-s3; env GOOS=linux   GOARCH=arm GOARM=7 go build -o ../dist/arm7/$(DIST)    ./...
	cd uhppoted-app-s3; env GOOS=darwin  GOARCH=amd64       go build -o ../dist/darwin/$(DIST)  ./...
	cd uhppoted-app-s3; env GOOS=windows GOARCH=amd64       go build -o ../dist/windows/$(DIST) ./...	

	cd uhppoted-app-sheets; env GOOS=linux   GOARCH=amd64       go build -o ../dist/linux/$(DIST)   ./...
	cd uhppoted-app-sheets; env GOOS=linux   GOARCH=arm GOARM=7 go build -o ../dist/arm7/$(DIST)    ./...
	cd uhppoted-app-sheets; env GOOS=darwin  GOARCH=amd64       go build -o ../dist/darwin/$(DIST)  ./...
	cd uhppoted-app-sheets; env GOOS=windows GOARCH=amd64       go build -o ../dist/windows/$(DIST) ./...	

	cd uhppote-simulator; env GOOS=linux   GOARCH=amd64       go build -o ../dist/linux/$(DIST)   ./...
	cd uhppote-simulator; env GOOS=linux   GOARCH=arm GOARM=7 go build -o ../dist/arm7/$(DIST)    ./...
	cd uhppote-simulator; env GOOS=darwin  GOARCH=amd64       go build -o ../dist/darwin/$(DIST)  ./...
	cd uhppote-simulator; env GOOS=windows GOARCH=amd64       go build -o ../dist/windows/$(DIST) ./...

	cp uhppoted-rest/documentation/uhppoted-api.yaml documentation/openapi
	cp uhppoted-rest/documentation/uhppoted-api.yaml install/openapi
	cp -r install/openapi/* dist/openapi/$(DIST)/

release: build-all docker integration-tests
	find . -name ".DS_Store" -delete
	tar --directory=dist/linux  --exclude=".DS_Store" -cvzf dist/$(DIST)-linux.tar.gz $(DIST)
	tar --directory=dist/arm7   --exclude=".DS_Store" -cvzf dist/$(DIST)-arm7.tar.gz $(DIST)
	tar --directory=dist/darwin --exclude=".DS_Store" -cvzf dist/$(DIST)-darwin.tar.gz $(DIST)
	cd dist/windows; zip --recurse-paths ../$(DIST)-windows.zip $(DIST)

build-github: 
	cd uhppote-core; go build ./...

debug: build
	./bin/uhppote-cli --debug --broadcast 192.168.1.100:54321 get-events 201020304

simulator: 
	./bin/uhppote-simulator --debug --bind 0.0.0.0:60000 --rest 0.0.0.0:8000 --devices "./runtime/simulation/devices"

uhppoted-rest:
	./bin/uhppoted-rest --console

uhppoted-mqtt: 
	./bin/uhppoted-mqtt --console

swagger: 
	docker run --detach --publish 80:8080 --name swagger --rm swaggerapi/swagger-editor 
	open http://127.0.0.1:80

docker:
	cd uhppote-simulator; env GOOS=linux GOARCH=amd64 go build -o ../docker/simulator     ./...
	cd uhppote-simulator; env GOOS=linux GOARCH=amd64 go build -o ../docker/uhppoted-rest ./...
	cd uhppote-simulator; env GOOS=linux GOARCH=amd64 go build -o ../docker/integration-tests/simulator ./...
	cd uhppoted-rest;     env GOOS=linux GOARCH=amd64 go build -o ../docker/uhppoted-rest ./...
	
	docker image     prune -f
	docker container prune -f
	docker build -f ./docker/simulator/Dockerfile     -t simulator       . 
	docker build -f ./docker/uhppoted-rest/Dockerfile -t uhppoted        . 
	docker build -f ./docker/hivemq/Dockerfile        -t hivemq/uhppoted . 
	docker build -f ./docker/integration-tests/simulator/Dockerfile -t integration-tests/simulator . 

docker-simulator:
	docker run --detach --publish 8000:8000 --publish 60000:60000/udp --name simulator --rm simulator
	sleep 1
	./bin/uhppote-cli --debug set-listener 405419896 192.168.1.100:60001
	./bin/uhppote-cli --debug set-listener 303986753 192.168.1.100:60001

docker-hivemq:
	docker run --detach --publish 8081:8080 --publish 1883:1883 --publish 8883:8883 --name hivemq --rm hivemq/uhppoted

docker-rest:
	docker run --detach --publish 8080:8080 --rm uhppoted

docker-stop:
	docker stop simulator
	docker stop hivemq

docker-integration-tests:
	docker run --detach --publish 8000:8000 --publish 60000:60000/udp --name qwerty --rm integration-tests/simulator

hivemq-listen:
#	mqtt subscribe --topic 'twystd/uhppoted/#'
	open runtime/mqtt-spy-0.5.4-jar-with-dependencies.jar



