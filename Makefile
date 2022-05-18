PACKAGE := github.com/farryl/project-mars/cmd/mars

run:
	go.exe run ${PACKAGE}

build:
	go.exe build -o build/mars ${PACKAGE}

fmt:
	go.exe fmt ./...

test:
	go.exe test -v ./...
