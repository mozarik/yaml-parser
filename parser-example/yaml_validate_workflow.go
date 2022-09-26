package parser

func ValidateYamlWorkflow(yamlData []byte) error {
	data, err := ParseToStructActivity(yamlData)
	if err != nil {
		return err
	}
	err = ValidateYamlDataTypeActivity(data)
	if err != nil {
		return err
	}
	return nil
}
