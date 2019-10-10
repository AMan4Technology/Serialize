package internal

import (
    "fmt"
    "reflect"

    "Serialize/codec"
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
            fmt.Println(err)
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

func structDeserialize(tp reflect.Type, data, tag string) (interface{}, error) {
    val := reflect.New(tp).Elem()
    fields, _, err := Deserialize(data, codec.String, tag)
    if err != nil {
        return val.Interface(), err
    }
    fieldWithName := StructFieldWithName(val, tag)
    for _, data := range fields.(StringSlice) {
        i, name, err := Deserialize(data, codec.String, tag)
        if err != nil {
            fmt.Println(err)
            continue
        }
        field := fieldWithName[name]
        if field.CanSet() {
            field.Set(reflect.ValueOf(i))
        }
    }
    return val.Interface(), nil
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
            fmt.Println(err)
            continue
        }
        structValue[name] = value
    }
    return
}
