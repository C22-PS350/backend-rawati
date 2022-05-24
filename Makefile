MAIN := github.com/C22-PS350/backend-rawati/cmd/rawati

run:
	@go.exe run ${MAIN}

build:
	@go.exe build -o build/rawati ${MAIN}

fmt:
	go fmt ./...

db-up:
	@docker-compose up -d db

db-stop:
	@docker-compose stop db

db-reset:
	mysql -h 127.0.0.1 -u root -proot -D rawati < ./scripts/sql/a.sql

docs:
	@swag init --dir cmd/rawati --output docs --ot "go,json" --parseDepth 10 --parseDependency

docs-fmt:
	swag fmt

docs-read:
	@docker compose up -d docs

test-prepare:
	@docker-compose -f ./tests/docker-compose.yml up -d

test-start:
	go test -v ./...

test-clean:
	@docker-compose -f ./tests/docker-compose.yml down
