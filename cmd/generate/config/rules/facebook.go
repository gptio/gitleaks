package rules

import (
	"github.com/gptio/gitleaks/v8/cmd/generate/secrets"
	"github.com/gptio/gitleaks/v8/config"
)

func Facebook() *config.Rule {
	// define rule
	r := config.Rule{
		Description: "Facebook Access Token",
		RuleID:      "facebook",
		Regex:       generateSemiGenericRegex([]string{"facebook"}, hex("32")),
		SecretGroup: 1,
		Keywords:    []string{"facebook"},
	}

	// validate
	tps := []string{
		generateSampleSecret("facebook", secrets.NewSecret(hex("32"))),
	}
	return validate(r, tps, nil)
}
