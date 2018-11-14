package main

import (
	"fmt"
	"os"
	"testing"
)

var testTax tax

// TestMain is called first. see additional: https://golang.org/pkg/testing/#hdr-Main
func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	// teardown()

	os.Exit(exitCode)
}

func setup() {
	fmt.Println("called setup()")
	testTax = newTax()
}

func understandForHelper(t *testing.T) {
	t.Helper() // Helper marks the calling function as a test helper function
	t.Fatal("failure here")
}

func TestTaxExcludeAmount(t *testing.T) {

	// testCase として、割り切れる場合・切り上げる場合・切り下げる場合を用意する
	testCase := []struct {
		testName string
		input    int
		want     int
	}{
		{
			testName: "divided",
			input:    2160,
			want:     2000,
		},
		{
			testName: "roundUp",
			input:    2005,
			want:     1856,
		},
		{
			testName: "roundDown",
			input:    2000,
			want:     1852,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.testName, func(t *testing.T) {
			got := testTax.calcurateTaxExcludeAmount(tc.input)
			if got != tc.want {
				t.Fatalf("want %v, but %v:", tc.want, got)
			}
			t.Log("logger at ", t.Name())
		}) //END table driven test
	}

	// understandForHelper(t)
}
