package internal

import (
    "errors"
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

func baseSerialize(value reflect.Value) string {
    return fmt.Sprint(value.Interface())
}

func baseDeserialize(tp reflect.Type, data string) (result reflect.Value, err error) {
    result = reflect.New(tp).Elem()
    value, err := baseDeserializeWith(tp.Kind().String(), data)
    if err == nil {
        result.Set(reflect.ValueOf(value).Convert(tp))
    }
    return
}

func baseDeserializeWith(kind, data string) (interface{}, error) {
    switch kind {
    case reflect.String.String():
        return data, nil
    case reflect.Bool.String():
        b, err := strconv.ParseBool(data)
        return b, err
    case reflect.Int.String():
        i64, err := strconv.ParseInt(data, 0, 0)
        return int(i64), err
    case reflect.Int8.String():
        i64, err := strconv.ParseInt(data, 0, 8)
        return int8(i64), err
    case reflect.Int16.String():
        i64, err := strconv.ParseInt(data, 0, 16)
        return int16(i64), err
    case reflect.Int32.String():
        i64, err := strconv.ParseInt(data, 0, 32)
        return int32(i64), err
    case reflect.Int64.String():
        i64, err := strconv.ParseInt(data, 0, 64)
        return i64, err
    case reflect.Float32.String():
        f64, err := strconv.ParseFloat(data, 32)
        return float32(f64), err
    case reflect.Float64.String():
        return strconv.ParseFloat(data, 64)
    case reflect.Uint.String():
        u64, err := strconv.ParseUint(data, 0, 0)
        return uint(u64), err
    case reflect.Int8.String():
        u64, err := strconv.ParseUint(data, 0, 8)
        return uint8(u64), err
    case reflect.Int16.String():
        u64, err := strconv.ParseUint(data, 0, 16)
        return uint16(u64), err
    case reflect.Int32.String():
        u64, err := strconv.ParseUint(data, 0, 32)
        return uint(u64), err
    case reflect.Int64.String():
        return strconv.ParseUint(data, 0, 64)
    case reflect.Complex64.String():
        var (
            split = strings.LastIndexByte(data, '+') + strings.LastIndexByte(data, '-')
            a, _  = baseDeserializeWith("float32", data[:split])
            b, _  = baseDeserializeWith("float32", data[split:len(data)-1])
        )
        return complex(a.(float32), b.(float32)), nil
    case reflect.Complex128.String():
        var (
            split = strings.LastIndexByte(data, '+') + strings.LastIndexByte(data, '-')
            a, _  = baseDeserializeWith("float64", data[:split])
            b, _  = baseDeserializeWith("float64", data[split:len(data)-1])
        )
        return complex(a.(float64), b.(float64)), nil
    default:
        return nil, errors.New("unsupported kind:" + kind)
    }
}
