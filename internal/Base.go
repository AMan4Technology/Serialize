package internal

import (
    "errors"
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

const (
    Int     = "int"
    Uint    = "uint"
    Float   = "float"
    Complex = "complex"
)

func baseSerialize(value interface{}) string {
    return fmt.Sprint(value)
}

func baseDeserialize(tp reflect.Type, data string) (interface{}, error) {
    value, _ := baseDeserializeWith(tp.Kind().String(), data)
    return reflect.ValueOf(value).Convert(tp).Interface(), nil
}

func baseDeserializeWith(kind, data string) (interface{}, error) {
    switch kind {
    case reflect.String.String():
        return data, nil
    case reflect.Bool.String():
        return strconv.ParseBool(data)
    case reflect.Int.String():
        return strconv.Atoi(data)
    case reflect.Int8.String(), reflect.Int16.String(), reflect.Int32.String(), reflect.Int64.String():
        bitSize, _ := strconv.Atoi(kind[len(Int):])
        return strconv.ParseInt(data, 0, bitSize)
    case reflect.Float32.String(), reflect.Float64.String():
        bitSize, _ := strconv.Atoi(kind[len(Float):])
        return strconv.ParseFloat(data, bitSize)
    case reflect.Uint.String(), reflect.Int8.String(), reflect.Int16.String(), reflect.Int32.String(), reflect.Int64.String():
        bitSize, _ := strconv.Atoi(kind[len(Uint):])
        return strconv.ParseUint(data, 0, bitSize)
    case reflect.Complex64.String():
        var (
            bitSize, _ = strconv.Atoi(kind[len(Complex):])
            tp         = Float + strconv.Itoa(bitSize/2)
            split      = strings.LastIndexByte(data, '+') + strings.LastIndexByte(data, '-')
            a, _       = baseDeserializeWith(tp, data[:split])
            b, _       = baseDeserializeWith(tp, data[split:len(data)-1])
        )
        return complex(a.(float32), b.(float32)), nil
    case reflect.Complex128.String():
        var (
            bitSize, _ = strconv.Atoi(kind[len(Complex):])
            tp         = Float + strconv.Itoa(bitSize/2)
            split      = strings.LastIndexByte(data, '+') + strings.LastIndexByte(data, '-')
            a, _       = baseDeserializeWith(tp, data[:split])
            b, _       = baseDeserializeWith(tp, data[split:len(data)-1])
        )
        return complex(a.(float64), b.(float64)), nil
    default:
        return nil, errors.New("unsupported kind:" + kind)
    }
}
