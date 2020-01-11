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

func structSerialize(value reflect.Value, tag, codecID string, varMap StringMap) string {
    var (
        length = value.NumField()
        fields = make(StringSlice, 0, length)
    )
    for i := 0; i < length; i++ {
        field := value.Field(i)
        if !field.CanInterface() {
            continue
        }
        fieldName := value.Type().Field(i).Tag.Get(tag)
        if fieldName == "" {
            fieldName = value.Type().Field(i).Name
        }
        tpID, data, err := serialize(field, tag, codecID, varMap)
        if err != nil {
            continue
        }
        fields = append(fields, codec.Encode(codecID, fieldName, tpID, data))
    }
    data, err := SerializerWithID[StringSliceID].Serialize(fields, tag, codecID, varMap)
    if err != nil {
        fmt.Println(err)
    }
    return data
}

func structDeserialize(tp reflect.Type, data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (result reflect.Value, err error) {
    result = reflect.New(tp).Elem()
    fields, err := SerializerWithID[StringSliceID].Deserialize(data, tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    fieldWithName := StructFieldWithName(result, tag)
    for _, data := range fields.(StringSlice) {
        name, typeID, value := codec.Decode(codecID, data)
        val, err := deserialize(typeID, value, tag, codecID, varMap, ptrMap)
        if err != nil {
            continue
        }
        field := fieldWithName[name]
        if !field.CanSet() {
            continue
        }
        field.Set(val)
    }
    return
}

func structDeserializeWithMap(data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (result reflect.Value, err error) {
    fields, err := SerializerWithID[StringSliceID].Deserialize(data, tag, codecID, varMap, ptrMap)
    if err != nil {
        return
    }
    fieldSlice := fields.(StringSlice)
    structValue := make(map[string]interface{}, len(fieldSlice))
    for _, data := range fieldSlice {
        name, typeID, value := codec.Decode(codecID, data)
        val, err := deserialize(typeID, value, tag, codecID, varMap, ptrMap)
        if err != nil {
            continue
        }
        structValue[name] = val.Interface()
    }
    return reflect.ValueOf(structValue), nil
}
