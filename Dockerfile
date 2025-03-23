FROM golang:alpine AS build
WORKDIR /build
COPY ../../OneDrive/Desktop/test-task .
RUN CGO_ENABLED=0 go build -o main ./main.go

FROM alpine
WORKDIR /app
COPY --from=build /build/main .
COPY config /app/config/
COPY migrations /app/migrations/
CMD ./main