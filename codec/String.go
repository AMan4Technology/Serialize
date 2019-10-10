package codec

import (
    "fmt"
    "strings"
)

func init() {
    Register(String, stringCodec{})
}

const String = "string"

type stringCodec struct{}

func (stringCodec) Encode(typeID, name, value string) (data string) {
    return fmt.Sprintf("%s%c%s%c%s", typeID, Split, name, Split, value)
}

func (stringCodec) Decode(data string) (typeID, name, value string) {
    var (
        one = strings.IndexByte(data, Split)
        two = one + 1 + strings.IndexByte(data[one+1:], Split)
    )
    return data[:one], data[one+1 : two], data[two+1:]
}
