# Our base image that we use to build
FROM golang:1.16-alpine as builder
# Where we are going to work from
WORKDIR /go/src/app
# Copy everything in our repo into the workdir
COPY . .
# Get the dependencies and then do the build
# This is in one line to reduce layers. 
# While that doesn't really matter for this image as we are doing a multi-stage build, it makes a good example of the technique
RUN go get -d -v ./... && go build -v
# The base of our final image that we will deploy
FROM alpine
# Copying the binary we built into a smaller base image without all of the build dependencies
COPY --from=builder /go/src/app/go-mysql-crud /bin/
# Declaring the port that the app will run on and exposing it.
EXPOSE 8005
# The entrypoint for the container.
ENTRYPOINT [ "/bin/go-mysql-crud" ]
