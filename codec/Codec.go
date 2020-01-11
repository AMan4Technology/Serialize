package codec

const Split = '|'

func Encode(codec, name, typeID, value string) (data string) {
    c := codecs[codec]
    if c == nil {
        c = codecs[String]
    }
    return c.Encode(name, typeID, value)
}

func Decode(codec, data string) (name, typeID, value string) {
    if data == "" {
        return "", "", ""
    }
    c := codecs[codec]
    if c == nil {
        c = codecs[String]
    }
    return c.Decode(data)
}

type Codec interface {
    Encode(name, typeID, value string) (data string)
    Decode(data string) (name, typeID, value string)
}
