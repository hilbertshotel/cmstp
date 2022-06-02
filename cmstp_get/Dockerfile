FROM golang:1.18-alpine
RUN mkdir cmstp_get
WORKDIR /cmstp_get
COPY . .
RUN go mod tidy
RUN go build
EXPOSE 6789
CMD ["./cmstp_get"]