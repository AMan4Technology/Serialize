package serialize

import (
    "github.com/AMan4Technology/Serialize/internal"
)

func Serialize(value interface{}, codecID, name, tag string) (data string, err error) {
    return internal.Serialize(value, codecID, name, tag)
}

func Deserialize(data string, codecID, tag string) (value interface{}, name string, err error) {
    return internal.Deserialize(data, codecID, tag)
}

type Serializer = internal.Serializer
