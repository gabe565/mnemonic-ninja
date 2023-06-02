variable "dev_hostname" {
  type        = string
  default     = "dev.mnemonic.ninja"
  description = "Dev hostname"
}

module "dev" {
  source = "./modules/mnemonic-ninja"

  aws_account_id      = data.aws_caller_identity.current.account_id
  env                 = "dev"
  hostname            = var.dev_hostname
  acm_certificate_arn = aws_acm_certificate.app.arn
}

output "dev_cloudfront_hostname" {
  value       = module.dev.distribution_hostname
  description = "Dev Cloudfront hostname"
}

output "dev_hostname" {
  value       = var.dev_hostname
  description = "Dev hostname"
}
