FROM golang:1.22-alpine AS builder


WORKDIR /usr/local/src
RUN apk --no-cache add bash git make
COPY ["app/go.mod","app/go.sum","./"]
RUN go mod download

COPY app ./
RUN go build -o ./bin/app cmd/rest/main.go

FROM alpine AS runner
COPY --from=builder /usr/local/src/bin/app /

EXPOSE 8090

CMD ["./app"]