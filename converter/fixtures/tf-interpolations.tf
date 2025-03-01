data "aws_iam_policy_document" "policy" {
  statement {
    sid    = "VisualEditor0"
    effect = "Allow"

    resources = [
      "arn:aws:s3:::${aws_s3_bucket.output.id}/${var.foldername}/*",
      "arn:aws:s3:::${aws_s3_bucket.output.id}",
    ]

    actions = [
      "s3:PutObject",
      "s3:GetBucketLocation",
      "s3:GetObjectAcl",
      "s3:GetObject",
      "s3:ListBucket",
      "s3:ListBucketVersions",
      "s3:DeleteObject",
      "s3:DeleteObjectVersion",
    ]
  }

  statement {
    sid    = "VisualEditor1"
    effect = "Allow"

    resources = [
      "arn:aws:s3:::${aws_s3_bucket.upload.id}/*",
      "arn:aws:s3:::${aws_s3_bucket.upload.id}",
      "arn:aws:s3:::foo/${join(var.separator,local.path_elements)}",
      "arn:aws:s3:::foo/${join("/",local.path_elements)}",
      "arn:aws:s3:::foo/${join("/",["foo${lower("BAR")}"])}",
    ]

    actions = [
      "s3:GetBucketLocation",
      "s3:GetObjectAcl",
      "s3:GetObject",
      "s3:ListBucket",
      "s3:ListBucketVersions",
      "s3:DeleteObject",
      "s3:DeleteObjectVersion",
    ]
  }

  statement {
    effect    = "Allow"
    resources = ["${aws_kms_key.key.arn}"]

    actions = [
      "kms:Encrypt",
      "kms:Decrypt",
      "kms:ReEncrypt*",
      "kms:GenerateDataKey*",
      "kms:DescribeKey",
    ]
  }

  statement {
    sid       = "AWSCloudTrailCreateLogStream2014110"
    effect    = "Allow"
    resources = ["arn:aws:logs:${data.aws_region.reg_current.name}:${data.aws_caller_identity.acc_current.account_id}:log-group:${aws_cloudwatch_log_group.trail-log-group.name}:*"]
    actions   = ["logs:CreateLogStream"]
  }
}
