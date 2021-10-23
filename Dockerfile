ARG GO_VERSION=latest

FROM golang:$GO_VERSION as builder
WORKDIR /build/
COPY go.* .
COPY /cmd cmd
COPY /pkg /pkg
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o . ./...

FROM golang:$GO_VERSION as tester
WORKDIR /test/
COPY go.* .
COPY /cmd /cmd
COPY /pkg /pkg
# dependency
RUN go list -m -u all
#  static analysis
RUN go vet ./...
# test and benchmark
RUN go test -bench=. -benchmem ./...
# race condition
RUN CGO_ENABLED=1 GOOS=linux go build -v -a -race -installsuffix cgo -o . ./...

FROM scratch as runner
EXPOSE 80
ENTRYPOINT ["/microservice"]
ENV SERVICE "default"
ENV MESSAGE ""
COPY --from=builder /build/microservice /