FROM golang:alpine

RUN mkdir /files
COPY errors.go /files
WORKDIR /files
RUN go build -o /files/errors errors.go
ENTRYPOINT [ "/files/errors" ]