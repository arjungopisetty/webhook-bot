# webhook-bot
REST API that, when sent a GitLab webhook, sends the git commit SHA to a Rocket Chat user as a direct message. I found this challenge on Reddit and wanted to attempt it as a learning experience.

## Layout 
k8s-manifests/
- Manifests for Rocket Chat, GitLab
webhook-app/
- GitLab webhook -> Rocket Chat message application
monitoring/
- prometheus-grafana-values.yaml
- webhook-app-service-monitor.yaml
- webhook-app-dashboard.yaml
- prometheus-metrics-adaptor(-values).yaml
 
## Overview
Create and deploy k8s manifests that contain the following applications using publicly available images from Docker Hub:
- Rocket Chat
- GitLab

Do not use a manifest package framework (e.g. Helm) for this step.

Write a REST API application that, when sent a GitLab webhook, sends the git commit SHA to a Rocket Chat user as a direct message (language of your choosing). Push this to a project on your GitLab instance created above.
Write a Dockerfile and k8s manifest(s) for that application and configure GitLab CI such that it builds and pushes the image to the bundled GitLab container registry, then deploys to your k8s cluster.
Configure GitLab CI to send a webhook to your deployed application, triggered by commits to the application’s GitLab project.
 
### Monitoring and Custom Metrics
Deploy the Prometheus and Grafana to your k8s cluster using the kube-prometheus-stack Helm chart, saving settings to a `values.yaml`.
Add a Prometheus metrics endpoint to your REST API that exposes a count of received webhooks.
Configure Prometheus to scrape that endpoint using ServiceMonitor, and create a provisioned Grafana dashboard for the metric.
Create another custom metric for your REST API that exposes a count of in-flight requests.
Deploy a custom metrics adaptor for Prometheus, saving your k8s manifest(s) and/or chart `values.yaml`.
Add an HPA manifest to your REST API that’s sensitive to the custom metric.
