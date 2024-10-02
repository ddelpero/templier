package main

import (
	_ "embed"
	"os"
	"testing"

	"github.com/ddelpero/templier/internal/config"
	"github.com/romshark/yamagiconf"
	"github.com/stretchr/testify/require"
)

//go:embed example-config.yml
var exampleConfig string

func TestExampleConfig(t *testing.T) {
	var c config.Config

	err := yamagiconf.Load(exampleConfig, &c)
	require.NoError(t, err)
	c.App.SetEnv()
	require.Equal(t, "example-value", os.Getenv("EXAMPLE_ENV_VAR"))
	require.Equal(t, "https://example.com", os.Getenv("EXAMPLE_ENV_URL"))
}
