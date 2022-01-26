FROM golang:1.17 as builder

WORKDIR /workspace

COPY main.go main.go
COPY webpage/ webpage/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bookroom main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/bookroom .
COPY webpage/ webpage/
USER 65532:65532

ENTRYPOINT ["/bookroom"]
