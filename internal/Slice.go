package internal

import (
    "fmt"
    "reflect"
    "strconv"

    "Serialize/codec"
)

func sliceSerialize(value interface{}, tag string) string {
    var (
        val    = reflect.ValueOf(value)
        length = val.Len()
        slice  = make(StringSlice, length)
    )
    for i := 0; i < length; i++ {
        elemData, err := Serialize(val.Index(i).Interface(), codec.String, strconv.Itoa(i), tag)
        if err != nil {
            continue
        }
        slice[i] = elemData
    }
    sliceData, _ := Serialize(slice, codec.String, "sliceData", tag)
    return sliceData
}

func sliceDeserialize(tp reflect.Type, data, tag string) (interface{}, error) {
    val := reflect.New(tp).Elem()
    sliceData, _, err := Deserialize(data, codec.String, tag)
    if err != nil {
        return val.Interface(), err
    }
    slice := sliceData.(StringSlice)
    for _, elem := range slice {
        value, _, err := Deserialize(elem, codec.String, tag)
        if err != nil {
            fmt.Printf("parse %s failed, error: %e", elem, err)
        }
        reflect.Append(val, reflect.ValueOf(value))
    }
    return val.Interface(), nil
}

func sliceDeserializeWith(data, tag string) (sliceValue []interface{}, err error) {
    sliceData, _, err := Deserialize(data, codec.String, tag)
    if err != nil {
        return nil, err
    }
    var (
        slice  = sliceData.(StringSlice)
        length = len(slice)
    )
    sliceValue = make([]interface{}, length)
    for i, elem := range slice {
        value, _, err := Deserialize(elem, codec.String, tag)
        if err != nil {
            fmt.Printf("parse %s failed, error: %e", elem, err)
        }
        sliceValue[i] = value
    }
    return
}