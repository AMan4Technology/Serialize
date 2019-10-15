package internal

import (
    "fmt"
    "reflect"
    "strings"
)

func Register(tp reflect.Type, s Serializer, update bool) (err error) {
    switch tp.Kind() {
    case reflect.Struct:
        for i := 0; i < tp.NumField(); i++ {
            _ = Register(tp.Field(i).Type, nil, false)
        }
    case reflect.Slice, reflect.Array:
        _ = Register(tp.Elem(), nil, false)
    case reflect.Map:
        _, _ = Register(tp.Key(), nil, false), Register(tp.Elem(), nil, false)
    case reflect.Ptr:
        _ = Register(tp.Elem(), nil, false)
        return
    }
    return register(IDOf(tp), tp, s, update)
}

func IDOf(tp reflect.Type) string {
    if pkgPath := strings.TrimSpace(tp.PkgPath()); pkgPath != "" {
        return fmt.Sprintf("%s.%s", pkgPath, strings.TrimSpace(tp.Name()))
    }
    return tp.Kind().String()
}

var SerializerWithID = make(map[string]*serializer)

func register(id string, tp reflect.Type, s Serializer, update bool) (err error) {
    if SerializerWithID[id] != nil && !update {
        return fmt.Errorf("serializer %s is exist", id)
    }
    SerializerWithID[id] = &serializer{ID: id, TP: tp, Serializer: s}
    return nil
}
