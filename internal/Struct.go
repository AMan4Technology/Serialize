package internal

import (
    "fmt"
    "reflect"

    "github.com/AMan4Technology/Serialize/codec"
)

func StructFieldWithName(val reflect.Value, tag string) (fieldWithName map[string]reflect.Value) {
    length := val.NumField()
    fieldWithName = make(map[string]reflect.Value, length)
    for i := 0; i < length; i++ {
        fieldName := val.Type().Field(i).Tag.Get(tag)
        if fieldName == "" {
            fieldName = val.Type().Field(i).Name
        }
        fieldWithName[fieldName] = val.Field(i)
    }
    return
}

func structSerialize(value interface{}, tag string) string {
    var (
        val    = reflect.ValueOf(value)
        length = val.NumField()
        fields = make(StringSlice, 0, length)
    )
    for i := 0; i < length; i++ {
        field := val.Field(i)
        if !field.CanInterface() {
            continue
        }
        fieldName := val.Type().Field(i).Tag.Get(tag)
        if fieldName == "" {
            fieldName = val.Type().Field(i).Name
        }
        data, err := Serialize(field.Interface(), codec.String, fieldName, tag)
        if err != nil {
            continue
        }
        fields = append(fields, data)
    }
    data, err := Serialize(fields, codec.String, "fields", tag)
    if err != nil {
        fmt.Println(err)
    }
    return data
}

func structDeserialize(tp reflect.Type, data, tag string, isPtr bool) (interface{}, error) {
    var (
        result = reflect.New(tp)
        val    = result.Elem()
    )
    if !isPtr {
        result = val
    }
    fields, _, err := Deserialize(data, codec.String, tag)
    if err != nil {
        return result.Interface(), err
    }
    fieldWithName := StructFieldWithName(val, tag)
    for _, data := range fields.(StringSlice) {
        value, name, err := Deserialize(data, codec.String, tag)
        if err != nil {
            continue
        }
        field := fieldWithName[name]
        if !field.CanSet() {
            continue
        }
        field.Set(reflect.ValueOf(value))
    }
    return result.Interface(), nil
}

func structDeserializeWithMap(data, tag string) (structValue map[string]interface{}, err error) {
    fields, _, err := Deserialize(data, codec.String, tag)
    if err != nil {
        return nil, err
    }
    fieldSlice := fields.(StringSlice)
    structValue = make(map[string]interface{}, len(fieldSlice))
    for _, data := range fieldSlice {
        value, name, err := Deserialize(data, codec.String, tag)
        if err != nil {
            continue
        }
        structValue[name] = value
    }
    return
}
