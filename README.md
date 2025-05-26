🚀 Pod Kicker is a lightweight web server built in Go, designed to automatically restart existing Kubernetes deployments whenever a new image is pushed to Docker Hub.

To get Pod Kicker up and running, make sure you’ve got the following in place:

1, GitHub Actions CI/CD pipeline in your application’s GitHub repo — it should be configured to build and push your image to Docker Hub.
2. A webhook set up on Docker Hub to trigger whenever a new image is pushed. 🔁🐳

That’s it! 🎉

To run Pod Kicker inside Kubernetes:

Build and package it into a Docker image (feel free to customize the code as needed).
Or, use the latest pre-built image from XXXXXXXXX 🧪.

By default it uses port 8450, and as of now only has /KickThatPod post endpoint (more will come :3)

⚠️ Don’t forget to grant RBAC permissions to the Pod Kicker deployment, so it has the rights to restart your Kubernetes workloads. 🔐!
