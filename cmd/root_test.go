package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testingTAdapter struct {
    t *testing.T
}

/*
 * Testify framework doesn't natively implement assertion for strings, so we
 * need to create a custom testing structure that implementes the TestingT
 * interface
 */
func (ta *testingTAdapter) Errorf(format string, args ...interface{}) {
    ta.t.Errorf(format, args...)
}

func TestRootCmd(t *testing.T) {
    // Create a buffer to capture the command's output
    var outputBuffer bytes.Buffer

    // Set the output of the RootCmd to our buffer
    RootCmd.SetOut(&outputBuffer)

    // Create a TestingT adapter
    tt := &testingTAdapter{t}

    // Run the RootCmd
    RootCmd.SetArgs([]string{"--help"}) // You can use any valid command line arguments
    if err := RootCmd.Execute(); err != nil {
        tt.Errorf("RootCmd execution error: %v", err)
    }

    // Check if the expected message is present in the output
    output := outputBuffer.String()
    expectedMessage := "Welcome to DLP's CLI!"
    assert.Contains(tt, output, expectedMessage)
}
