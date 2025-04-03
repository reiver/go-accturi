package accturi_test

import (
	"testing"

	"github.com/reiver/go-accturi"
)

func TestIsAcctURIBytes(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected bool
	}{
		{
			Value: []byte(""),
			Expected: false,
		},



		{
			Value: []byte("apple"),
			Expected: false,
		},
		{
			Value: []byte("banana"),
			Expected: false,
		},
		{
			Value: []byte("cherry"),
			Expected: false,
		},



		{
			Value: []byte("http://example.com/"),
			Expected: false,
		},
		{
			Value: []byte("http://joeblow@example.com/"),
			Expected: false,
		},



		{
			Value: []byte("a"),
			Expected: false,
		},
		{
			Value: []byte("ac"),
			Expected: false,
		},
		{
			Value: []byte("acc"),
			Expected: false,
		},
		{
			Value: []byte("acct"),
			Expected: false,
		},
		{
			Value: []byte("acct:"),
			Expected: false,
		},



		{
			Value: []byte("acct:a"),
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := accturi.IsAcctURIBytes(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-acct-uri-bytes is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}

func TestIsAcctURIString(t *testing.T) {

	tests := []struct{
		Value string
		Expected bool
	}{
		{
			Value: "",
			Expected: false,
		},



		{
			Value: "apple",
			Expected: false,
		},
		{
			Value: "banana",
			Expected: false,
		},
		{
			Value: "cherry",
			Expected: false,
		},



		{
			Value: "http://example.com/",
			Expected: false,
		},
		{
			Value: "http://joeblow@example.com/",
			Expected: false,
		},



		{
			Value: "a",
			Expected: false,
		},
		{
			Value: "ac",
			Expected: false,
		},
		{
			Value: "acc",
			Expected: false,
		},
		{
			Value: "acct",
			Expected: false,
		},
		{
			Value: "acct:",
			Expected: false,
		},



		{
			Value: "acct:a",
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := accturi.IsAcctURIString(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual is-acct-uri-string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
