package serialize

import (
    "fmt"
    "reflect"
    "testing"

    "github.com/AMan4Technology/Serialize/internal"
)

func TestRegister(t *testing.T) {
    fmt.Println(internal.Register(reflect.TypeOf(new(node)), nil, true))
    fmt.Println(internal.Register(reflect.TypeOf(new(val)), nil, true))
    fmt.Println(internal.SerializerWithID)
}

type node struct {
    next *node
}

type val int
