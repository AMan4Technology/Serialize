package internal

import (
    "fmt"
    "testing"

    "github.com/AMan4Technology/Serialize/codec"
)

func Test_stringSliceSerializer_Serialize(t *testing.T) {
    data := StringSlice{"WangTing", "1", "2"}
    fmt.Println(Serialize(data, "sliceData", "", codec.String))
}

func Test_StringSliceSerializer_Deserialize(t *testing.T) {
    data := "2|86|target||sliceData|github.com/AMan4Technology/Serialize/internal.StringSlice|c00006066030|c000060660||3|8|WangTing1|11|2"
    serializable, _, err := Deserialize(data, "", codec.String)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(serializable.(StringSlice))
}
