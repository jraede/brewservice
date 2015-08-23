package main

import (
	"github.com/mgutz/ansi"
	"os/exec"
	"regexp"
	"strings"
)

var match = regexp.MustCompile("homebrew(\\.([A-Za-z0-9]+))+")

func getStatus() (map[string]string, error) {
	cache := loadCacheOrFail()

	output, err := exec.Command("launchctl", "list").Output()
	if err != nil {
		return nil, err
	}

	matches := match.FindAll(output, -1)

	statuses := make(map[string]string)

	for serviceName, plistPath := range cache {
		for _, m := range matches {
			if strings.Contains(string(plistPath), string(m)) {
				statuses[serviceName] = ansi.Green + "✓" + ansi.Reset
				break
			}

		}

		if _, ok := statuses[serviceName]; !ok {
			statuses[serviceName] = ansi.Red + "×" + ansi.Reset
		}
	}

	return statuses, nil
}
