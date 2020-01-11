package internal

import "reflect"

func defaultSerialize(value reflect.Value, tag, codecID string, varMap StringMap) (string, error) {
    switch kind := value.Kind(); kind {
    case reflect.Slice, reflect.Array:
        return sliceSerialize(value, tag, codecID, varMap), nil
    case reflect.Map:
        return mapSerialize(value, tag, codecID, varMap), nil
    case reflect.Struct:
        return structSerialize(value, tag, codecID, varMap), nil
    default:
        return baseSerialize(value), nil
    }
}

func defaultDeserialize(tp reflect.Type, data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (reflect.Value, error) {
    switch kind := tp.Kind().String(); kind {
    case reflect.Array.String():
        return sliceDeserialize(tp, data, tag, codecID, varMap, ptrMap)
    case reflect.Slice.String():
        ptr, data := data, varMap[data]
        if ptr, ok := ptrMap[ptr]; ok {
            return ptr, nil
        }
        value, err := sliceDeserialize(tp, data, tag, codecID, varMap, ptrMap)
        ptrMap[ptr] = value
        return value, err
    case reflect.Map.String():
        ptr, data := data, varMap[data]
        if ptr, ok := ptrMap[ptr]; ok {
            return ptr, nil
        }
        value, err := mapDeserialize(tp, data, tag, codecID, varMap, ptrMap)
        ptrMap[ptr] = value
        return value, err
    case reflect.Struct.String():
        return structDeserialize(tp, data, tag, codecID, varMap, ptrMap)
    default:
        return baseDeserialize(tp, data)
    }
}

func defaultDeserializeWith(kind, data, tag, codecID string, varMap StringMap, ptrMap map[string]reflect.Value) (reflect.Value, error) {
    switch kind {
    case reflect.Array.String():
        return sliceDeserializeWith(data, tag, codecID, varMap, ptrMap)
    case reflect.Slice.String():
        ptr, data := data, varMap[data]
        if ptr, ok := ptrMap[ptr]; ok {
            return ptr, nil
        }
        value, err := sliceDeserializeWith(data, tag, codecID, varMap, ptrMap)
        ptrMap[ptr] = value
        return value, err
    case reflect.Map.String():
        ptr, data := data, varMap[data]
        if ptr, ok := ptrMap[ptr]; ok {
            return ptr, nil
        }
        value, err := mapDeserializeWith(data, tag, codecID, varMap, ptrMap)
        ptrMap[ptr] = value
        return value, err
    case reflect.Struct.String():
        return structDeserializeWithMap(data, tag, codecID, varMap, ptrMap)
    default:
        value, err := baseDeserializeWith(kind, data)
        return reflect.ValueOf(value), err
    }
}
