package definitions

type fieldInstance struct {
	FieldName string
	*fieldInfo
	FieldHeader *fieldHeader
}

type fieldInfo struct {
	Nth            int
	IsVLEncoded    bool
	IsSerialized   bool
	IsSigningField bool
	Type           string
}

type fieldHeader struct {
	TypeCode  byte
	FieldCode byte
}
