🚀 Pod Kicker
Pod Kicker is a lightweight web server written in Go, designed to automatically restart existing Kubernetes deployments whenever a new image is pushed to Docker Hub.

🔧 Prerequisites
To get Pod Kicker working, make sure the following are in place:

GitHub Actions CI/CD pipeline
Your application’s GitHub repository should include a workflow that builds and pushes your Docker image to Docker Hub.
🐙⚙️

Webhook on Docker Hub
Configure a webhook to trigger when a new image is pushed.
🔁🐳

That's all it takes! 🎉

📦 Running Pod Kicker in Kubernetes
You can run Pod Kicker in your Kubernetes cluster in one of two ways:

Build your own image
Clone the repo and build your own Docker image — customize as needed.

Use the latest pre-built image
Pull the latest from: XXXXXXXXX
🧪

Pod Kicker listens on port 8450 by default.

🔁 Available Endpoint
POST /KickThatPod
Triggers a rollout restart of the target deployment.

More endpoints coming soon! 😺

🔐 RBAC Permissions
Make sure your Pod Kicker deployment has the necessary RBAC permissions to restart deployments within your cluster.

You’ll need to grant it access to:

get, list, and patch on deployments in the appropriate namespace.

📄 License
MIT — feel free to use, modify, and contribute! 🤘
