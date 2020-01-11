package serialize

import (
    "reflect"

    "github.com/AMan4Technology/Serialize/internal"
)

func Register(tp reflect.Type, s Serializer, update bool) (err error) {
    return internal.Register(tp, s, update)
}

func IDOf(tp reflect.Type) string {
    return internal.IDOf(tp)
}

func StringFrom(value reflect.Value) string {
    return internal.StringFrom(value)
}

func NumOfSerializers() int {
    return len(internal.SerializerWithID)
}

func RangeSerializers(callback func(name string) bool) {
    for name := range internal.SerializerWithID {
        if !callback(name) {
            break
        }
    }
}
