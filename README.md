# Marvel Webhook for Kubernetes 🦸‍♂️🦸‍♀️
![Go](https://img.shields.io/badge/Go-1.20.4-blue)
![Kubernetes](https://img.shields.io/badge/Kubernetes-1.27.4-blue)
![Cert-Manager](https://img.shields.io/badge/Cert--Manager-1.13.1-green)
![License](https://img.shields.io/badge/License-MIT-purple)

<img src="https://github.com/kubernetes/kubernetes/raw/master/logo/logo.png" width="100">

# Introduction
Marvel Webhook is a Kubernetes Mutating Admission Webhook that adds a Marvel superhero name as a label to every new pod. Built with ❤️ in Go, it's a fun way to add some Marvel magic to your Kubernetes cluster!

# Features
🦸‍♂️ Adds a random Marvel superhero name as a label to new pods

🧪 Includes unit test

🛡️ Secure TLS communication using Cert-Manager

# Prerequisites
- Kubernetes 1.18+
- Go 1.16+
- Cert-Manager 1.3.1+
- Marvel API keys

# Installation
## Cert-Manager Setup
Navigate to the deploy/cert-manager directory and apply all YAML files:

```bash
kubectl apply -f deploy/cert-manager/
```

## Build the webhook (optional)
docker build -t marvel-webhook:<your_tag> -f cmd/marvel-webhook/Dockerfile .

# Deploy the Webhook
```bash
kubectl apply -f deploy/webhook-configuration/marvel-webhook.yaml
```

# Usage
Once installed, every new pod will automatically receive a Marvel superhero name as a label. To check the labels, run:

```bash
kubectl get pods --show-labels
```

# Development

## Running Tests
Navigate to the cmd/marvel-webhook directory and run:

```bash
export MARVEL_PRIVATE_KEY=<your_private_key>
cd cmd/marvel-webhook/ && go test -v
```

# License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

