package internal

import (
    "errors"
    "fmt"
    "reflect"

    "github.com/AMan4Technology/Serialize/codec"
)

func mapSerialize(value reflect.Value, tag, codecID string, varMap StringMap) string {
    var (
        length = value.Len()
        keys   = make(StringSlice, length)
        values = make(StringSlice, length)
        iter   = value.MapRange()
    )
    for i := 0; iter.Next(); i++ {
        k := iter.Key()
        v := iter.Value()
        tpID, data, err := serialize(k, tag, codecID, varMap)
        if err != nil {
            continue
        }
        key := codec.Encode(codecID, "", tpID, data)
        tpID, data, err = serialize(v, tag, codecID, varMap)
        if err != nil {
            continue
        }
        value := codec.Encode(codecID, "", tpID, data)
        keys[i] = key
        values[i] = value
    }
    mapDataSlice := make(StringSlice, 2)

    data, err := SerializerWithID[StringSliceID].Serialize(keys, tag, codecID, varMap)
    if err != nil {
        fmt.Println(err)
    }
    mapDataSlice[0] = data

    data, err = SerializerWithID[StringSliceID].Serialize(values, tag, codecID, varMap)
    if err != nil {
        fmt.Println(err)
    }
    mapDataSlice[1] = data

    data, err = SerializerWithID[StringSliceID].Serialize(mapDataSlice, tag, codecID, varMap)
    if err != nil {
        fmt.Println(err)
    }
    return data
}

func mapDeserialize(tp reflect.Type, data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (result reflect.Value, err error) {
    mapData, err := SerializerWithID[StringSliceID].Deserialize(data, tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    keysAndValues := mapData.(StringSlice)
    if len(keysAndValues) != 2 {
        return result, errors.New("valid data keys&values")
    }
    keys, err := SerializerWithID[StringSliceID].Deserialize(keysAndValues[0], tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    values, err := SerializerWithID[StringSliceID].Deserialize(keysAndValues[1], tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    var (
        keySlice   = keys.(StringSlice)
        valueSlice = values.(StringSlice)
        mapValue   = reflect.MakeMap(tp)
    )
    for i, key := range keySlice {
        _, typeID, value := codec.Decode(codecID, key)
        k, err := deserialize(typeID, value, tag, codecID, varMap, ptrMap)
        if err != nil {
            if err.Error() != Nil {
                continue
            }
            k = reflect.New(tp.Key()).Elem()
        }
        _, typeID, value = codec.Decode(codecID, valueSlice[i])
        v, err := deserialize(typeID, value, tag, codecID, varMap, ptrMap)
        if err != nil {
            continue
        }
        mapValue.SetMapIndex(k, v)
    }
    return mapValue, nil
}

func mapDeserializeWith(data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (result reflect.Value, err error) {
    mapData, err := SerializerWithID[StringSliceID].Deserialize(data, tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    keysAndValues := mapData.(StringSlice)
    if len(keysAndValues) != 2 {
        return result, errors.New("valid data keys&values")
    }
    keys, err := SerializerWithID[StringSliceID].Deserialize(keysAndValues[0], tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    values, err := SerializerWithID[StringSliceID].Deserialize(keysAndValues[1], tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    var (
        keySlice   = keys.(StringSlice)
        valueSlice = values.(StringSlice)
    )
    mapValue := make(map[interface{}]interface{}, len(keySlice))
    for i, key := range keySlice {
        _, typeID, value := codec.Decode(codecID, key)
        k, err := deserialize(typeID, value, tag, codecID, varMap, ptrMap)
        if err != nil && err.Error() != Nil {
            continue
        }
        _, typeID, value = codec.Decode(codecID, valueSlice[i])
        v, err := deserialize(typeID, value, tag, codecID, varMap, ptrMap)
        if err != nil {
            continue
        }
        mapValue[k.Interface()] = v.Interface()
    }
    return reflect.ValueOf(mapValue), nil
}
