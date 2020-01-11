package internal

import (
    "fmt"
    "reflect"
    "strconv"

    "github.com/AMan4Technology/Serialize/codec"
)

func sliceSerialize(value reflect.Value, tag, codecID string, varMap StringMap) string {
    var (
        length = value.Len()
        slice  = make(StringSlice, length)
    )
    for i := 0; i < length; i++ {
        tpID, data, err := serialize(value.Index(i), tag, codecID, varMap)
        if err != nil {
            continue
        }
        slice[i] = codec.Encode(codecID, strconv.Itoa(i), tpID, data)
    }
    data, _ := SerializerWithID[StringSliceID].Serialize(slice, tag, codecID, varMap)
    return data
}

func sliceDeserialize(tp reflect.Type, data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (result reflect.Value, err error) {
    sliceValue, err := SerializerWithID[StringSliceID].Deserialize(data, tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    slice := sliceValue.(StringSlice)
    sliceVal := reflect.MakeSlice(tp, len(slice), len(slice))
    for i, elem := range slice {
        _, typeID, value := codec.Decode(codecID, elem)
        val, err := deserialize(typeID, value, tag, codecID, varMap, ptrMap)
        if err != nil {
            continue
        }
        sliceVal.Index(i).Set(val)
    }
    return sliceVal, nil
}

func sliceDeserializeWith(data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (result reflect.Value, err error) {
    sliceData, err := SerializerWithID[StringSliceID].Deserialize(data, tag, codecID, varMap, ptrMap)
    if err != nil {
        return reflect.Value{}, err
    }
    var (
        slice  = sliceData.(StringSlice)
        length = len(slice)
    )
    sliceValue := make([]interface{}, length)
    for i, elem := range slice {
        _, typeID, value := codec.Decode(codecID, elem)
        val, err := deserialize(typeID, value, tag, codecID, varMap, ptrMap)
        if err != nil {
            fmt.Printf("parse %s failed, error: %e", elem, err)
        }
        sliceValue[i] = val.Interface()
    }
    return reflect.ValueOf(sliceValue), nil
}
