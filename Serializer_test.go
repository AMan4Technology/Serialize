package serialize

import (
    "fmt"
    "reflect"
    "testing"

    "github.com/AMan4Technology/Serialize/codec"
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
    fmt.Println(Serialize(&my{A: "WangTing", b: 1, C: new(string)}, codec.String, "test", "string"))
}

func TestDeserialize(t *testing.T) {
    fmt.Println(Deserialize("slice|data|github.com/AMan4Technology/Serialize/internal.StringSlice|sliceData|3|7|int|0|17|int|1|27|int|2|3", codec.String, ""))
    fmt.Println(Deserialize("map|data|github.com/AMan4Technology/Serialize/internal.StringSlice|mapData|2|98|github.com/AMan4Technology/Serialize/internal.StringSlice|keys|3|9|string||b9|string||c9|string||a91|github.com/AMan4Technology/Serialize/internal.StringSlice|values|3|6|int||26|int||36|int||1", codec.String, ""))
    fmt.Println(Deserialize("struct|data|github.com/AMan4Technology/Serialize/internal.StringSlice|fields|1|17|string|a|wangting", codec.String, "string"))
    fmt.Println(Deserialize("github.com/AMan4Technology/Serialize.my|test|github.com/AMan4Technology/Serialize/internal.StringSlice|fields|1|17|string|a|WangTing", codec.String, "string"))
    fmt.Println(Deserialize("*github.com/AMan4Technology/Serialize.my|test|github.com/AMan4Technology/Serialize/internal.StringSlice|fields|2|17|string|a|WangTing10|*string|C|", codec.String, "string"))
}

func TestIDOf(t *testing.T) {
    fmt.Println(IDOf(reflect.TypeOf(my{})))
    fmt.Println(IDOf(reflect.TypeOf(my{}.A)))
    fmt.Println(IDOf(reflect.TypeOf(my{}.C)))
}

type my struct {
    A string `string:"a"`
    b int
    C *string
}
