resource "aws_iam_policy" "ci" {
  name   = "mnemonic-ninja-${var.env}-ci"
  policy = data.aws_iam_policy_document.policy.json

  tags = {
    app = "mnemonic-ninja"
    env = var.env
  }
}

resource "aws_iam_role" "ci" {
  name                 = "mnemonic-ninja-${var.env}-ci"
  assume_role_policy   = data.aws_iam_policy_document.role.json
  managed_policy_arns  = [aws_iam_policy.ci.arn]
  max_session_duration = "3600"
}

resource "aws_iam_role_policy_attachment" "ci" {
  role       = aws_iam_role.ci.name
  policy_arn = aws_iam_policy.ci.arn
}

data "aws_iam_policy_document" "policy" {
  statement {
    actions = [
      "cloudfront:CreateInvalidation",
      "s3:GetBucketLocation",
      "s3:ListBucket",
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject",
    ]

    resources = [
      aws_s3_bucket.app.arn,
      "${aws_s3_bucket.app.arn}/*",
      aws_cloudfront_distribution.app.arn,
    ]
  }
}

data "aws_iam_policy_document" "role" {
  statement {
    actions = [
      "sts:AssumeRoleWithWebIdentity",
    ]

    condition {
      test     = "StringEquals"
      variable = "token.actions.githubusercontent.com:aud"
      values   = ["sts.amazonaws.com"]
    }

    condition {
      test     = "StringLike"
      variable = "token.actions.githubusercontent.com:sub"
      values   = ["repo:gabe565/mnemonic-ninja:environment:${var.env}"]
    }

    principals {
      type = "Federated"
      identifiers = [
        "arn:aws:iam::${var.aws_account_id}:oidc-provider/token.actions.githubusercontent.com"
      ]
    }
  }
}
