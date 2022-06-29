package definitions

func init() {
	err := loadDefinitions()
	if err != nil {
		return
	}
}
