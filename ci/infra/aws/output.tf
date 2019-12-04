output "control_plane_public_ip" {
  value = aws_instance.control_plane.*.public_ip
}

output "control_plane_private_dns" {
  value = aws_instance.control_plane.*.private_dns
}

output "nodes_private_dns" {
  value = aws_instance.nodes.*.private_dns
}

