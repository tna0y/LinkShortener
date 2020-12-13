FROM golang:1.15

WORKDIR /build

COPY go.mod go.mod

RUN go mod download

COPY pkg ./pkg

COPY cmd ./cmd

RUN mkdir build

# CGO_ENABLED=0 cant work with go-sqlite3 but is required for alpine
RUN GOOS=linux go build -o app cmd/main.go cmd/config.go

FROM ubuntu:20.04

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates && \
    apt-get clean autoclean && \
    apt-get autoremove --yes && \
    rm -rf /var/lib/{apt,dpkg,cache,log}/

RUN groupadd -r user && useradd -r -s /bin/false -g user user

WORKDIR /app

COPY --from=0 /build/app .

RUN chown -R user:user /app

USER user

ENTRYPOINT ["/app/app"]
