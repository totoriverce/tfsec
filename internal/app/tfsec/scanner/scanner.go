package scanner

import (
	"fmt"
	"github.com/tfsec/tfsec/internal/app/tfsec/metrics"
	"io/ioutil"
	"strings"



	"github.com/tfsec/tfsec/internal/app/tfsec/debug"

	"github.com/tfsec/tfsec/internal/app/tfsec/parser"
)

// Scanner scans HCL blocks by running all registered checks against them
type Scanner struct {
}

type ScannerOption int

const (
	IncludePassed ScannerOption = iota
)

// New creates a new Scanner
func New() *Scanner {
	return &Scanner{}
}

// Find element in list
func checkInList(code RuleCode, list []string) bool {
	codeCurrent := fmt.Sprintf("%s", code)
	for _, codeIgnored := range list {
		if codeIgnored == codeCurrent {
			return true
		}
	}
	return false
}

func (scanner *Scanner) Scan(blocks []*parser.Block, excludedChecksList []string, options ...ScannerOption) []Result {

	includePassed := false

	for _, option := range options {
		if option == IncludePassed {
			includePassed = true
		}
	}

	if len(blocks) == 0 {
		return nil
	}

	checkTime := metrics.Start(metrics.Check)
	defer checkTime.Stop()
	var results []Result
	context := &Context{blocks: blocks}
	checks := GetRegisteredChecks()
	for _, block := range blocks {
		for _, check := range checks {
			func(check Check) {
				if check.IsRequiredForBlock(block) {
					debug.Log("Running check for %s on %s.%s (%s)...", check.Code, block.Type(), block.FullName(), block.Range().Filename)
					var res = check.Run(block, context)
					if includePassed && res == nil {
						results = append(results, check.NewPassingResult(block.Range()))
					} else {
						for _, result := range res {
							if !scanner.checkRangeIgnored(result, result.Range) && !checkInList(result.RuleID, excludedChecksList) {
								results = append(results, result)
							}
						}
					}
				}
			}(check)
		}
	}
	return results
}

func (scanner *Scanner) checkRangeIgnored(rule Result, r parser.Range) bool {
	raw, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return false
	}

	ignoresToCheck := []string{
		"tfsec:ignore:*", // ignore all
		fmt.Sprintf("tfsec:ignore:%s", rule.RuleID), // ignore by Code
	}

	if len(rule.RuleAlias) > 0 {
		ignoreAlias := fmt.Sprintf("tfsec:ignore:%s", rule.RuleAlias)
		ignoresToCheck = append(ignoresToCheck, ignoreAlias)
	}

	lines := append([]string{""}, strings.Split(string(raw), "\n")...)
	for number := r.StartLine; number <= r.EndLine; number++ {
		if number <= 0 || number >= len(lines) {
			continue
		}

		for _, ignore := range ignoresToCheck {
			if strings.Contains(lines[number], ignore) {
				return true
			}
		}
	}

	if r.StartLine-1 > 0 {
		line := lines[r.StartLine-1]
		line = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(line, "//", ""), "#", ""))
		segments := strings.Split(line, " ")
		for _, segment := range segments {
			for _, ignore := range ignoresToCheck {
				if segment == ignore {
					return true
				}
			}
		}
	}

	return false
}
