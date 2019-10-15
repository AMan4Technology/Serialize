package internal

import (
    "errors"
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

func baseSerialize(value interface{}) string {
    return fmt.Sprint(value)
}

func baseDeserialize(tp reflect.Type, data string, isPtr bool) (value interface{}, err error) {
    value, err = baseDeserializeWith(tp.Kind().String(), data, false)
    if !isPtr {
        return reflect.ValueOf(value).Convert(tp).Interface(), err
    }
    val := reflect.New(tp)
    val.Elem().Set(reflect.ValueOf(value).Convert(tp))
    return val.Interface(), err
}

func baseDeserializeWith(kind, data string, isPtr bool) (interface{}, error) {
    switch kind {
    case reflect.String.String():
        if isPtr {
            return &data, nil
        }
        return data, nil
    case reflect.Bool.String():
        b, err := strconv.ParseBool(data)
        if isPtr {
            return &b, err
        }
        return b, err
    case reflect.Int.String():
        i64, err := strconv.ParseInt(data, 0, 0)
        i := int(i64)
        if isPtr {
            return &i, err
        }
        return i, err
    case reflect.Int8.String():
        i64, err := strconv.ParseInt(data, 0, 8)
        i8 := int8(i64)
        if isPtr {
            return &i8, err
        }
        return i8, err
    case reflect.Int16.String():
        i64, err := strconv.ParseInt(data, 0, 16)
        i16 := int16(i64)
        if isPtr {
            return &i16, err
        }
        return i16, err
    case reflect.Int32.String():
        i64, err := strconv.ParseInt(data, 0, 32)
        i32 := int32(i64)
        if isPtr {
            return &i32, err
        }
        return i32, err
    case reflect.Int64.String():
        i64, err := strconv.ParseInt(data, 0, 64)
        if isPtr {
            return &i64, err
        }
        return i64, err
    case reflect.Float32.String():
        f64, err := strconv.ParseFloat(data, 32)
        f32 := float32(f64)
        if isPtr {
            return &f32, err
        }
        return f32, err
    case reflect.Float64.String():
        f64, err := strconv.ParseFloat(data, 64)
        if isPtr {
            return &f64, err
        }
        return f64, err
    case reflect.Uint.String():
        u64, err := strconv.ParseUint(data, 0, 0)
        u := uint(u64)
        if isPtr {
            return &u, err
        }
        return u, err
    case reflect.Int8.String():
        u64, err := strconv.ParseUint(data, 0, 8)
        u8 := uint8(u64)
        if isPtr {
            return &u8, err
        }
        return u8, err
    case reflect.Int16.String():
        u64, err := strconv.ParseUint(data, 0, 16)
        u16 := uint16(u64)
        if isPtr {
            return &u16, err
        }
        return u16, err
    case reflect.Int32.String():
        u64, err := strconv.ParseUint(data, 0, 32)
        u32 := uint(u64)
        if isPtr {
            return &u32, err
        }
        return u32, err
    case reflect.Int64.String():
        u64, err := strconv.ParseUint(data, 0, 64)
        if isPtr {
            return &u64, err
        }
        return u64, err
    case reflect.Complex64.String():
        var (
            split = strings.LastIndexByte(data, '+') + strings.LastIndexByte(data, '-')
            a, _  = baseDeserializeWith("float32", data[:split], false)
            b, _  = baseDeserializeWith("float32", data[split:len(data)-1], false)
        )
        return complex(a.(float32), b.(float32)), nil
    case reflect.Complex128.String():
        var (
            split = strings.LastIndexByte(data, '+') + strings.LastIndexByte(data, '-')
            a, _  = baseDeserializeWith("float64", data[:split], false)
            b, _  = baseDeserializeWith("float64", data[split:len(data)-1], false)
        )
        return complex(a.(float64), b.(float64)), nil
    default:
        return nil, errors.New("unsupported kind:" + kind)
    }
}
