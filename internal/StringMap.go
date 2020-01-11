package internal

import (
    "reflect"

    "github.com/AMan4Technology/Serialize/codec"
)

func init() {
    _ = Register(reflect.TypeOf(StringMap{}), StringMapSerializer{}, false)
}

type StringMap map[string]string

type StringMapSerializer struct{}

func (s StringMapSerializer) Serialize(value interface{}, tag, codecID string, varMap StringMap) (string, error) {
    var (
        stringMap   = value.(StringMap)
        stringSlice = make(StringSlice, 0, len(stringMap))
    )
    for k, v := range stringMap {
        stringSlice = append(stringSlice, codec.Encode(codecID, k, "", v))
    }
    return SerializerWithID[StringSliceID].Serialize(stringSlice, tag, codecID, varMap)
}

func (s StringMapSerializer) Deserialize(data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (interface{}, error) {
    val, err := SerializerWithID[StringSliceID].Deserialize(data, tag, codecID, varMap, ptrMap)
    if err != nil {
        return nil, err
    }
    stringSlice := val.(StringSlice)
    stringMap := make(StringMap, len(stringSlice))
    for _, kAndV := range stringSlice {
        k, _, v := codec.Decode(codecID, kAndV)
        stringMap[k] = v
    }
    return stringMap, nil
}
