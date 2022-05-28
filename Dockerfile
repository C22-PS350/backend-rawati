FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV LISTEN_IP=0.0.0.0
ENV LISTEN_PORT=8080

RUN go build -o build/rawati ./cmd/rawati/

CMD [ "./build/rawati" ]

