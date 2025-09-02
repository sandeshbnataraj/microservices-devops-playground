# Microservices DevOps Playground — Monorepo CI/CD, Dockerfiles, and Swarm (fork)

### *Exploring how complex engineering challenges can be solved at scale.*

> **Fork notice:** This repository is a **fork** of a microservices sample.
> **My contribution:** the **root GitLab CI/CD pipeline** and the **Dockerfiles** under `src/*` (per service), plus **Swarm stack** definitions with **Traefik** reverse proxy. Application code is not mine; my work is the delivery system around it.

---

## What this shows (in one minute)

* **Monorepo CI/CD** that builds images for multiple services (Node, Python, Go, .NET) using **Docker Buildx**, with SHA tagging and registry push.&#x20;
* **Production-minded Dockerfiles** per service (cache-efficient layering, deterministic installs, minimal runtimes).
* **Docker Swarm stack** with **Traefik** as the edge router, health checks, VIP load balancing, and clean deploy/rollback settings. &#x20;

---

## CI/CD pipeline (root `.gitlab-ci.yml`)

* **Stages:** `build`, `test`, `deploy` (primary work is in `build` for image creation). **Default runner image:** `docker:28.2` with **DinD**.&#x20;
* **Per-service jobs:** Each service gets its own Buildx job; images are pushed to `$CI_REGISTRY_IMAGE/<service>[:tag]`. Most jobs tag with **`$CI_COMMIT_SHORT_SHA`** for traceability.&#x20;
* **Buildx setup:** Each job configures Docker context and **creates a named Buildx builder** with `--bootstrap --use`, then runs `docker buildx build --push`.&#x20;
* **Multi-arch example:** `adservice` builds **linux/amd64, linux/arm64** (QEMU prepped) and pushes a manifest list.&#x20;
* **Intermediate artifacts → runtime image:** For `adservice`, there’s a **two-phase flow**:

  1. `Dockerfile.build` produces a local artifact output.
  2. A **runtime** image is built using `--build-context gitlab-artifact=./src/adservice/build` + `Dockerfile.runtime`.
     This mirrors “build once, copy only what’s needed” for slim runtime images.&#x20;
* **Dev playground image:** A generic **dev** image can be built with a selectable `IMAGE_TAG` (gradle/node/python/golang) using `Dockerfile.dev`.&#x20;

**Why it matters:**

* **Reproducibility** (SHA tags, deterministic steps), **performance** (Buildx caching pattern), and **portability** (multi-arch where it counts).

---

## Dockerfiles per service (polyglot patterns)

Each service under `src/<service>/Dockerfile` follows its ecosystem’s production practices:

* **Node.js** → `npm ci`, lockfile-first layering, lean runtime (copy built artifacts only, clear CMD/ENTRYPOINT).
* **Python** → pinned `requirements.txt`, slim base (e.g. `python:*-slim`), Gunicorn/uvicorn when applicable, non-root runtime recommended.
* **Go** → multi-stage (builder → scratch/distroless), copy only the final static binary, healthchecks optional.
* **.NET** → multi-stage (`sdk` → `aspnet`), publish into runtime-only image, predictable ENTRYPOINT.

**Why it matters:** across languages, the **image principles stay consistent** (deterministic installs, cache-friendly layers, minimal runtime surface).

---

## Swarm stack (compose) + Traefik reverse proxy

You can run the whole system on **Docker Swarm** using the provided **stack file** and **Traefik** config.

### Traefik entry & providers

* **Entrypoints:** `web :80` and `secure_web :443`.
* **Providers:** Docker socket, **`exposedByDefault: false`** (opt-in exposure only).
* **Dashboard:** enabled via the **internal** API.&#x20;

### Stack services & routing

* **reverse-proxy (Traefik v3.4.1)** on the `front-tier` network, publishes :80, mounts `traefik.yml` (read-only) and the Docker socket. Uses **labels** to expose the dashboard under `/dashboard` and `/api`. Deploy config sets **replicas: 1**, **VIP mode**, restart/backoff policies, and **rollback/update** settings (parallelism, delay, start-first).&#x20;
* **frontend** and all backend services (e.g., `adservice`, `recommendationservice`, `checkoutservice`, `paymentservice`, `emailservice`, `currencyservice`, `cartservice`, `productcatalogservice`, `shippingservice`) run as **replicated** services on `front-tier`, many with **healthchecks** and **Traefik route labels**. The `frontend` exposes `/_healthz` and maps service ports via Traefik service labels.&#x20;
* **redis** uses a passworded startup command and is placed on the same network with **VIP** and deploy policies.&#x20;

**Why this design:**

* **Traefik labels** colocate routing config with service definitions (self-documenting).
* **Swarm deploy policies** (restart, healthcheck, update/rollback with **start-first**) reduce downtime and limit blast radius.
* **VIP mode** keeps service discovery simple for HTTP/gRPC services.

### Quick start (Swarm)

```bash
# 1) Init swarm (if not already)
docker swarm init

# 2) Deploy the stack (from repo root)
docker stack deploy -c compose.yml microdemo

# 3) Check services
docker stack services microdemo

# 4) Open the app (frontend)
# http://localhost:8080  (per compose ports)
# Traefik dashboard:
# http://localhost/dashboard
```

(Ports and routes per `compose.yml` + labels.)&#x20;

---

## Pipeline/Swarm map

```text
           ┌─────────────── GitLab CI ───────────────┐
           │                                          │
           │   buildx jobs per service (parallel)     │
           │   - docker context + buildx create       │
           │   - docker buildx build --push           │
           │   - tags: $CI_COMMIT_SHORT_SHA           │
           └───────────────┬──────────────────────────┘
                           │   pushed images
                           ▼
                 GitLab Container Registry
                           │
                           ▼
                 ┌──────────────────────┐
                 │ Docker Swarm (stack) │
                 │  - Traefik reverse   │
                 │    proxy (labels)    │
                 │  - VIP services      │
                 │  - healthchecks      │
                 │  - start-first roll  │
                 └──────────────────────┘
```

---

## Why this matters

* **Monorepo efficiency:** each service builds independently, in parallel, **only** what you need.&#x20;
* **Operational maturity:** Traefik + Swarm configs show health-aware rollouts, rollback safety, and clean ingress routing. &#x20;
* **Traceability:** every image is **immutably** linked to a commit SHA for audits and rollbacks.&#x20;
* **Polyglot readiness:** demonstrates solid Docker practices across **Node, Python, Go, .NET** in one repo.

---

## Attribution & contact

* **Fork of:** a microservices sample (see repo history for upstream).
* **My contribution:** root `.gitlab-ci.yml`, **Dockerfiles under `src/*`**, and **Swarm stack with Traefik** (`compose.yml`, `traefik.yml`).
* **Maintainer:** Sandesh Nataraj — [sandeshb.nataraj@gmail.com](mailto:sandeshb.nataraj@gmail.com)
