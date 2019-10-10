package internal

import (
	"reflect"
	"strconv"
	"strings"

	"Serialize/codec"
)

func init() {
	_ = Register(reflect.TypeOf(StringSlice{}), StringSliceSerializer{}, false)
}

const Nil = "nil"

type StringSlice []string

type StringSliceSerializer struct{}

func (StringSliceSerializer) Serialize(value interface{}, tag string) (string, error) {
	stringSlice := value.(StringSlice)
	if stringSlice == nil {
		return Nil, nil
	}
	var (
		length = len(stringSlice)
		data   = strings.Builder{}
	)
	data.WriteString(strconv.Itoa(length))
	data.WriteByte(codec.Split)
	for _, s := range stringSlice {
		data.WriteString(strconv.Itoa(len(s)))
		data.WriteByte(codec.Split)
		data.WriteString(s)
	}
	return data.String(), nil
}

func (StringSliceSerializer) Deserialize(data string, tag string) (interface{}, error) {
	var stringSlice StringSlice
	if strings.EqualFold(data, Nil) {
		return stringSlice, nil
	}
	index := strings.IndexByte(data, codec.Split)
	length, err := strconv.Atoi(data[:index])
	if err != nil {
		return stringSlice, err
	}
	stringSlice = make(StringSlice, length)
	index++
	for i := 0; i < length; i++ {
		next := index + strings.IndexByte(data[index:], codec.Split)
		length, err := strconv.Atoi(data[index:next])
		if err != nil {
			return stringSlice, err
		}
		index = next + 1 + length
		stringSlice[i] = data[next+1 : index]
	}
	return stringSlice, nil
}
