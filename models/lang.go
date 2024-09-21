package models

import (
	"fmt"
)

type Lang int

const (
	DE Lang = iota
	EN
)

func (l Lang) String() string {
	return [...]string{"deu", "eng"}[l]
}

func (w Lang) EnumIndex() int {
	return int(w)
}

func ParseLangString(langStr string) (Lang, error) {
	if langStr == "" {
		return EN, nil
	}

	switch langStr {
	case DE.String():
		return DE, nil
	case EN.String():
		return EN, nil
	default:
		return -1, fmt.Errorf("invalid 'lang' parameter: %s", langStr)
	}
}
