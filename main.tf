provider "aws" {
  region = "ap-south-1"
}

# Reference existing key pair by name
resource "aws_key_pair" "deployer" {
  key_name   = "muhammedsirajudeen"
  public_key = file("~/.ssh/id_rsa.pub")
}

# Reference existing security group by name (SG name must match exactly)
resource "aws_security_group" "ec2_sg" {
  name        = "ec2_sg"
  description = "Allow SSH, HTTP, and HTTPS traffic"
  vpc_id      = "vpc-0cb5308c845901ab7"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# EC2 Instance
resource "aws_instance" "web" {
  ami             = "ami-0f5ee92e2d63afc18"
  instance_type   = "t2.micro"
  key_name        = aws_key_pair.deployer.key_name
  security_groups = [aws_security_group.ec2_sg.name]

  root_block_device {
    volume_size           = 28
    volume_type           = "gp3"
    delete_on_termination = true
  }

  tags = {
    Name = "UbuntuInstance"
  }
}

output "instance_public_ip" {
  description = "Public IP of the EC2 instance"
  value       = aws_instance.web.public_ip
}

output "instance_id" {
  description = "ID of the EC2 instance"
  value       = aws_instance.web.id
}
