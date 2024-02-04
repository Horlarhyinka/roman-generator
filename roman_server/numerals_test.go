package main

import (
	"fmt"
	"romanNumerals"
	// "strconv"
	"testing"
)

func TestGetRomanNumeral(t *testing.T) {
	type testAnswerMap = map[int]string
	testSuites := []testAnswerMap{
		{20: "XX"},
		{21: "XXI"},
		{22: "XXII"},
		{23: "XXIII"},
		{24: "XXIV"},
		{25: "XXV"},
		{26: "XXVI"},
		{27: "XXVII"},
		{28: "XXVIII"},
		{29: "XXIX"},
		{30: "XXX"},
		{51: "LI"},
		{101: "CI"},
		{110: "CX"},
		{111: "CXI"},
		{1001: "MI"},
		{2000: "MM"},
		{0: ""},
	}
	for key, val := range romanNumerals.Numerals{
		testSuites = append(testSuites, testAnswerMap{key: val})
	}
	//
	for _, suite := range testSuites{
		for k, v := range suite{
			r := GetRomanNumeral(k)
			if v != r{
				t.Errorf("expected GetRomanNumeral(%d) to be (%q), recieved %q", k, v, r)
			}else {
				fmt.Printf("%-10d --> %+15q  %-10s \n", k, r, "OK")
			}
		}
	}
}
