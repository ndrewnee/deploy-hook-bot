FROM golang:1.15 as modules

COPY go.mod go.sum /modules/
RUN cd /modules && go mod download

FROM golang:1.15 as builder

COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./deploy-hook-bot .

FROM alpine:3.13

RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/deploy-hook-bot .

EXPOSE 9998
ENTRYPOINT ["./deploy-hook-bot"]
