swagger-generate:
	mkdir swagger-generate
	swagger generate server --exclude-main -f ./swagger.yml -t swagger-generate

.PHONY: swagger-regenerate
swagger-regenerate:
	rm -rf swagger-generate
	make swagger-generate

.PHONY: clean
clean:
	rm -rf build

.PHONY: test
test:
	go mod tidy
	DYNAMO_TABLE_USERS=local_users AWS_REGION=ap-northeast-1  go test service/*.go -v

build:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o build/lambda cmd/lambda.go
	zip -j build/lambda.zip build/lambda

.PHONY: rebuild
rebuild: clean
	make build

deploy:
	aws lambda update-function-code --region ap-northeast-1 --function-name general-lambda-api --zip-file fileb://build/lambda.zip
