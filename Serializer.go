package serialize

import (
    "github.com/AMan4Technology/Serialize/internal"
)

func Serialize(value interface{}, name, tag, codecID string) (data string, err error) {
    return internal.Serialize(value, name, tag, codecID)
}

func Deserialize(data, tag, codecID string) (value interface{}, name string, err error) {
    return internal.Deserialize(data, tag, codecID)
}

type Serializer = internal.Serializer
