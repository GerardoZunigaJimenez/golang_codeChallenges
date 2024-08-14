package unosquare

import (
	"fmt"
	"testing"
)

//daniel.dorado@unosquare.com
/*
package main

import "fmt"

/*

/// DESCRIPTION

Julius Caesar protected his confidential information by encrypting it using a cipher.
Caesar's cipher shifts each letter by a number of letters. For example, in the case of
a rotation by 3, w, x, y and z would map to z, a, b and c.

Original alphabet:      abcdefghijklmnopqrstuvwxyz
Alphabet rotated +3:    defghijklmnopqrstuvwxyzabc

/// REQUIREMENTS.

1. If the shift takes you past the end of the alphabet, just rotate back to the front of the alphabet.
2. If the shift number is a negative number, log an error.
3. The cipher only encrypts letters; all other characters remain unencrypted.
4. Build a function CaesarCipher(s string, n int) (string, error) to handle your implementation.
5. Create unit tests to validate previous requirements.

/// EXAMPLE.

s := "uno-square *es* %una gran empresa."
n := 3

The alphabet is rotated by 3, matching the mapping above.

The encrypted string is: "xqr-vtxduh *hv* %xqd judq hpsuhvd."

*/

func TestCipher(t *testing.T) {
	tests := []struct {
		description   string
		input         string
		rotation      int
		expected      string
		expectedError error
		validate      func(input, expected string, inputErr, expError error)
	}{
		{
			description:   "Invalid Scenario because the rotation value is negative",
			input:         "abcdefghijklmnopqrstuvwxyz",
			rotation:      -1,
			expectedError: fmt.Errorf("the shift number is a negative number, it is not possible to continue"),
			validate: func(input, expected string, inputErr, expError error) {
				if inputErr == nil {
					fmt.Println("There should to be an error, and it is not passing")
				}
				if inputErr.Error() != expError.Error() {
					fmt.Println("the result error is not the expected one", inputErr, expError)
				}
			},
		},
		{
			description: "Valid Scenario No rotations",
			input:       "abcdefghijklmnopqrstuvwxyz",
			rotation:    0,
			expected:    "abcdefghijklmnopqrstuvwxyz",
			validate: func(input, expected string, inputErr, expError error) {
				if inputErr != nil {
					fmt.Println("There is an error, and it should not", inputErr)
				}
				if !(input == expected) {
					fmt.Println("There is an error, because the result is not equals to the expected", input, expected)
				}
			},
		},
		{
			description: "Valid Scenario 3 rotations",
			input:       "uno-square *es* %una gran empresa.",
			rotation:    3,
			expected:    "xqr-vtxduh *hv* %xqd judq hpsuhvd.",
			validate: func(input, expected string, inputErr, expError error) {
				if inputErr != nil {
					fmt.Println("There is an error, and it should not", inputErr)
				}
				if !(input == expected) {
					fmt.Println("There is an error, because the result is not equals to the expected", input, expected)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			input, err := Cipher(tt.input, tt.rotation)
			tt.validate(input, tt.expected, err, tt.expectedError)
		})
	}
}
