package goconf

import (
    "io/ioutil"
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

var bytes = []byte(`
a: foo
camelcase: true
manualoverride1: "foobar"
splitword1: "hello world"
id: 123456
`)

func TestEnv(t *testing.T) {
    os.Setenv("GO_A", "foo")
    os.Setenv("GO_CamelCase", "true")
    os.Setenv("GO_ManualOverride1", "foobar")
    os.Setenv("GO_SplitWord1", "hello world")
    os.Setenv("GO_ID", "123456")
    os.Setenv("GO_DefaultValue", "default")

    sample := &SampleA{}
    err := Parse(sample, WithEnv("go"))
    require.NoError(t, err)
    require.EqualValues(t, expSample, sample)
}

func TestYaml(t *testing.T) {
    sample := &SampleA{DefaultValue: "default"}
    err := Parse(sample, WithYamlFromBytes(bytes))
    require.NoError(t, err)
    require.EqualValues(t, expSample, sample)
}

func TestYamlFromFile(t *testing.T) {
    err := ioutil.WriteFile("config1.yml", bytes, 0777)
    require.NoError(t, err)

    sample := &SampleA{DefaultValue: "default"}
    err = Parse(sample, WithYaml("config1.yml"))
    require.NoError(t, err)
    require.EqualValues(t, expSample, sample)
}

func TestCombile(t *testing.T) {
    os.Setenv("GOO_A", "baz")

    cbytes := []byte(`
a: foo
b:
  e:
  - 3
  - 3
  - 3
manualoverride1: "foobar"
`)

    cfile := []byte(`
a: "bar"
b:
  e:
  - 4
  - 4
  - 4
`)

    cexp := &SampleA{
        A:               "baz",
        ManualOverride1: "foobar",
        DefaultValue:    "123",
    }

    sample := &SampleA{
        A:               "000",
        ManualOverride1: "111",
        DefaultValue:    "default",
    }

    err := ioutil.WriteFile("config2.yml", cfile, 0777)
    require.NoError(t, err)

    err = Parse(sample,
        WithYamlFromBytes(cbytes),
        WithYaml("config2.yml"),
        WithEnv("goo"),
    )
    require.NoError(t, err)
    require.EqualValues(t, cexp, sample)
}
