package internal

import (
    "reflect"
    "strings"

    "Serialize/codec"
)

const Ptr = "*"

func Serialize(value interface{}, codecID, name, tag string) (data string, err error) {
    tpID, data, err := serialize(value, tag)
    return codec.Encode(codecID, tpID, name, data), err
}

func Deserialize(data string, codecID, tag string) (value interface{}, name string, err error) {
    tpID, name, data := codec.Decode(codecID, data)
    value, err = deserialize(tpID, data, tag)
    return
}

type Serializer interface {
    Serialize(value interface{}, tag string) (string, error)
    Deserialize(data string, tag string) (interface{}, error)
}

func serialize(value interface{}, tag string) (tpID, data string, err error) {
    val := reflect.ValueOf(value)
    if val.Kind() == reflect.Ptr {
        tpID, data, err = serialize(val.Elem().Interface(), tag)
        return Ptr + tpID, data, err
    }
    tpID = IDOf(val.Type())
    if serializer := SerializerWithID[tpID]; serializer == nil {
        tpID = val.Kind().String()
    } else if serializer.Serializer != nil {
        data, err = serializer.Serialize(value, tag)
        return
    }
    data, err = defaultSerialize(value, tag)
    return
}

func deserialize(tpID, data, tag string) (interface{}, error) {
    if strings.HasPrefix(tpID, Ptr) {
        return deserialize(tpID[1:], data, tag)
    }
    if serializer := SerializerWithID[tpID]; serializer != nil {
        if serializer.Serializer != nil {
            return serializer.Deserialize(data, tag)
        }
        return defaultDeserialize(serializer.TP, data, tag)
    }
    return defaultDeserializeWith(tpID, data, tag)
}

type serializer struct {
    ID string
    TP reflect.Type
    Serializer
}
