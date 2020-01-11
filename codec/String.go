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

func (stringCodec) Encode(name, typeID, value string) (data string) {
    return fmt.Sprintf("%s%c%s%c%s", name, Split, typeID, Split, value)
}

func (stringCodec) Decode(data string) (name, typeID, value string) {
    var (
        one = strings.IndexByte(data, Split)
        two = one + 1 + strings.IndexByte(data[one+1:], Split)
    )
    if one == -1 {
        return "", "", data
    }
    if two == one {
        return "", data[:one], data[one+1:]
    }
    return data[:one], data[one+1 : two], data[two+1:]
}
