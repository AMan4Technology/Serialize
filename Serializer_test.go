package serialize

import (
    "fmt"
    "reflect"
    "testing"

    "github.com/AMan4Technology/Serialize/codec"
)

func init() {
    _ = Register(reflect.TypeOf(mys{}), nil, true)
}

func TestSerialize(t *testing.T) {
    fmt.Println(Serialize([]int{1, 2, 3}, "data", "", codec.String))
    fmt.Println(Serialize(map[string]int{"a": 1, "b": 2, "c": 3}, "data", "", codec.String))
    fmt.Println(Serialize(struct {
        A string `string:"a"`
        a int
    }{A: "wangting", a: 24}, "data", "string", codec.String))
    m := my{A: "WangTing", b: 1, C: new(string)}
    fmt.Println(Serialize(m, "m", "string", codec.String))
    fmt.Println(Serialize(&m, "pOfM", "string", codec.String))
    fmt.Println(Serialize(mys{nil, &m}, "mys", "string", codec.String))
}

func TestDeserialize(t *testing.T) {
    fmt.Println(Deserialize("2|29|target||data|slice|c00001048041|c000010480||3|7|0|int|17|1|int|27|2|int|3", "", codec.String))
    fmt.Println(Deserialize("2|81|c000060750||2|35|3|9||string|a9||string|b9||string|c26|3|6||int|16||int|26||int|327|target||data|map|c000060750", "", codec.String))
    fmt.Println(Deserialize("1|42|target||data|struct|1|17|a|string|wangting", "string", codec.String))
    fmt.Println(Deserialize("2|20|c000046990|||string|95|target||m|github.com/AMan4Technology/Serialize.my|2|17|a|string|WangTing20|C|*string|c000046990", "string", codec.String))
    fmt.Println(Deserialize("3|20|c000046990|||string|98|c000004760|||github.com/AMan4Technology/Serialize.my|2|17|a|string|WangTing20|C|*string|c00004699064|target||pOfM|*github.com/AMan4Technology/Serialize.my|c000004760", "string", codec.String))
    fmt.Println(Deserialize("4|20|c000046990|||string|98|c000004760|||github.com/AMan4Technology/Serialize.my|2|17|a|string|WangTing20|C|*string|c00004699080|c000046c30||2|8|0|*ptr|053|1|*github.com/AMan4Technology/Serialize.my|c00000476063|target||mys|github.com/AMan4Technology/Serialize.mys|c000046c30", "string", codec.String))
}

func TestIDOf(t *testing.T) {
    fmt.Println(IDOf(reflect.TypeOf(my{})))
    fmt.Println(IDOf(reflect.TypeOf(my{}.A)))
    fmt.Println(IDOf(reflect.TypeOf(my{}.C)))
}

type mys []*my

type my struct {
    A string `string:"a"`
    b int
    C *string
}
