package main

import (
	"fmt"
	"regexp"
)

var rulesMatcher = regexp.MustCompile(`(\d+)|(\d+)`)

func buildRulesMap(rules []string, debug bool) map[string][]string {
	rulesMap := make(map[string][]string)

	for i, unparsedRule := range rules {
		match := rulesMatcher.FindAllStringSubmatch(unparsedRule, -1)
		if match == nil {
			msg := fmt.Sprintf("Unable to parse rule: %s (line %d)", unparsedRule, i)
			panic(msg)
		}

		// x, y := match[0], match[1]
		x := match[0][1]
		y := match[1][1]

		if debug {
			fmt.Printf("* %v\n", match)
		}

		var rulesForX []string
		if len(rulesMap[x]) > 0 {
			rulesForX = append(rulesMap[x], y)
		} else {
			rulesForX = []string{y}
		}
		rulesMap[x] = rulesForX
	}

	return rulesMap
}
