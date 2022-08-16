package types

type SerializedType interface {
	SerializeJson(json any) ([]byte, error)
}

func GetSerializedType(t string) SerializedType {
	switch t {
	case "UInt32":
		return &UInt32{}
	}
	return nil
}
