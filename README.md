# Pod Kicker

Pod Kicker is a lightweight Go-based web server designed to automatically restart Kubernetes deployments when a JSON payload is recived from hub.docker.com webhook. Working and tested in Kubernetes 1.32.3, and with Ubuntu server 24.04.2 LTS

Still a work in progress, but feel free to add PR if you want to help out!

# Overview

When integrated into your CI/CD pipeline, Pod Kicker listens for webhook events from Docker Hub and triggers a rollout restart of the specified deployment in your Kubernetes cluster.

# Requirements

To use Pod Kicker, ensure the following:

A GitHub Actions workflow configured to build and push your Docker image to Docker Hub.

A Docker Hub webhook configured to send a POST request to Pod Kicker when a new image is pushed.

# Usage

By default, the server listens on port 8450 and currently supports a single POST endpoint:

www.yoururl.com/updateDeployment

Additional endpoints and features are planned for future releases.

# Deployment Permissions 

Ensure the Pod Kicker deployment has the necessary RBAC permissions to restart deployments. At a minimum, it requires get, list, and patch permissions on deployments in the relevant namespaces. There is a default yaml file, where it i's already configured.

# Docker image

https://hub.docker.com/r/hauro/pod-kicker

If you dont want to edit the code, just configure the deployment.yaml file to your needs - then it works out of the box. Remember to set target repo, deployment, in the ENV vars.

# License

MIT — feel free to use, modify, and contribute! 🤘

