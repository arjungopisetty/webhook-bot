resource "google_container_cluster" "primary" {
  name               = "chatbot-dev-1"
  location           = "us-west1"
  initial_node_count = 1

  # Enabled for GCP free tier
  enable_autopilot = true

  # master_authorized_networks_config {
  #   cidr_blocks {
  #     cidr_block = "REDACTED"
  #     display_name = "Home"
  #   }
  # }

  # private_cluster_config {
  #   enable_private_nodes = true
  #   enable_private_endpoint = false
  #   master_ipv4_cidr_block = "172.16.0.0/28"
  # }
}