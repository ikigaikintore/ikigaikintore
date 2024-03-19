# Architecture

Table of Contents

1. [Overview](#overview)
2. [Release flow](#release-flow)
3. [Tech stack](#tech-stack)
4. [Architect Layout](#architect-layout)

---

## Overview

The architecture is based in a monolith. The main reason is to keep it simple and easy to maintain. The idea is to have low costs and low complexity and low latency.

## Release flow

The release flow is simple, only master branch is used. The idea is to have a simple flow, using gitflow for feature branches and tags for deploying new versions.

### Backend release

When deploying a new version, a new tag is created and the image is stored in google artifact. Then, the image is deployed in google cloud run.

This involves the proxy and the backend.

### Frontend release

The frontend uses firebase. It also uses tags for deploying new versions.

### Libs release

Each library can be released independently. The idea is to have a simple release flow, using tags and github actions.
Using the libs_tagger workflow action, a new tag is created and the library is released in github packages.

## Tech stack

Backend: Go
Frontend: NextJS
Use IaC for managing the infrastructure

## Architect Layout

TBD with some ideas:

- DDD with hexagonal layered
- Cache, Database: TBD
- Communicate with telegram!

![poc](./assets/poc_view.png)

### Frontend

Using nextjs, building a simple dashboard using web components which communicates with the backend via the proxy.

### Backend

The backend is a simple go application, using the standard library and some external libraries and call other third party services using their APIs.

The endpoints are built using a proto file.

### Proxy

The proxy is the connector between the backend and the frontend. It controls the access and expose the service publicly.



## Local development

Using docker-compose and [tilt](https://tilt.dev/) for local development. The idea is to have a local environment as close as possible to the production environment.

Ports:

- proxy:
  - 8080
  - 18080
- proxy_bot:
  - 7080
  - 17080
- backend:
  - 8999
  - 18999
- frontend:
  - 9099:9099
  - 5001:5001
  - 8080:8080
  - 9000:9000
  - 5000:5000
  - 8085:8085
  - 9199:9199
  - 4000:4000
