package main

import (
	"fmt"
	"os"
	"strings"
)

func getBrand(number string, brands map[string]string) string {
	if brands == nil {
		fmt.Fprintln(os.Stderr, "Brands map is nil")
		return "-"
	}
	var best string
	for prefix, brand := range brands {
		if strings.HasPrefix(number, prefix) && len(prefix) > len(best) {
			best, _ = prefix, brand
		}
	}
	if best != "" {
		return brands[best]
	}
	return "-"
}

func getIssuer(number string, issuers map[string]string) string {
	if issuers == nil {
		fmt.Fprintln(os.Stderr, "Issuers map is nil")
		return "-"
	}
	var best string
	for prefix, issuer := range issuers {
		if strings.HasPrefix(number, prefix) && len(prefix) > len(best) {
			best, _ = prefix, issuer
		}
	}
	if best != "" {
		return issuers[best]
	}
	return "-"
}
