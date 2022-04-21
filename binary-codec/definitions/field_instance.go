package definitions

type fieldInstance struct {
	FieldName   string
	FieldInfo   fieldInfo
	FieldHeader fieldHeader
}

type fieldInfo struct {
	Nth            int64
	IsVLEncoded    bool
	IsSerialized   bool
	IsSigningField bool
	Type           string
}

type fieldHeader struct {
	TypeCode  byte
	FieldCode byte
}
