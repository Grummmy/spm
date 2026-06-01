package main

import (
	"regexp"
	"slices"
	"strings"
)

type Matcher struct {
	Full  []string `toml:"full,omitempty"`
	Start []string `toml:"start,omitempty"`
	End   []string `toml:"end,omitepmty"`
	Regex []string `toml:"regex,omitempty"`
}

var regexCache map[string]*regexp.Regexp

// check if string can be matched to matcher
func (m Matcher) Match(name string) bool {
	// check if name matchs any full string in matcher
	if slices.Contains(m.Full, name) {
		return true
	}

	// check if name starts with any string in matcher
	for _, str := range m.Start {
		if strings.HasPrefix(name, str) {
			return true
		}
	}

	// check if name ends with any string in matcher
	for _, str := range m.End {
		if strings.HasSuffix(name, str) {
			return true
		}
	}

	// check if name matchs any regex in matcher
	for _, str := range m.Regex {
		if _, ok := regexCache[str]; !ok {
			// cache regex for fastel iteration
			r, err := regexp.Compile(str)
			if err != nil {
				logger.Error("Could not compile '"+str+"' regex, skipping it:", err)
				continue
			}

			regexCache[str] = r
		}

		if regexCache[str].MatchString(name) {
			return true
		}
	}

	return false
}
