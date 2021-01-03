# First stage: build the executable.
FROM golang:buster AS builder

# Create the user and group files that will be used in the running 
# container to run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'random:x:65534:65534:random:/:' > /user/passwd && \
    echo 'random:x:65534:' > /user/group

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/shitpostingio/randomapi

# Import the code from the context.
COPY . .

# Build the executable
RUN go install

# Final stage: the running container.
FROM debian:buster

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Copy the built executable
COPY --from=builder /go/bin/randomapi /home/random/randomapi

# Install dependencies and create home directory
RUN apt update && apt install -y ca-certificates; \ 
    chown -R random /home/random

# Set the workdir
WORKDIR /home/random

# Perform any further action as an unprivileged user.
USER random:random

# Run the compiled binary.
CMD ["./randomapi" "-config", "config/conf.toml"]