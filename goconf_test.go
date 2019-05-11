package goconf

import (
    "os"
    "testing"

    "github.com/stretchr/testify/require"
)

type SampleA struct {
    A               string
    CamelCase       bool
    ManualOverride1 string
    SplitWord1      string
    ID              string
    DefaultValue    string `default:"123"`
}

var expSample = &SampleA{
    A:               "foo",
    CamelCase:       true,
    ManualOverride1: "foobar",
    SplitWord1:      "hello world",
    ID:              "123456",
    DefaultValue:    "default",
}

func TestEnv(t *testing.T) {
    _ = os.Setenv("GO_A", "foo")
    _ = os.Setenv("GO_CamelCase", "true")
    _ = os.Setenv("GO_ManualOverride1", "foobar")
    _ = os.Setenv("GO_SplitWord1", "hello world")
    _ = os.Setenv("GO_ID", "123456")
    _ = os.Setenv("GO_DefaultValue", "default")

    sample := &SampleA{}
    Parse("go", sample)
    require.EqualValues(t, expSample, sample)
}
