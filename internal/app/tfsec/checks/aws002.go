package checks

import (
	"fmt"

	"github.com/tfsec/tfsec/internal/app/tfsec/scanner"

	"github.com/tfsec/tfsec/internal/app/tfsec/parser"
)

const AWSNoBucketLogging scanner.RuleCode = "AWS002"
const AWSNoBucketLoggingDescription scanner.RuleSummary = "S3 Bucket does not have logging enabled."
const AWSNoBucketLoggingImpact = "There is no way to determine the access to this bucket"
const AWSNoBucketLoggingResolution = "Add a logging block to the resource to enable access logging"
const AWSNoBucketLoggingExplanation = `
Buckets should have logging enabled so that access can be audited. 
`
const AWSNoBucketLoggingBadExample = `
resource "aws_s3_bucket" "bad_example" {

}
`
const AWSNoBucketLoggingGoodExample = `
resource "aws_s3_bucket" "good_example" {
	logging {
		target_bucket = "target-bucket"
	}
}
`

func init() {
	scanner.RegisterCheck(scanner.Check{
		Code: AWSNoBucketLogging,
		Documentation: scanner.CheckDocumentation{
			Summary:     AWSNoBucketLoggingDescription,
			Explanation: AWSNoBucketLoggingExplanation,
			Impact:      AWSNoBucketLoggingImpact,
			Resolution:  AWSNoBucketLoggingResolution,
			BadExample:  AWSNoBucketLoggingBadExample,
			GoodExample: AWSNoBucketLoggingGoodExample,
			Links: []string{
				"https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html",
				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket",
			},
		},
		Provider:       scanner.AWSProvider,
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_s3_bucket"},
		CheckFunc: func(check *scanner.Check, block *parser.Block, _ *scanner.Context) []scanner.Result {
			if loggingBlock := block.GetBlock("logging"); loggingBlock == nil {
				if block.GetAttribute("acl") != nil && block.GetAttribute("acl").Equals("log-delivery-write") {
					return nil
				}
				return []scanner.Result{
					check.NewResult(
						fmt.Sprintf("Resource '%s' does not have logging enabled.", block.FullName()),
						block.Range(),
						scanner.SeverityError,
					),
				}
			}
			return nil
		},
	})
}
