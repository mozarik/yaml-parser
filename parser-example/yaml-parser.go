package parser

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	InvalidOrUnsupportedDataType = errors.New("invalid or unsupported data type, data type given")
)

type YamlData struct {
	Name        string     `yaml:"name"`
	Type        string     `yaml:"type"`
	Description string     `yaml:"description"`
	RecordData  []YamlData `yaml:"fields"`
}

func (y YamlData) isOneOfDataType() error {
	switch DataType(y.Type) {
	case BOOLEAN, BYTES, DATE, DATETIME, FLOAT, GEOGRAPHY, INTEGER, NUMERIC, STRING, TIME, TIMESTAMP:
		return nil
	case RECORD:
		for _, value := range y.RecordData {
			err := value.isOneOfDataType()
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("%w %s", InvalidOrUnsupportedDataType, y.Type)
	}
}

type RecordData struct {
	YamlData []YamlData
}

type DataType string

// Refer to this https://cloud.google.com/bigquery/docs/data-types
const (
	BOOLEAN   DataType = "BOOLEAN"
	BYTES     DataType = "BYTES"
	DATE      DataType = "DATE"
	DATETIME  DataType = "DATETIME"
	FLOAT     DataType = "FLOAT"
	GEOGRAPHY DataType = "GEOGRAPHY"
	INTEGER   DataType = "INTEGER"
	NUMERIC   DataType = "NUMERIC"
	RECORD    DataType = "RECORD"
	STRING    DataType = "STRING"
	TIME      DataType = "TIME"
	TIMESTAMP DataType = "TIMESTAMP"
)

func ValidateYamlDataTypeActivity(data []YamlData) (err error) {
	for _, value := range data {
		err = value.isOneOfDataType()
		if err != nil {
			return fmt.Errorf("%w: %s", InvalidOrUnsupportedDataType, value.Type)
		}
	}
	return nil
}

func ParseToStructActivity(rawYaml []byte) (data []YamlData, err error) {
	err = yaml.Unmarshal(rawYaml, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
