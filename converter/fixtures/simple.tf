data "aws_iam_policy_document" "policy" {
  statement {
    sid       = "AllowAllUsersToListAccounts"
    effect    = "Allow"
    resources = ["*"]

    actions = [
      "iam:ListAccountAliases",
      "iam:ListUsers",
      "iam:ListVirtualMFADevices",
      "iam:GetAccountPasswordPolicy",
      "iam:GetAccountSummary",
    ]
  }

  statement {
    sid    = "AllowIndividualUserToListOnlyTheirOwnMFA"
    effect = "Allow"

    resources = [
      "arn:aws:iam::*:mfa/*",
      "arn:aws:iam::*:user/$${aws:username}",
    ]

    actions = ["iam:ListMFADevices"]
  }

  statement {
    sid       = "BlockMostAccessUnlessSignedInWithMFA"
    effect    = "Deny"
    resources = ["*"]

    not_actions = [
      "iam:CreateVirtualMFADevice",
      "iam:DeleteVirtualMFADevice",
      "iam:ListVirtualMFADevices",
      "iam:EnableMFADevice",
      "iam:ResyncMFADevice",
      "iam:ListAccountAliases",
      "iam:ListUsers",
      "iam:ListSSHPublicKeys",
      "iam:ListAccessKeys",
      "iam:ListServiceSpecificCredentials",
      "iam:ListMFADevices",
      "iam:GetAccountSummary",
      "sts:GetSessionToken",
    ]

    condition {
      test     = "BoolIfExists"
      variable = "aws:MultiFactorAuthPresent"
      values   = ["false"]
    }
  }

  statement {
    sid    = ""
    effect = "Deny"

    not_resources = [
      "arn:aws:s3:::HRBucket/Payroll",
      "arn:aws:s3:::HRBucket/Payroll/*",
    ]

    actions = ["s3:*"]

    condition {
      test     = "StringLike"
      variable = "s3:prefix"

      values = [
        "",
        "home/",
        "home/$${aws:username}/",
      ]
    }
  }

  statement {
    sid     = ""
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type = "Service"

      identifiers = [
        "elasticmapreduce.amazonaws.com",
        "datapipeline.amazonaws.com",
      ]
    }
  }

  statement {
    sid     = ""
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    not_principals {
      type        = "AWS"
      identifiers = ["arn:aws:sts::1234567:assumed-role/role-name/role-session-name"]
    }
  }

  statement {
    sid       = "DenyUnEncryptedObjectUploads"
    effect    = "Deny"
    resources = ["arn:aws:s3:::<bucket_name>/*"]
    actions   = ["s3:PutObject"]

    condition {
      test     = "Null"
      variable = "s3:x-amz-server-side-encryption"
      values   = ["true"]
    }

    principals {
      type        = "*"
      identifiers = ["*"]
    }
  }
}
