package usecase

import (
	"errors"
	"strings"
)

type packageMonitored struct {
	Values []string
}

func NewPackageMonitored(packagesString string) (*packageMonitored, error) {
	packagesString = strings.TrimSpace(packagesString)
	if packagesString == "" {
		return nil, errors.New("packages monitored is empty")
	}

	packagesList := strings.Split(packagesString, ",")
	var packages packageMonitored

	for _, p := range packagesList {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		packages.Values = append(packages.Values, p)
	}

	if len(packages.Values) == 0 {
		return nil, errors.New("packages monitored is empty")
	}
	return &packages, nil
}

func (p *packageMonitored) Contains(packageName string) bool {
	for _, p := range p.Values {
		if p == packageName {
			return true
		}
	}
	return false
}
