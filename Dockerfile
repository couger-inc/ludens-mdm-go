FROM couger-inc/mdm-go-lambda-base as builder
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