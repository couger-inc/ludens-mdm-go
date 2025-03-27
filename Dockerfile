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
COPY ./${BUILD_DIR}/go.mod ./${BUILD_DIR}/
COPY ./user-console/go.mod ./user-console/
COPY ./middleware/ ./middleware/
COPY ./openapi/go.mod ./openapi/
COPY ./openapi/go.sum ./openapi/
COPY ./openapi/*.go ./openapi/
COPY ./user-console/*.go ./user-console/
COPY ./${BUILD_DIR}/go.sum ./${BUILD_DIR}/
COPY ./${BUILD_DIR}/${HANDLER} ./${BUILD_DIR}/${HANDLER}
RUN GOOS=linux GOARCH=amd64 go build -C ./${BUILD_DIR} -tags lambda.norpc -o /app/main ${HANDLER}
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=builder /app/main ./main
ENTRYPOINT [ "./main" ]