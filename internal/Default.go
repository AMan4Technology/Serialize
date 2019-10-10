package internal

import "reflect"

func defaultSerialize(value interface{}, tag string) (string, error) {
    switch kind := reflect.ValueOf(value).Kind(); kind {
    case reflect.Slice, reflect.Array:
        return sliceSerialize(value, tag), nil
    case reflect.Map:
        return mapSerialize(value, tag), nil
    case reflect.Struct:
        return structSerialize(value, tag), nil
    default:
        return baseSerialize(value), nil
    }
}

func defaultDeserialize(tp reflect.Type, data string, tag string) (interface{}, error) {
    switch kind := tp.Kind().String(); kind {
    case reflect.Slice.String(), reflect.Array.String():
        return sliceDeserialize(tp, data, tag)
    case reflect.Map.String():
        return mapDeserialize(tp, data, tag)
    case reflect.Struct.String():
        return structDeserialize(tp, data, tag)
    default:
        return baseDeserialize(tp, data)
    }
}

func defaultDeserializeWith(kind, data, tag string) (interface{}, error) {
    switch kind {
    case reflect.Slice.String(), reflect.Array.String():
        return sliceDeserializeWith(data, tag)
    case reflect.Map.String():
        return mapDeserializeWith(data, tag)
    case reflect.Struct.String():
        return structDeserializeWithMap(data, tag)
    default:
        return baseDeserializeWith(kind, data)
    }
}
