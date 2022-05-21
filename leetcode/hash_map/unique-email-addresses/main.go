package main

import "strings"

const (
	local int = iota
	domain
)

func numUniqueEmails(emails []string) int {
	emailTable := make(map[string]struct{}, len(emails))
	for _, email := range emails {
		e := validateEmail(email)
		if _, exists := emailTable[e]; !exists {
			emailTable[e] = struct{}{}
		}
	}

	return len(emailTable)
}

func validateEmail(email string) string {
	ss := strings.Split(email, "@")
	ss[local] = strings.ReplaceAll(ss[local], ".", "")
	if strings.Contains(ss[local], "+") {
		ss[local] = strings.Split(ss[local], "+")[0]
	}

	return ss[local] + "@" + ss[domain]
}
