package codec

const Split = '|'

func Encode(codec, typeID, name, value string) (data string) {
    c := codecs[codec]
    if c == nil {
        c = codecs[String]
    }
    return c.Encode(typeID, name, value)
}

func Decode(codec, data string) (typeID, name, value string) {
    c := codecs[codec]
    if c == nil {
        c = codecs[String]
    }
    return c.Decode(data)
}

type Codec interface {
    Encode(typeID, name, value string) (data string)
    Decode(data string) (typeID, name, value string)
}