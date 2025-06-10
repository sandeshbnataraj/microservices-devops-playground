FROM --platform=$BUILDPLATFORM eclipse-temurin:21.0.5_11-jdk@sha256:a20cfa6afdbf57ff2c4de77ae2d0e3725a6349f1936b5ad7c3d1b06f6d1b840a AS build

ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM" 

WORKDIR /app

COPY . ./

RUN chmod +x ./gradlew

RUN ./gradlew installDist

EXPOSE 9555

CMD ["/app/build/install/hipstershop/bin/AdService"]