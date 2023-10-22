package cmd

import (
	"bytes"
	//"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootCmd_Execute(t *testing.T) {

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	RootCmd.SetOut(&stdout)

	// Execute the root command
	err := RootCmd.Execute()

	// Check for any execution error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the output
	expectedOutput := "Welcome to DLP's CLI!\n\n"
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, stdout.String())
	}
}
