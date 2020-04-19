data "aws_iam_policy_document" "policy" {
  statement {
    sid       = ""
    effect    = "Deny"
    resources = ["*"]
    actions   = ["*"]

    condition {
      test     = "NotIpAddress"
      variable = "aws:SourceIp"

      values = [
        "xxx.xxx.xxx.xxx/xx",
        "yyy.yyy.yyy.yyy",
      ]
    }

    condition {
      test     = "StringNotLike"
      variable = "aws:invokedBy"
      values   = ["*.amazonaws.com"]
    }
  }
}
