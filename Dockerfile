FROM golang as builder
WORKDIR /builder
ADD . /builder
RUN GOPATH=.dependencies CGO_ENABLED=0 go build -o app

FROM scratch
COPY --from=builder /builder/app /app
ENTRYPOINT ["/app"]