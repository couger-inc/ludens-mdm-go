FROM golang:1.24 AS crud
WORKDIR /app
ARG BUILD_DIR
ARG HANDLER
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y nodejs \
    npm                       # note this one
RUN npm install -g @prisma/client
# Copy dependencies list
COPY ./crud ./crud

RUN (cd ./crud && go run github.com/steebchen/prisma-client-go generate)
FROM crud AS builder
WORKDIR /app
ARG BUILD_DIR
ARG HANDLER
# Build with optional lambda.norpc tag
COPY ./${BUILD_DIR}/ ./${BUILD_DIR}/
COPY ./firebase/ ./firebase/
COPY ./user-console/ ./user-console/
COPY ./middleware/ ./middleware/
COPY ./aws/parameter-and-secrets/ ./aws/parameter-and-secrets/
COPY ./openapi/ ./openapi/
COPY ./user-console/ ./user-console/
RUN GOOS=linux GOARCH=amd64 go build -C ./${BUILD_DIR} -tags lambda.norpc -o /app/main ${HANDLER}
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=builder /app/main ./main
ENTRYPOINT [ "./main" ]