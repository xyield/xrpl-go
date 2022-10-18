package types

type SerializedType interface {
	SerializeJson(json any) ([]byte, error)
}

// Returns the serialized type for the given type description.
func GetSerializedType(t string) SerializedType {
	switch t {
	case "UInt8":
		return &UInt8{}
	case "UInt16":
		return &UInt16{}
	case "UInt32":
		return &UInt32{}
	case "UInt64":
		return &UInt64{}
	case "Hash128":
		return NewHash128()
	case "Hash160":
		return NewHash160()
	case "Hash256":
		return NewHash256()
	case "AccountID":
		return &AccountID{}
	case "Amount":
		return &Amount{}
	case "Vector256":
		return &Vector256{}
	}
	return nil
}
