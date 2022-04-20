package definitions

type fieldInfo struct {
	Nth            int64
	IsVLEncoded    bool
	IsSerialized   bool
	IsSigningField bool
	Type           string
}
