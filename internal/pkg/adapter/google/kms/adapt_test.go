package kms

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/google/kms"
	"github.com/aquasecurity/tfsec/internal/pkg/adapter/testutils"
)

func Test_Adapt(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name      string
		terraform string
		expected  kms.KMS
	}{
		{
			name: "basic",
			terraform: `
resource "" "example" {
    
}
`,
			expected: kms.KMS{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := testutils.CreateModulesFromSource(test.terraform, ".tf", t)
			adapted := Adapt(modules)
			testutils.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

func Test_adaptKeyRings(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name      string
		terraform string
		expected  []kms.KeyRing
	}{
		{
			name: "basic",
			terraform: `
resource "" "example" {
    
}
`,
			expected: []kms.KeyRing{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := testutils.CreateModulesFromSource(test.terraform, ".tf", t)
			adapted := adaptKeyRings(modules)
			testutils.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

func Test_adaptKey(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name      string
		terraform string
		expected  kms.Key
	}{
		{
			name: "basic",
			terraform: `
resource "" "example" {
    
}
`,
			expected: kms.Key{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := testutils.CreateModulesFromSource(test.terraform, ".tf", t)
			adapted := adaptKey(modules.GetBlocks()[0])
			testutils.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}
