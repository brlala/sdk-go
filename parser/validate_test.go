package parser

import (
	"embed"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/workflows/*
var contentFiles embed.FS

func Test_parseSchema(t *testing.T) {
	path := "testdata/workflows"
	sch, err := jsonschema.Compile("https://serverlessworkflow.io/schemas/0.8/workflow.json")
	assert.NoError(t, err)

	err = validateSchema(contentFiles, path, sch)
	assert.NoError(t, err)
}
