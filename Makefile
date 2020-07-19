all: debug

debug:
	@clear
	@echo "APP debug is loading"
	@go run ./api/build/main.go -conf=./api/build/app.yaml

build:
	@clear
	@echo "New APP is creating. Please wait ..."
	@make -s clean
	@echo "App compiling ..."
	@go build -o ./api/build/node_puppeteer_example_go -v ./api/build/main.go
	@echo "App is created"

run:
	@clear
	@echo "APP is runing. Please wait ..."
	@./api/build/node_puppeteer_example_go  -conf=./api/build/app.yaml

ini:
	@clear
	@cp ./api/build/app.example.yaml ./api/build/app.yaml
	@echo "Copy yaml success"

tool:
	@clear
	@go vet ./...; true
	@gofmt -w .

clean:
	@echo "Remove Old APP ... "
	@rm -rf ./api/build/node_puppeteer_example_go
	@echo "Remove Old App --- Done"

test:
	@clear
	@echo "Test --- START"
	@# 全量测试---暂时不考虑
	@# go test -v ./...
	@# 指定测试
	@go test -v ./api/service/ -conf=../../api/build/app.yaml
	@echo "Test --- END"