FROM golang:latest

LABEL maintainer="dg.mondini@gmail.com"

WORKDIR /app

COPY main.go ./

RUN main.go

EXPOSE 2000

CMD ["./main"]