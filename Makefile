all: build

build:
	@echo "App is creating. Please wait ..."
	@go build -o puppeteer.hlzblog.top -v . 
	@echo "App is created"
	@echo "Copy conf/app.ini To /usr/local/etc/puppeteer.hlzblog.top.ini ---DOING"
	@cp conf/app.ini /usr/local/etc/puppeteer.hlzblog.top.ini
	@echo "Copy ---DONE"

ini:
	@cp conf/app.ini /usr/local/etc/puppeteer.hlzblog.top.ini

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf puppeteer.hlzblog.top
	go clean -i .

test:
	@echo "Test --- START"
	@go test -v service/comic_service/*.go
	@echo "Test --- END"
