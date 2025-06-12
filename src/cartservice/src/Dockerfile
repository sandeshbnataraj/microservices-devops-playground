# Use the .NET SDK base image for building the application
# --platform=$BUILDPLATFORM ensures cross-platform builds using BuildKit/Buildx
FROM --platform=$BUILDPLATFORM mcr.microsoft.com/dotnet/sdk:9.0.101-noble@sha256:1f13e67d295e02abdfd187c341f887442bad611eda536766172ced401fc8b9fa AS builder

# Build-time variables used to manage architecture/platform-specific logic
ARG BUILDPLATFORM        # E.g., linux/amd64 (your local machine)
ARG TARGETPLATFORM       # E.g., linux/arm64 (target deployment platform)
ARG TARGETARCH           # E.g., amd64, arm64 (used by dotnet)

# Set working directory inside the container
WORKDIR /app

# Copy only the project file first to take advantage of Docker layer caching
COPY cartservice.csproj ./

# Restore .NET dependencies for the target architecture
# This avoids re-downloading packages if only code changes
RUN dotnet restore cartservice.csproj -a ${TARGETARCH}

# Copy the rest of the source code into the container
COPY . ./

# Publish the app:
# -p:PublishSingleFile=true          → Output one standalone binary
# -a $TARGETARCH                     → Compile for target architecture (e.g., arm64)
# --self-contained true              → Package the .NET runtime (no install needed on target)
# -p:PublishTrimmed=true             → Remove unused assemblies (makes binary smaller)
# -p:TrimMode=full                   → Aggressive trimming (may break dynamic code)
# -c release                         → Use release build configuration
# -o /cartservice                    → Output directory for published binary
RUN dotnet publish \
    -p:PublishSingleFile=true \
    -a $TARGETARCH \
    --self-contained true \
    -p:PublishTrimmed=true \
    -p:TrimMode=full \
    -c release \
    -o /cartservice

# ─────────────────────────────
# Runtime Image (minimal and secure)
# ─────────────────────────────

# Use runtime-deps base image with no .NET SDK (chiseled = extra minimal & secure)
FROM mcr.microsoft.com/dotnet/runtime-deps:9.0.1-noble-chiseled@sha256:6f7466eda39e24efaf7eab2325e15d776a685d13cc93b4ea0cde9ee4f7982210

# Set working directory for runtime
WORKDIR /app

# Copy the published single-file app from the builder stage
COPY --from=builder /cartservice .

# Expose the port the app listens on
EXPOSE 7070

# Disable diagnostics tools (makes runtime lighter and more secure)
# ASPNETCORE_HTTP_PORTS = which port the app listens on internally
ENV DOTNET_EnableDiagnostics=0 \
    ASPNETCORE_HTTP_PORTS=7070

# Drop root privileges → follow security best practices
USER 1000

# Define the app entrypoint (run the compiled binary)
ENTRYPOINT ["/app/cartservice"]

#RUN dotnet publish -c Release -o app
#CMD ["dotnet", "app/cartservice.dll"]