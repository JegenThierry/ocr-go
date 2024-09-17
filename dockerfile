FROM golang:1.23-alpine AS builder

RUN apk add --no-cache build-base gcc python3 py3-pip tesseract-ocr \
    tesseract-ocr-data-eng tesseract-ocr-data-deu poppler-utils ghostscript ocrmypdf

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ocr-api

FROM alpine:latest

RUN apk add --no-cache python3 py3-pip tesseract-ocr \
    tesseract-ocr-data-eng tesseract-ocr-data-deu poppler-utils ghostscript ocrmypdf

COPY --from=builder /app/ocr-api /usr/local/bin/ocr-api

ENV PATH="/venv/bin:${PATH}"

EXPOSE 8080
CMD ["ocr-api"]
