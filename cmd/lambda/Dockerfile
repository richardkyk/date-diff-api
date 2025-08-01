FROM golang:1.24-alpine AS builder

RUN apk add --no-cache curl

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bootstrap ./cmd/lambda/main.go \
 && curl -Lo /usr/local/bin/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie \
 && chmod +x /usr/local/bin/aws-lambda-rie

FROM scratch

COPY --from=builder /app/bootstrap /var/task/bootstrap

COPY --from=builder /usr/local/bin/aws-lambda-rie /usr/local/bin/aws-lambda-rie

ENTRYPOINT ["/usr/local/bin/aws-lambda-rie"]

CMD ["/var/task/bootstrap"]
