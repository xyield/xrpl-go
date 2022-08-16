package types

type SerializedType interface {
	SerializeJson(json any) ([]byte, error)
}

func GetSerializedType(t string) SerializedType {
	switch t {
	case "UInt32":
		return &UInt32{}
	case "UInt16":
		return &UInt16{}
	case "UInt64":
		return &UInt64{}
	case "Hash256":
		return &Hash256{}
	}
	return nil
}
