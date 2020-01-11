package internal

import (
    "errors"
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

func Register(tp reflect.Type, s Serializer, update bool) (err error) {
    if SerializerWithID[IDOf(tp)] != nil && !update {
        return errors.New("type has been registered")
    }
    if tp.Kind() == reflect.Ptr {
        return Register(tp.Elem(), s, update)
    }
    err = register(IDOf(tp), tp, s, update)
    switch tp.Kind() {
    case reflect.Struct:
        for i := 0; i < tp.NumField(); i++ {
            _ = Register(tp.Field(i).Type, nil, false)
        }
    case reflect.Slice, reflect.Array:
        _ = Register(tp.Elem(), nil, false)
    case reflect.Map:
        _, _ = Register(tp.Key(), nil, false), Register(tp.Elem(), nil, false)
    }
    return
}

func IDOf(tp reflect.Type) string {
    if pkgPath := strings.TrimSpace(tp.PkgPath()); pkgPath != "" {
        return fmt.Sprintf("%s.%s", pkgPath, strings.TrimSpace(tp.Name()))
    }
    return tp.Kind().String()
}

func StringFrom(value reflect.Value) (ptr string) {
    switch value.Kind() {
    case reflect.Map, reflect.Ptr, reflect.Slice:
        ptr = strconv.FormatUint(uint64(value.Pointer()), 16)
    }
    return
}

var SerializerWithID = make(map[string]*serializer)

func register(id string, tp reflect.Type, s Serializer, update bool) (err error) {
    if SerializerWithID[id] != nil && !update {
        return fmt.Errorf("serializer %s is exist", id)
    }
    SerializerWithID[id] = &serializer{ID: id, TP: tp, Serializer: s}
    return nil
}
