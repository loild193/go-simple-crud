# --------------------------------
# Base stage
FROM golang:1.20-alpine as base

WORKDIR /app

RUN apk add --no-cache bash git openssh

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# --------------------------------
# Dev stage
FROM base as dev

WORKDIR /app

EXPOSE 8080

# Run the executable as dev
CMD ["tail", "-f", "/dev/null"]

# --------------------------------
# Builder stage
FROM base as builder

WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY --from=base /app/. ./

# Build the Go app
RUN go build -o main .

# --------------------------------
# Production stage
FROM builder as production

WORKDIR /app

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]


## Deploy
# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=dev /main /main
# COPY --from=dev /.env.* /

# EXPOSE 8080

# USER nonroot:nonroot

# ENTRYPOINT ["/main"]