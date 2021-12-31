package s3

import (
	"github.com/aquasecurity/defsec/rules/aws/s3"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "AWS076",
		BadExample: []string{`
resource "aws_s3_bucket" "example" {
  bucket = "mybucket"
}

resource "aws_s3_bucket_public_access_block" "bad_example" {
  bucket = aws_s3_bucket.example.id
}
 
resource "aws_s3_bucket_public_access_block" "bad_example" {
  bucket = aws_s3_bucket.example.id 
  block_public_policy = false
}
 `},
		GoodExample: []string{`
resource "aws_s3_bucket" "example" {
  bucket = "mybucket"
}

resource "aws_s3_bucket_public_access_block" "good_example" {
  bucket = aws_s3_bucket.example.id 
  block_public_policy = true 
}
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_public_access_block#block_public_policy",
		},
		Base: s3.CheckPublicPoliciesAreBlocked,
	})
}
