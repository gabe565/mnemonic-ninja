variable "prod_hostname" {
  type        = string
  default     = "mnemonic.ninja"
  description = "Prod hostname"
}

module "prod" {
  source = "./modules/mnemonic-ninja"

  aws_account_id         = data.aws_caller_identity.current.account_id
  env                    = "prod"
  hostname               = var.prod_hostname
  acm_certificate_arn    = aws_acm_certificate.app.arn
  cloudfront_price_class = "PriceClass_200"
}

output "prod_cloudfront_hostname" {
  value       = module.prod.distribution_hostname
  description = "Prod Cloudfront hostname"
}

output "prod_hostname" {
  value       = var.prod_hostname
  description = "Prod hostname"
}
