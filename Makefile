all: debug

debug:
	@clear
	@echo "App API debug is loading"
	@go run ./api/build/main.go -conf=./api/build/app.yaml

build:
	@clear
	@echo "App API is creating. Please wait ..."
	@rm -rf ./app_api
	@echo "App API compiling ..."
	@go build -o ./app_api -v ./api/build/main.go
	@echo "App API is created"

run:
	@clear
	@echo "APP is loading. Please wait ..."
	@./app_api  -conf=./api/build/app.yaml

ini:
	@clear
	@cp ./api/build/app.example.yaml ./api/build/app.yaml
	@echo "API config.yaml initial success"

tool:
	@clear
	@go vet ./...; true
	@gofmt -w .

clean:
	@clear
	@echo "Remove Old Apps ... "
	@rm -rf ./app_api
	@rm -rf ./app_job
	@rm -rf ./app_admin
	@go clean -i .
	@echo "Remove Old Apps --- Done"

easyjson:
	@clear
    # 以下会为 api/model 目录下的所有结构体生成 easyjson 文件
	@echo "Creating easyjson file for api/model"
	@rm -rf api/model/*_easyjson.go
	@easyjson -all api/model/*.go # 格式化 json ，需要 https://github.com/mailru/easyjson
	@echo "Created API easyjson file Success"

test:
	@clear
	@echo "Test --- START"
	@# 全量测试---暂时不考虑
	@# go test -v ./...
	@# 指定测试
	@go test -v ./api/service/ -conf=../../api/build/app.yaml
	@echo "Test --- END"