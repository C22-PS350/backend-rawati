MAIN := github.com/farryl/project-mars/cmd/mars

run:
	@go.exe run ${MAIN}

build:
	@go.exe build -o build/mars ${MAIN}

fmt:
	@go.exe fmt ./...

test:
	@go.exe test -v ./...

db-up:
	@docker-compose up -d db

clean:
	@docker-compose down
