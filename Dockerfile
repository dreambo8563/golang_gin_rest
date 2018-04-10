FROM golang:1.10
LABEL maintainer="vincent"
RUN mkdir -p /go/src/vincent.com/golangginrest
WORKDIR /go/src/vincent.com/golangginrest
COPY ./docker-bash.sh /
RUN chmod +x /docker-bash.sh
EXPOSE 8080
CMD ["go run main.go"] 
