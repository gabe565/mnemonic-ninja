resource "aws_acm_certificate" "app" {
  provider = aws.us-east-1

  domain_name       = "mnemonic.ninja"
  validation_method = "DNS"

  subject_alternative_names = ["*.mnemonic.ninja"]

  tags = {
    app = "mnemonic-ninja"
  }

  lifecycle {
    create_before_destroy = true
  }
}

output "cert_validation_records" {
  value       = aws_acm_certificate.app.domain_validation_options
  description = "ACM certificate validation DNS records"
}
