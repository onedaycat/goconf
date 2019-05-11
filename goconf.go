package goconf

import (
    "io/ioutil"
    "os"
    "reflect"
    "strconv"
    "strings"

    "gopkg.in/yaml.v2"
)

func Parse(v interface{}, options ...Option) error {
    var err error
    o := &opts{}
    for _, option := range options {
        option(o)
    }

    if o.yaml {
        if err = parseYaml(v, o); err != nil {
            return err
        }
    }

    if o.env {
        rv := reflect.ValueOf(v)
        rt := reflect.TypeOf(v)

        for i := 0; i < rt.Elem().NumField(); i++ {
            rti := rt.Elem().Field(i)
            rvi := rv.Elem().Field(i)
            data := os.Getenv(o.envPrefix + "_" + rti.Name)
            if data == "" {
                data = rti.Tag.Get("default")
            }

            if rvi.Type().Kind().String() == "slice" {
                setSliceValue(&rvi, rvi.Type().String(), data)
            } else {
                setValue(&rvi, rvi.Type().Name(), data)
            }
        }
    }

    return nil
}

func parseYaml(v interface{}, o *opts) error {
    var err error
    if o.yamlBytes != nil {
        err = yaml.Unmarshal(o.yamlBytes, v)
    }

    for _, path := range o.yamlPaths {
        yamlBytes, err := ioutil.ReadFile(path)
        if err != nil && !os.IsNotExist(err) {
            return err
        }
        return yaml.Unmarshal(yamlBytes, v)
    }

    return err
}

func setValue(rv *reflect.Value, valueType string, value string) {
    if value == "" {
        return
    }

    switch valueType {
    case "string":
        rv.SetString(value)
    case "bool":
        if value == "1" || value == "true" {
            rv.SetBool(true)
        }
    case "int64":
        i, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            panic(err)
        }
        rv.SetInt(i)
    case "int":
        i, err := strconv.ParseInt(value, 10, 32)
        if err != nil {
            panic(err)
        }
        rv.SetInt(i)
    case "float64":
        i, err := strconv.ParseFloat(value, 64)
        if err != nil {
            panic(err)
        }
        rv.SetFloat(i)
    }
}

func cretaeSlice(valueType string, n int) reflect.Value {
    switch valueType {
    case "[]string":
        return reflect.MakeSlice(reflect.TypeOf([]string{}), n, n)
    case "[]bool":
        return reflect.MakeSlice(reflect.TypeOf([]bool{}), n, n)
    case "[]int64":
        return reflect.MakeSlice(reflect.TypeOf([]int64{}), n, n)
    case "[]int":
        return reflect.MakeSlice(reflect.TypeOf([]int{}), n, n)
    case "[]float64":
        return reflect.MakeSlice(reflect.TypeOf([]float64{}), n, n)
    }

    return reflect.Value{}
}

func setSliceValue(rv *reflect.Value, valueType string, value string) {
    if value == "" {
        return
    }
    vals := strings.Split(value, ",")
    slice := cretaeSlice(valueType, len(vals))
    for i, val := range vals {
        switch valueType {
        case "[]string":
            slice.Index(i).SetString(val)
        case "[]bool":
            if val == "1" || val == "true" {
                slice.Index(i).SetBool(true)
            }
        case "[]int64":
            ii, err := strconv.ParseInt(val, 10, 64)
            if err != nil {
                panic(err)
            }
            slice.Index(i).SetInt(ii)
        case "[]int":
            ii, err := strconv.ParseInt(val, 10, 32)
            if err != nil {
                panic(err)
            }
            slice.Index(i).SetInt(ii)
        case "[]float64":
            ii, err := strconv.ParseFloat(val, 64)
            if err != nil {
                panic(err)
            }
            slice.Index(i).SetFloat(ii)
        }
    }
    rv.Set(slice)
}
