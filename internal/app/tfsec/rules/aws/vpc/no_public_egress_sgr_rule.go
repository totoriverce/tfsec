package vpc

// generator-locked
import (
	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/severity"

	"github.com/aquasecurity/tfsec/pkg/provider"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/cidr"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/hclcontext"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"

	"github.com/aquasecurity/tfsec/pkg/rule"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID:  "AWS007",
		Service:   "vpc",
		ShortCode: "no-public-egress-sgr",
		Documentation: rule.RuleDocumentation{
			Summary: "An egress security group rule allows traffic to /0.",
			Explanation: `
Opening up ports to connect out to the public internet is generally to be avoided. You should restrict access to IP addresses or ranges that are explicitly required where possible.
`,
			Impact:     "Your port is egressing data to the internet",
			Resolution: "Set a more restrictive cidr range",
			BadExample: []string{`
resource "aws_security_group_rule" "bad_example" {
	type = "egress"
	cidr_blocks = ["0.0.0.0/0"]
}
`},
			GoodExample: []string{`
resource "aws_security_group_rule" "good_example" {
	type = "egress"
	cidr_blocks = ["10.0.0.0/16"]
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule",
			},
		},
		Provider:        provider.AWSProvider,
		RequiredTypes:   []string{"resource"},
		RequiredLabels:  []string{"aws_security_group_rule"},
		DefaultSeverity: severity.Critical,
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context) {

			typeAttr := resourceBlock.GetAttribute("type")
			if typeAttr.IsNil() || !typeAttr.IsString() || typeAttr.NotEqual("egress") {
				return
			}

			if cidrBlocksAttr := resourceBlock.GetAttribute("cidr_blocks"); cidrBlocksAttr.IsNotNil() {

				if cidr.IsAttributeOpen(cidrBlocksAttr) {
					set.AddResult().
						WithDescription("Resource '%s' defines a fully open egress security group rule.", resourceBlock.FullName()).
						WithAttribute(cidrBlocksAttr)
				}
			}

			if ipv6CidrBlocksAttr := resourceBlock.GetAttribute("ipv6_cidr_blocks"); ipv6CidrBlocksAttr.IsNotNil() {

				if cidr.IsAttributeOpen(ipv6CidrBlocksAttr) {
					set.AddResult().
						WithDescription("Resource '%s' defines a fully open egress security group rule.", resourceBlock.FullName()).
						WithAttribute(ipv6CidrBlocksAttr)
				}
			}
		},
	})
}
