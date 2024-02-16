output "vpc_connector_name" {
  value = google_vpc_access_connector.vpc-connector.name
}

output "vpc_connector_ip_range" {
  value = google_vpc_access_connector.vpc-connector.ip_cidr_range
}