FROM golang:1.10
ENV GO_ENV=production
ENV GIN_MODE=release
EXPOSE 3000
EXPOSE 3010