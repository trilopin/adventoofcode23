package main

import (
	"strings"
	"testing"
)

func TestCalibrateBlock(t *testing.T) {
	testCases := []struct {
		Input  string
		Output int
	}{
		// Part One
		{"1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet", 142},
		{"3abc2\npqr3stu8vwx\na1b2c3d4e5f\n7reb7uchet", 162},
		// Part Two
		{"two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen", 281},
	}
	for _, tc := range testCases {
		got := calibrateBlock(strings.NewReader(tc.Input))
		if got != tc.Output {
			t.Errorf("calibrateBlock(%q) = %d; want %d", tc.Input, got, tc.Output)
		}
	}
}

func TestCalibrateLine(t *testing.T) {
	testCases := []struct {
		Input  string
		Output int
	}{
		// Part One
		{"trrmhcmto1rpkb27fh", 17},
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		// Part Two
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}
	for _, tc := range testCases {
		got := calibrateLine(tc.Input)
		if got != tc.Output {
			t.Errorf("extractCalibration(%q) = %d; want %d", tc.Input, got, tc.Output)
		}
	}
}

func TestIsNumberContained(t *testing.T) {
	testCases := []struct {
		Input  string
		Output string
	}{
		{"sdgsfdgonedsvfggvds", "1"},
		{"nine", "9"},
		{"loltwoppp", "2"},
		{"threeman", "3"},
		{"tadfghsdfheeman", ""},
	}
	for _, tc := range testCases {
		_, got := isNumberContained(tc.Input, true)
		if got != tc.Output {
			t.Errorf("isNumberContained(%q) = %s; want %s", tc.Input, got, tc.Output)
		}
		_, got = isNumberContained(tc.Input, false)
		if got != tc.Output {
			t.Errorf("isNumberContained(%q) = %s; want %s", tc.Input, got, tc.Output)
		}
	}
}
