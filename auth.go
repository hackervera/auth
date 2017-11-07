package auth

import (
	"net"
	"regexp"
)

func FindKeys(domain string) ([]string, error) {
	re := regexp.MustCompile("ed25519=(.*)")
	records, err := net.LookupTXT(domain)
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, record := range records {
		res := re.FindStringSubmatch(record)
		if res != nil {
			keys = append(keys, res[1])
		}
	}
	return keys, nil
}
