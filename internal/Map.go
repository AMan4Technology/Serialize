package internal

import (
    "errors"
    "fmt"
    "reflect"

    "Serialize/codec"
)

func mapSerialize(value interface{}, tag string) string {
    var (
        val    = reflect.ValueOf(value)
        length = val.Len()
        keys   = make(StringSlice, length)
        values = make(StringSlice, length)
        iter   = val.MapRange()
    )
    for i := 0; iter.Next(); i++ {
        k := iter.Key()
        v := iter.Value()
        key, err := Serialize(k.Interface(), codec.String, "", tag)
        if err != nil {
            fmt.Println(err)
            continue
        }
        value, err := Serialize(v.Interface(), codec.String, "", tag)
        if err != nil {
            fmt.Println(err)
            continue
        }
        keys[i] = key
        values[i] = value
    }
    mapDataSlice := make(StringSlice, 2)
    keysData, err := Serialize(keys, codec.String, "keys", tag)
    if err != nil {
        fmt.Println(err)
    }
    mapDataSlice[0] = keysData
    valuesData, err := Serialize(values, codec.String, "values", tag)
    if err != nil {
        fmt.Println(err)
    }
    mapDataSlice[1] = valuesData
    mapData, err := Serialize(mapDataSlice, codec.String, "mapData", tag)
    if err != nil {
        fmt.Println(err)
    }
    return mapData
}

func mapDeserialize(tp reflect.Type, data, tag string) (interface{}, error) {
    val := reflect.New(tp).Elem()
    mapData, _, err := Deserialize(data, codec.String, tag)
    if err != nil {
        return val.Interface(), err
    }
    keysAndValues := mapData.(StringSlice)
    if len(keysAndValues) != 2 {
        return val.Interface(), errors.New("valid data keys&values")
    }
    keys, _, err := Deserialize(keysAndValues[0], codec.String, tag)
    if err != nil {
        return val.Interface(), err
    }
    values, _, err := Deserialize(keysAndValues[1], codec.String, tag)
    if err != nil {
        return val.Interface(), err
    }
    var (
        keySlice   = keys.(StringSlice)
        valueSlice = values.(StringSlice)
    )
    for i, key := range keySlice {
        k, _, err := Deserialize(key, codec.String, tag)
        if err != nil {
            continue
        }
        v, _, err := Deserialize(valueSlice[i], codec.String, tag)
        if err != nil {
            continue
        }
        val.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
    }
    return val.Interface(), nil
}

func mapDeserializeWith(data, tag string) (mapValue map[interface{}]interface{}, err error) {
    mapData, _, err := Deserialize(data, codec.String, tag)
    if err != nil {
        return nil, err
    }
    keysAndValues := mapData.(StringSlice)
    if len(keysAndValues) != 2 {
        return nil, errors.New("valid data keys&values")
    }
    keys, _, err := Deserialize(keysAndValues[0], codec.String, tag)
    if err != nil {
        return nil, err
    }
    values, _, err := Deserialize(keysAndValues[1], codec.String, tag)
    if err != nil {
        return nil, err
    }
    var (
        keySlice   = keys.(StringSlice)
        valueSlice = values.(StringSlice)
    )
    mapValue = make(map[interface{}]interface{}, len(keySlice))
    for i, key := range keySlice {
        k, _, err := Deserialize(key, codec.String, tag)
        if err != nil {
            continue
        }
        v, _, err := Deserialize(valueSlice[i], codec.String, tag)
        if err != nil {
            continue
        }
        mapValue[k] = v
    }
    return
}
