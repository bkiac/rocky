package main

import "strings"

const (
	separator = ", "
)

func Split(s string) []string {
	return strings.Split(s, separator)
}

func Join(s []string) string {
	return strings.Join(s, separator)
}
