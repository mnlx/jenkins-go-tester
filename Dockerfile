FROM golang:1.12-alpine

COPY build/ /app


CMD /app/main