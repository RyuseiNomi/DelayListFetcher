.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world

setup-dev:
	aws2 --endpoint-url http://127.0.0.1:9000 --profile minio_test s3 mb delay-list

exec-dev:
	sam build; sam local invoke DelayListFetcher \
		--region ap-northeast-1 \
		--docker-network delay_list_fetcher_network \
		--skip-pull-image
