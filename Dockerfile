FROM golang:1.26.3-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN go build -o sprint12 ./main.go

FROM alpine:latest AS app
WORKDIR /app
COPY --from=build /app/sprint12 /app/sprint12
ENTRYPOINT ["./sprint12"]
