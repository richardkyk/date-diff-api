FROM golang:1.24 AS build

WORKDIR /app

COPY . .

RUN go build -tags lambda.norpc -o main ./cmd/lambda/main.go

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /app/main ./main

ENTRYPOINT [ "./main" ]
