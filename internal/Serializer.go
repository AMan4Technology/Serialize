package internal

import (
    "errors"
    "reflect"
    "strings"

    "github.com/AMan4Technology/Serialize/codec"
)

const Ptr = "*"

func Serialize(value interface{}, codecID, name, tag string) (data string, err error) {
    tpID, data, err := serialize(value, tag)
    if err != nil {
        return "", err
    }
    return codec.Encode(codecID, tpID, name, data), err
}

func Deserialize(data string, codecID, tag string) (value interface{}, name string, err error) {
    tpID, name, data := codec.Decode(codecID, data)
    if tpID == "" {
        return nil, "", errors.New(Nil)
    }
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
        if val.Elem().IsValid() {
            tpID, data, err = serialize(val.Elem().Interface(), tag)
            return Ptr + tpID, data, err
        }
        return "", "", errors.New(Nil)
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
    var (
        index = strings.LastIndex(tpID, Ptr)
        isPtr = index != -1
    )
    if isPtr {
        tpID = tpID[index+1:]
    }
    if serializer := SerializerWithID[tpID]; serializer != nil {
        val := reflect.New(serializer.TP)
        if serializer.Serializer != nil {
            value, err := serializer.Deserialize(data, tag)
            if !isPtr {
                return value, err
            }
            if err == nil {
                val.Elem().Set(reflect.ValueOf(value))
            }
            return val.Interface(), err
        }
        return defaultDeserialize(serializer.TP, data, tag, isPtr)
    }
    return defaultDeserializeWith(tpID, data, tag, isPtr)
}

type serializer struct {
    ID string
    TP reflect.Type
    Serializer
}
