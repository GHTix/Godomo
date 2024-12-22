# syntax=docker/dockerfile:1

##
## Build the application from source
##

FROM docker.io/golang:1.23.3 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o release/ ./cmd/services/overkiz/...

##
## Run the tests in the container
##

# FROM build AS run-test
# RUN go test -v ./...

##
## Deploy the application binary into a lean image
##

# AS release
FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=build app/release/overkiz /overkiz

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/overkiz"]