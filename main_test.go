package main

import "testing"

func Test_reverseString(t *testing.T) {
	testInput := "hello world!"
	expectedOutput := "!dlrow olleh"

	output := reverseString(testInput)
	if output != expectedOutput {
		t.Errorf("unexpected output: %s", output)
		t.Fail()
	}
}
