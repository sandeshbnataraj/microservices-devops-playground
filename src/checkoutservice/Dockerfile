FROM golang:1.24

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# it will build a binary for the application as shipping-app which can be run like ./shipping-app
# -v : verbose ; -o:ouptu dir name/path
# Use binary-name if it's placed inside a directory that’s in $PATH i.e /usr/local/bin is the path in $PATH so now i can run without ./app just app in CMD.
RUN go build -v -o /usr/local/bin/app

# this works without ./app because it is inside /usr/local/bin/ so no need for ./ since this path is included in $PATH. but if we run locllay or local dir in docker we need ./
CMD ["app"]