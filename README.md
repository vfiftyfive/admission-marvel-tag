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
🛡️ Secure TLS communication using Cert-Manager
📦 Easy to install and configure
📝 Well-documented codebase
🧪 Includes unit tests
# Prerequisites
- Kubernetes 1.18+
- Go 1.16+
- Cert-Manager 1.3.1+

# Installation
Cert-Manager Setup
Navigate to the deploy/cert-manager directory and apply all YAML files:

```bash
kubectl apply -f deploy/cert-manager/
```

# Webhook Configuration
Apply the MutatingWebhookConfiguration:

```bash
kubectl apply -f deploy/webhook-configuration/marvel-webhook.yaml
```

# Deploy the Webhook
Coming soon!

# Usage
Once installed, every new pod will automatically receive a Marvel superhero name as a label. To check the labels, run:

```bash
kubectl get pods --show-labels
```

# Development

## Running Tests
Navigate to the cmd/marvel-webhook directory and run:

```bash
go test
```

# Contributing
We welcome contributions! Please see CONTRIBUTING.md for details.

# License
This project is licensed under the MIT License - see the LICENSE.md file for details.

