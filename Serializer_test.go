package serialize

import (
	"fmt"
	"reflect"
	"testing"

	"Serialize/codec"
)

func init() {
	_ = Register(reflect.TypeOf(my{}), nil, true)
}

func TestSerialize(t *testing.T) {
	fmt.Println(Serialize([]int{1, 2, 3}, codec.String, "data", ""))
	fmt.Println(Serialize(map[string]int{"a": 1, "b": 2, "c": 3}, codec.String, "data", ""))
	fmt.Println(Serialize(struct {
		A string `string:"a"`
		a int
	}{A: "wangting", a: 24}, codec.String, "data", "string"))
	fmt.Println(Serialize(my{A: "WangTing", b: 1}, codec.String, "test", "string"))
}

func TestDeserialize(t *testing.T) {
	fmt.Println(Deserialize("slice|data|Serialize/internal.StringSlice|sliceData|3|7|int|0|17|int|1|27|int|2|3", codec.String, ""))
	fmt.Println(Deserialize("map|data|Serialize/internal.StringSlice|mapData|2|71|Serialize/internal.StringSlice|keys|3|9|string||a9|string||b9|string||c64|Serialize/internal.StringSlice|values|3|6|int||16|int||26|int||3", codec.String, ""))
	fmt.Println(Deserialize("struct|data|Serialize/internal.StringSlice|fields|1|17|string|a|wangting", codec.String, "string"))
	fmt.Println(Deserialize("Serialize.my|test|Serialize/internal.StringSlice|fields|2|17|string|a|WangTing11|ptr|C|<nil>", codec.String, "string"))
}

func TestIDOf(t *testing.T) {
	value := reflect.New(reflect.TypeOf(my{}))
	c := reflect.New(reflect.TypeOf("c"))
	c.Elem().SetString("test")
	value.Elem().FieldByName("C").Set(c.Elem().Addr())
	fmt.Println(*value.Elem().Interface().(my).C)
}

type my struct {
	A string `string:"a"`
	b int
	C *string
}
