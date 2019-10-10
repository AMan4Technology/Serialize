package internal

import (
    "fmt"
    "reflect"
    "strings"
)

func Register(tp reflect.Type, s Serializer, update bool) (err error) {
    return register(IDOf(tp), tp, s, update)
}

func IDOf(tp reflect.Type) string {
    return fmt.Sprintf("%s.%s", strings.TrimSpace(tp.PkgPath()), strings.TrimSpace(tp.Name()))
}

var SerializerWithID = make(map[string]*serializer)

func register(id string, tp reflect.Type, s Serializer, update bool) (err error) {
    if SerializerWithID[id] != nil && !update {
        return fmt.Errorf("serializer %s is exist", id)
    }
    SerializerWithID[id] = &serializer{ID: id, TP: tp, Serializer: s}
    return nil
}
