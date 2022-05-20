MAIN := github.com/farryl/project-mars/cmd/mars

run:
	@go.exe run ${MAIN}

build:
	@go.exe build -o build/mars ${MAIN}

fmt:
	@go.exe fmt ./...

db-up:
	@docker-compose up -d db

db-reset:
	mysql -h 127.0.0.1 -u root -proot -D mars < ./scripts/sql/a.sql

test-up:
	@docker-compose -f ./tests/docker-compose.yml up -d

test-clean:
	@docker-compose -f ./tests/docker-compose.yml down
