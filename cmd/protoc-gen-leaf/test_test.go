package main

import (
	"regexp"
	"testing"
)

func TestParse(t *testing.T) {
	re := regexp.MustCompile(`\d+`)
	str := "/request/100/response/101"
	match := re.FindAllString(str,2)
	t.Log(match)
}
