data "aws_iam_policy_document" "deny_access_without_mfa" {
  statement {
    sid       = "BlockMostAccessUnlessSignedInWithMFA"
    effect    = "Deny"
    resources = ["*"]
  }
}
