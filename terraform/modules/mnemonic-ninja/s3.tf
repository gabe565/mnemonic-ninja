resource "aws_s3_bucket" "app" {
  bucket = var.hostname

  tags = {
    app = "mnemonic-ninja"
    env = var.env
  }

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_s3_bucket_policy" "app" {
  bucket = aws_s3_bucket.app.id
  policy = data.aws_iam_policy_document.cloudfront_to_s3.json
}

data "aws_iam_policy_document" "cloudfront_to_s3" {
  policy_id = "PolicyForCloudFrontPrivateContent"
  statement {
    sid = "AllowCloudFrontServicePrincipal"
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }
    actions = [
      "s3:ListBucket",
      "s3:GetObject",
    ]
    resources = [
      aws_s3_bucket.app.arn,
      "${aws_s3_bucket.app.arn}/*",
    ]

    condition {
      test     = "StringEquals"
      variable = "AWS:SourceArn"

      values = [
        aws_cloudfront_distribution.app.arn
      ]
    }
  }
}
