package utils

import "encoding/json"

type SchemaData struct {
	Model any
}

func NewSchemaData(OriginModel any) *SchemaData {
	return &SchemaData{
		Model: OriginModel,
	}
}

func (s *SchemaData) SchemaSwap(schema any) error {
	userJson, _ := json.Marshal(schema)
	err := json.Unmarshal(userJson, &s.Model)
	if err != nil {
		return err
	}
	return nil
}
