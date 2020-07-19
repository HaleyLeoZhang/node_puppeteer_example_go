all: debug

debug:
	@clear
	@echo "App debug is loading"
	@go run ./api/build/main.go -conf=./api/build/app.yaml

build:
	@echo "App is creating. Please wait ..."
	@make -s clean
	@echo "Copy ./api/build/app.example.yaml To ./api/build/app.yaml  ---DOING"
	@make -s ini
	@echo "Copy ---DONE"
	@echo "App compiling ..."
	@go build -o ./api/build/node_puppeteer_example_go -v ./api/build/main.go -conf=./api/build/app.yaml
	@echo "App is created"

run:
	./api/build/node_puppeteer_example_go  -conf=./api/build/app.example.yaml

ini:
	@cp ./api/build/app.example.yaml ./api/build/app.yaml

tool:
	@clear
	@go vet ./...; true
	@gofmt -w .

clean:
	@rm -rf ./node_puppeteer_example_go
	@go clean -i .

test:
	@clear
	@echo "Test --- START"
	@# 全量测试---暂时不考虑
	@# go test -v ./...
	@# 指定测试
	@go test -v ./api/service/ -conf=../../api/build/app.yaml
	@echo "Test --- END"