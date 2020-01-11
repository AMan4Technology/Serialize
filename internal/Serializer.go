package internal

import (
    "errors"
    "reflect"
    "strings"

    "github.com/AMan4Technology/Serialize/codec"
)

const (
    Ptr    = "*"
    Target = "target"
)

var VarMapID = IDOf(reflect.TypeOf(StringMap{}))
var StringSliceID = IDOf(reflect.TypeOf(StringSlice{}))

func Serialize(value interface{}, name, tag, codecID string) (data string, err error) {
    varMap := make(StringMap, 1)
    tpID, data, err := serialize(reflect.ValueOf(value), tag, codecID, varMap)
    if err != nil {
        return "", err
    }
    varMap[Target] = codec.Encode(codecID, name, tpID, data)
    return SerializerWithID[VarMapID].Serialize(varMap, tag, codecID, nil)
}

func Deserialize(data, tag, codecID string) (value interface{}, name string, err error) {
    stringMap, err := SerializerWithID[VarMapID].Deserialize(data, tag, codecID, nil, nil)
    if err != nil {
        return nil, "", err
    }
    varMap := stringMap.(StringMap)
    name, tpID, data := codec.Decode(codecID, varMap[Target])
    if tpID == "" {
        return nil, "", errors.New(Nil)
    }
    delete(varMap, Target)
    val, err := deserialize(tpID, data, tag, codecID, varMap, make(map[string]reflect.Value, len(varMap)))
    if err == nil && val.CanInterface() {
        value = val.Interface()
    }
    return
}

type Serializer interface {
    Serialize(value interface{}, tag, codecID string, varMap StringMap) (string, error)
    Deserialize(data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (interface{}, error)
}

func serialize(value reflect.Value, tag, codecID string, varMap StringMap) (tpID, data string, err error) {
    tpID = IDOf(value.Type())
    switch value.Kind() {
    case reflect.Ptr:
        ptr := StringFrom(value)
        if _, ok := varMap[ptr]; !ok && value.Elem().IsValid() {
            tpID, data, err = serialize(value.Elem(), tag, codecID, varMap)
            varMap[ptr] = codec.Encode(codec.String, "", tpID, data)
        }
        return Ptr + tpID, ptr, err
    case reflect.Map, reflect.Slice:
        ptr := StringFrom(value)
        if _, ok := varMap[ptr]; ok {
            return tpID, ptr, err
        }
    }
    if serializer := SerializerWithID[tpID]; serializer == nil {
        tpID = value.Kind().String()
    } else if serializer.Serializer != nil {
        data, err = serializer.Serialize(value.Interface(), tag, codecID, varMap)
        goto result
    }
    data, err = defaultSerialize(value, tag, codecID, varMap)
result:
    switch value.Kind() {
    case reflect.Map, reflect.Slice:
        ptr := StringFrom(value)
        varMap[ptr], data = data, ptr
    }
    return
}

func deserialize(tpID, data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (reflect.Value, error) {
    if strings.HasPrefix(tpID, Ptr) {
        if ptr, ok := ptrMap[data]; ok {
            return ptr, nil
        }
        _, tpID, value := codec.Decode(codecID, varMap[data])
        val, err := deserialize(tpID, value, tag, codecID, varMap, ptrMap)
        if err == nil && val.CanAddr() {
            val = val.Addr()
            ptrMap[data] = val
        }
        return val, err
    }
    if serializer := SerializerWithID[tpID]; serializer != nil {
        if serializer.Serializer != nil {
            var ptr string
            switch serializer.TP.Kind() {
            case reflect.Map, reflect.Slice:
                ptr, data = data, varMap[data]
                if ptr, ok := ptrMap[ptr]; ok {
                    return ptr, nil
                }
            }
            val := reflect.New(serializer.TP).Elem()
            value, err := serializer.Deserialize(data, tag, codecID, varMap, ptrMap)
            if err == nil {
                val.Set(reflect.ValueOf(value))
            }
            if ptr != "" {
                ptrMap[ptr] = val
            }
            return val, err
        }
        return defaultDeserialize(serializer.TP, data, tag, codecID, varMap, ptrMap)
    }
    return defaultDeserializeWith(tpID, data, tag, codecID, varMap, ptrMap)
}

type serializer struct {
    ID string
    TP reflect.Type
    Serializer
}
