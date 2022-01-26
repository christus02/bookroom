FROM golang:1.17 as builder

WORKDIR /workspace

COPY main.go main.go
COPY images/ images/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o image-source main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/image-source .
COPY images/ images/
USER 65532:65532

ENTRYPOINT ["/image-source"]
