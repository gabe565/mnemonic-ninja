variable "aws_account_id" {
  type        = string
  description = "AWS account ID"
}

variable "env" {
  type        = string
  description = "Environment name"
}

variable "hostname" {
  type        = string
  description = "Deployment hostname"
}

variable "acm_certificate_arn" {
  type        = string
  description = "ACM certificate ARN"
}
