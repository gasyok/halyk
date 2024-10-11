OUT_PATH:=$(CURDIR)/pkg
LOCAL_BIN:=$(CURDIR)/bin


build:
	go build -o $(LOCAL_BIN) ./cmd/app

.PHONY: build-image
build-image:
	$(info Building service image...)
	docker build . -f build/auction/Dockerfile -t assessments/auction

run:
	go run ./cmd/app/main.go

# ---------------------------
# Запуск базы данных в Docker
# ---------------------------

compose-up:
	docker-compose up -d postgres

compose-down:
	docker-compose down

compose-stop:
	docker-compose stop postgres

compose-start:
	docker-compose start postgres

compose-ps:
	docker-compose ps postgres

# ---------------------------
# Запуск миграций через Goose
# ---------------------------

goose-install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

goose-add:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" create rename_me sql

goose-up:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

goose-status:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" status

# ---------------------------------
# Запуск кодогенерации через protoc
# ---------------------------------

bin-deps: api/vendor-proto
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@latest

generate:
	mkdir -p ${OUT_PATH}
	protoc --proto_path api --proto_path api/vendor-proto \
		--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out=${OUT_PATH} --go_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out=${OUT_PATH} --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway --grpc-gateway_out ${OUT_PATH} --grpc-gateway_opt paths=source_relative \
		--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 --openapiv2_out=${OUT_PATH} \
		--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate --validate_out="lang=go,paths=source_relative:${OUT_PATH}" \
		./api/auction/v1/auction.proto

api/vendor-proto: api/vendor-proto/google/protobuf api/vendor-proto/google/api api/vendor-proto/protoc-gen-openapiv2/options api/vendor-proto/validate

api/vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway api/vendor-proto/grpc-ecosystem && \
 		cd api/vendor-proto/grpc-ecosystem && \
		git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
		git checkout
		mkdir -p api/vendor-proto/protoc-gen-openapiv2
		mv api/vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options api/vendor-proto/protoc-gen-openapiv2
		rm -rf api/vendor-proto/grpc-ecosystem

api/vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf api/vendor-proto/protobuf &&\
		cd api/vendor-proto/protobuf &&\
		git sparse-checkout set --no-cone src/google/protobuf &&\
		git checkout
		mkdir -p api/vendor-proto/google
		mv api/vendor-proto/protobuf/src/google/protobuf api/vendor-proto/google
		rm -rf api/vendor-proto/protobuf

api/vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis api/vendor-proto/googleapis && \
 		cd api/vendor-proto/googleapis && \
		git sparse-checkout set --no-cone google/api && \
		git checkout
		mkdir -p api/vendor-proto/google
		mv api/vendor-proto/googleapis/google/api api/vendor-proto/google
		rm -rf api/vendor-proto/googleapis

api/vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate api/vendor-proto/tmp && \
		cd api/vendor-proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p api/vendor-proto/validate
		mv api/vendor-proto/tmp/validate api/vendor-proto/
		rm -rf api/vendor-proto/tmp

