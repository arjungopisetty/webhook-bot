output "kubernetes_cluster_name" {
  value       = google_container_cluster.primary.name
  description = "GKE Cluster Name"
}

output "kubernetes_cluster_host" {
  value       = google_container_cluster.primary.endpoint
  description = "GKE Cluster Host"
}

output "kubernetes_cluster_url" {
  value       = google_container_cluster.primary.self_link
  description = "GKE Cluster URL"
}

output "kubernetes_cluster_id" {
  value       = google_container_cluster.primary.id
  description = "GKE Cluster ID"
}
