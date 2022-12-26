#build
FROM golang:1.19 AS Builder
WORKDIR /server

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -0 app .

# deploy
FROM alpine:latest AS Prod
WORKDIR /
COPY --from=Builder /server/app .
CMD ["./app"]