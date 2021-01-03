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
	@#rm -rf ./app_job # 2020-11-29 计划中
	@#rm -rf ./app_admin # 2020-11-14 暂无计划
	@go clean -i .
	@echo "Remove Old Apps --- Done"

jsonclean:
	@rm -rf common/model/*/*easyjson*.go
	@rm -rf api/model/*easyjson*.go

json:
	@clear
	@echo "Creating all easyjson file"
	@rm -rf common/model/*/*easyjson*.go
	@easyjson -all common/model/*/*.go # 格式化 json ，需要 https://github.com/mailru/easyjson
	@echo "Created common/model/* easyjson file Success"
	@make jsonvo

jsonbo:
	@clear
	@echo "Creating easyjson file for common/model/bo"
	@rm -rf common/model/bo/*easyjson*.go
	@easyjson -all common/model/bo/*.go # 格式化 json ，需要 https://github.com/mailru/easyjson
	@echo "Created common/model/bo easyjson file Success"

jsonpo:
	@clear
	@echo "Creating easyjson file for common/model/po"
	@rm -rf common/model/po/*easyjson*.go
	@easyjson -all common/model/po/*.go # 格式化 json ，需要 https://github.com/mailru/easyjson
	@echo "Created common/model/po easyjson file Success"

jsonvo:
	@clear
	@echo "Creating easyjson file for api/model/"
	@rm -rf api/model/*easyjson*.go
	@easyjson -all api/model/*.go # 格式化 json ，需要 https://github.com/mailru/easyjson
	@echo "Created api/model/ easyjson file Success"

test:
	@clear
	@echo "Test --- START"
	@# 全量测试---暂时不考虑
	@# go test -v ./...
	@# 指定测试
	@go test -v ./api/service/ -conf=../../api/build/app.yaml
	@echo "Test --- END"