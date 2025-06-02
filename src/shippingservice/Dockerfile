FROM golang:1.24

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./src/shippingservice/go.mod  ./src/shippingservice/go.sum ./
RUN go mod download

COPY ./src/shippingservice ./

# it will build a binary for the application as shipping-app which can be run like ./shipping-app
RUN go build -v -o shipping-app

# this works without ./shipping-app because it is useing exec form which looks in cur dir while in shell form it looks at $PATH
CMD ["shipping-app"]  