resource "aws_instance" "traefik" {
  ami           = data.aws_ami.ubuntu_jammy.id
  instance_type = "m6a.large"
  subnet_id     = aws_subnet.public_us_east_1a.id
  key_name      = "devops" # TODO: update to yours

  vpc_security_group_ids = [
    aws_security_group.proxy.id,
    aws_security_group.ssh.id
  ]

  tags = {
    Name          = "traefik"
    service       = "traefik"
    node-exporter = "true"
  }
}

output "traefik_public_ip" {
  value = aws_instance.traefik.public_ip
}
