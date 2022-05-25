FROM golang:1.17 AS docs

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/swaggo/swag/cmd/swag@latest && go mod download

COPY . .

RUN make docs



FROM swaggerapi/swagger-ui:v4.11.1

COPY --from=docs /app/docs/swagger.json /foo/

ENV SWAGGER_JSON=/foo/swagger.json

ENV BASE_URL=/swagger
