FROM --platform=$BUILDPLATFORM eclipse-temurin:21.0.5_11-jdk@sha256:a20cfa6afdbf57ff2c4de77ae2d0e3725a6349f1936b5ad7c3d1b06f6d1b840a AS builder

WORKDIR /app

#Only copy what's needed to download dependencies
COPY build.gradle gradlew ./
COPY gradle gradle
RUN chmod +x ./gradlew
RUN ./gradlew downloadRepos

COPY . .
RUN chmod +x ./gradlew
RUN ./gradlew installDist

#FROM gcr.io/distroless/java21-debian12 AS prod
#wont work here because the distroless has no other tools evcep app runtime!
#ARG BUILDPLATFORM
##ARG TARGETPLATFORM
#RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"
FROM eclipse-temurin:21.0.5_11-jre-alpine@sha256:4300bfe1e11f3dfc3e3512f39939f9093cf18d0e581d1ab1ccd0512f32fe33f0

WORKDIR /app

COPY --from=builder /app ./

EXPOSE 9555
ENTRYPOINT ["/app/build/install/hipstershop/bin/AdService"]