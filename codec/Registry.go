package codec

func Register(name string, codec Codec) {
    codecs[name] = codec
}

var codecs = make(map[string]Codec)
