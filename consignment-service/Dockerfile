# Step 1 - we create a docker image to create executable for the current application
# Step 2 - Copy this executable (only) to a new clean image

## STEP 1

# Creating builder image on top of official golang alpine image. This will have
# all the right tools required for building the executable
FROM golang:alpine as builder

# Assuming that git is required for pulling the golang modules. Else, can also
# be used to pull the latest code before building
RUN apk --no-cache add git

WORKDIR /app/shippy-service-consignment

# Copying current directory into the working directory on the image
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o shippy-service-consignment -a -installsuffix cgo main.go repository.go handler.go datastore.go

# STEP 2

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

# Copying executable(binary) from the builder image (instead of the host)
COPY --from=builder /app/shippy-service-consignment/shippy-service-consignment .

CMD ["./shippy-service-consignment"]
