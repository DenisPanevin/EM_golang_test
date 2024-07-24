FROM golang:1.22-alpine AS builder


WORKDIR /usr/local/src
RUN apk --no-cache add bash git make
COPY ["go.mod","go.sum","./"]
RUN go mod download

COPY . ./
RUN go build -o ./bin/app ./main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /


EXPOSE 9090:9090

CMD ["./app"]