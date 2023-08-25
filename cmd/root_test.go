package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func TestRootCmd(t *testing.T) {
	args := []string{}

	_, err := execute(t, rootCmd, args...)

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}
}

func TestRootCmdVersion(t *testing.T) {
	args := []string{"version"}

	_, err := execute(t, rootCmd, args...)

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}
}

func TestRootCmdDebugFlag(t *testing.T) {
	args := []string{"--debug"}

	_, err := execute(t, rootCmd, args...)

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}
}

func TestRootCmdUnknownFlag(t *testing.T) {
	args := []string{"--unknown"}

	_, err := execute(t, rootCmd, args...)

	if err == nil {
		t.Errorf("Error was expected but received: %s", err)
	}
}
