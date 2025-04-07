FROM golang:1.24 AS crud
WORKDIR /app
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y nodejs \
    npm                       # note this one
RUN npm install -g @prisma/client
# Copy dependencies list
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod tidy
RUN go install github.com/steebchen/prisma-client-go

COPY ./crud ./crud

RUN (cd ./crud && go run github.com/steebchen/prisma-client-go generate)
FROM crud AS builder
WORKDIR /app

COPY ./firebase/ ./firebase/
COPY ./user-console/ ./user-console/
COPY ./middlewares/ ./middlewares/
COPY ./aws/ ./aws/
COPY ./openapi/ ./openapi/
COPY ./user-console/ ./user-console/