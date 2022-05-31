FROM golang:1.18.2-buster AS builder

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-extldflags "-static"' -v -o /src/app cmd/goapiexample/main.go

FROM alpine:3.16
WORKDIR /root/
COPY --from=builder /src/app ./

CMD ["./app"]
