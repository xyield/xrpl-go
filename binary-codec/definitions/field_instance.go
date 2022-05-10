package definitions

type fieldInstance struct {
	FieldName string
	*fieldInfo
	FieldHeader *fieldHeader
}

type fieldInfo struct {
	Nth            int32
	IsVLEncoded    bool
	IsSerialized   bool
	IsSigningField bool
	Type           string
}

type fieldHeader struct {
	TypeCode  int32
	FieldCode int32
}

func CreateFieldHeader(tc, fc int32) fieldHeader {
	return fieldHeader{
		TypeCode:  tc,
		FieldCode: fc,
	}
}
