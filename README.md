# Date Difference API

A simple HTTP API written in Go that calculates the difference between two dates in various units (years, months, weeks, days).

---

## Features

- Calculate difference between two dates in years, months, weeks, or days.
- Defaults to current date if the end date is not provided.
- Simple plain-text HTTP API.

---

## API Usage

### Endpoint

GET /

### Query Parameters

| Parameter | Description                          | Format         | Required | Default         |
| --------- | ---------------------------------- | -------------- | -------- | --------------- |
| `start`   | Start date                         | `YYYY-MM-DD`   | Yes      | N/A             |
| `end`     | End date                           | `YYYY-MM-DD`   | No       | Today’s date    |
| `unit`    | Unit for difference calculation    | `years`, `months`, `weeks`, `days` | No       | `years`         |

# Docker

## Build

```bash
docker build -t date-diff-api .
```

## Run

```bash
docker run -p 8080:8080 date-diff-api
```

# Go

## Run

```bash
go run main.go
```
