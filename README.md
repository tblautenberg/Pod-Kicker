# Pod Kicker

Pod Kicker is a lightweight Go-based web server designed to automatically restart Kubernetes deployments when a new container image is pushed to Docker Hub. Working and tested in Kubernetes 1.32.3, and with Ubuntu server 24.04.2 LTS

Still a work in progress, but feel free to add PR if you help out!

# Overview

When integrated into your CI/CD pipeline, Pod Kicker listens for webhook events from Docker Hub and triggers a rollout restart of the specified deployment in your Kubernetes cluster.

# Requirements

To use Pod Kicker, ensure the following:

A GitHub Actions workflow configured to build and push your Docker image to Docker Hub.

A Docker Hub webhook configured to send a POST request to Pod Kicker when a new image is pushed.

# Usage

Pod Kicker can be run inside Kubernetes either by building your own Docker image or by using the latest pre-built image available at:
XXXXXXXXX

By default, the server listens on port 8450 and currently supports a single POST endpoint:

/KickThatPod
Additional endpoints and features are planned for future releases.

# Deployment Permissions 

Ensure the Pod Kicker deployment has the necessary RBAC permissions to restart deployments. At a minimum, it requires get, list, and patch permissions on deployments in the relevant namespaces.

# License

MIT — feel free to use, modify, and contribute! 🤘

