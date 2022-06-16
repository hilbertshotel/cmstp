FROM golang:1.18.3-alpine
RUN mkdir app
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build
EXPOSE 6789
CMD ["./get"]