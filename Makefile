all: build

build:
	@echo "App is creating. Please wait ..."
	@make -s clean
	@echo "Copy conf/app.ini To /usr/local/etc/puppeteer.hlzblog.top.ini ---DOING"
	@make -s ini
	@echo "Copy ---DONE"
	@echo "Serivce compiling ..."
	@go build -o puppeteer.hlzblog.top -v . 
	@echo "App is created"

ini:
	@cp conf/app.ini /usr/local/etc/puppeteer.hlzblog.top.ini

tool:
	@go vet ./...; true
	@gofmt -w .

clean:
	@rm -rf ./puppeteer.hlzblog.top
	@go clean -i .

test:
	@echo "Test --- START"
	@go test -v service/comic_service/*.go
	@echo "Test --- END"
