FROM eclipse-temurin:21.0.5_11-jdk@sha256:a20cfa6afdbf57ff2c4de77ae2d0e3725a6349f1936b5ad7c3d1b06f6d1b840a AS build

WORKDIR /app

COPY . ./

RUN chmod +x ./gradlew

RUN ./gradlew installDist

EXPOSE 9555

FROM --platform=$BUILDPLATFORM gcr.io/distroless/java21-debian12 AS prod

RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"

WORKDIR /app

COPY --from=build /app/build/install/hipstershop/ ./

CMD ["/app/bin/AdService"]