# Mnemonic Ninja Terraform

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.2.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.2.0 |
| <a name="provider_aws.us-east-1"></a> [aws.us-east-1](#provider\_aws.us-east-1) | 5.2.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_dev"></a> [dev](#module\_dev) | ./modules/mnemonic-ninja | n/a |
| <a name="module_prod"></a> [prod](#module\_prod) | ./modules/mnemonic-ninja | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_acm_certificate.app](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate) | resource |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_dev_hostname"></a> [dev\_hostname](#input\_dev\_hostname) | Dev hostname | `string` | `"dev.mnemonic.ninja"` | no |
| <a name="input_prod_hostname"></a> [prod\_hostname](#input\_prod\_hostname) | Prod hostname | `string` | `"mnemonic.ninja"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_cert_validation_records"></a> [cert\_validation\_records](#output\_cert\_validation\_records) | ACM certificate validation DNS records |
| <a name="output_dev_cloudfront_hostname"></a> [dev\_cloudfront\_hostname](#output\_dev\_cloudfront\_hostname) | Dev Cloudfront hostname |
| <a name="output_dev_hostname"></a> [dev\_hostname](#output\_dev\_hostname) | Dev hostname |
| <a name="output_prod_cloudfront_hostname"></a> [prod\_cloudfront\_hostname](#output\_prod\_cloudfront\_hostname) | Prod Cloudfront hostname |
| <a name="output_prod_hostname"></a> [prod\_hostname](#output\_prod\_hostname) | Prod hostname |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
