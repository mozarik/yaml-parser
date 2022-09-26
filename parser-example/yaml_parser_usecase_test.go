package parser_test

import (
	"log"
	"yaml-parsing-tracker/parser-example"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var VALID_testDataOneValue = []byte(`
- name: field_1
  type: STRING
  description: description_1
`)

var VALID_testDataTwovalue = []byte(`
- name: field_1
  type: STRING
  description: description_1
- name: field_2
  type: BOOLEAN
  description: description_2
`)

var VALID_testDataNestedRecord = []byte(`
- name: field_1
  type: RECORD
  description: description_1
  fields:
      - name: field_1_1
        type: STRING
        description: description_1_1
      - name: field_1_2
        type: STRING
        description: description_1_2
      - name: field_1_3
        type: INTEGER
        description: description_1_3
`)

var VALID_testDataNestedRecordModeRepeated = []byte(`
- name: field_1
  type: RECORD
  mode: REPEATED
  description: description_1
  fields:
      - name: field_1_1
        type: STRING
        description: description_1_1
      - name: field_1_2
        type: STRING
        description: description_1_2
      - name: field_1_3
        type: INTEGER
        description: description_1_3
`)

var INVALID_string_testDataOneValue = []byte(`
- name: field_1
  type: string
  description: description_1
`)

var INVALID_string_testDataTwoValue = []byte(`
- name: field_1
  type: STRING
  description: description_1
- name: field_2
  type: boolean
  description: description_2
`)

var INVALID_bool_testDataNestedRecord = []byte(`
- name: field_1
  type: RECORD
  description: description_1
  fields:
      - name: field_1_1
        type: STRING
        description: description_1_1
      - name: field_1_2
        type: STRING
        description: description_1_2
      - name: field_1_3
        type: bool
        description: description_1_3
`)

var _ = Describe("YamlParserUsecase", func() {
	Context("Parsing raw string byte to struct", func() {
		It("can decode to a simple struct", func() {
			parsedYamlToStruct, err := parser.ParseToStructActivity(VALID_testDataOneValue)
			Expect(err).To(BeNil())
			for _, value := range parsedYamlToStruct {
				Expect(value.Name).To(Equal("field_1"))
				Expect(value.Type).To(Equal("STRING"))
				Expect(value.Description).To(Equal("description_1"))
			}
		})

		It("can decode to a simple struct with two value", func() {
			expectedData1 := parser.YamlData{
				Name:        "field_1",
				Type:        "STRING",
				Description: "description_1",
			}
			expectedData2 := parser.YamlData{
				Name:        "field_2",
				Type:        "BOOLEAN",
				Description: "description_2",
			}

			parsedYamlToStruct, err := parser.ParseToStructActivity(VALID_testDataTwovalue)
			Expect(err).To(BeNil())
			Expect(parsedYamlToStruct).Should(HaveLen(2))
			Expect(parsedYamlToStruct).Should(ContainElements(expectedData1, expectedData2))
		})

		It("can decode to a nested struct", func() {
			expectedData1 := []parser.YamlData{
				{
					Name:        "field_1",
					Type:        "RECORD",
					Description: "description_1",
					RecordData: []parser.YamlData{
						{
							Name:        "field_1_1",
							Type:        "STRING",
							Description: "description_1_1",
						},
						{
							Name:        "field_1_2",
							Type:        "STRING",
							Description: "description_1_2",
						},
						{
							Name:        "field_1_3",
							Type:        "INTEGER",
							Description: "description_1_3",
						},
					},
				},
			}

			parsedYamlToStruct, err := parser.ParseToStructActivity(VALID_testDataNestedRecord)
			Expect(err).To(BeNil())
			Expect(parsedYamlToStruct).Should(HaveLen(1))
			Expect(parsedYamlToStruct).Should(BeEquivalentTo(expectedData1))
		})

		It("Introduce field `mode` but no check", func() {
			expectedData1 := []parser.YamlData{
				{
					Name:        "field_1",
					Type:        "RECORD",
					Description: "description_1",
					RecordData: []parser.YamlData{
						{
							Name:        "field_1_1",
							Type:        "STRING",
							Description: "description_1_1",
						},
						{
							Name:        "field_1_2",
							Type:        "STRING",
							Description: "description_1_2",
						},
						{
							Name:        "field_1_3",
							Type:        "INTEGER",
							Description: "description_1_3",
						},
					},
				},
			}

			parsedYamlToStruct, err := parser.ParseToStructActivity(VALID_testDataNestedRecordModeRepeated)
			Expect(err).To(BeNil())
			Expect(parsedYamlToStruct).Should(HaveLen(1))
			Expect(parsedYamlToStruct).Should(BeEquivalentTo(expectedData1))
		})
	})

	Context("Check Data Type", func() {
		When("Given correct data", func() {
			It("can validate type is one of data type supported by given one data", func() {
				parsedYamlToStruct, err := parser.ParseToStructActivity(VALID_testDataOneValue)
				Expect(err).To(BeNil())
				err = parser.ValidateYamlDataTypeActivity(parsedYamlToStruct)
				Expect(err).To(BeNil())
			})
		})

		When("Given incorrect data type", func() {
			It("given one data with type string", func() {
				parsedYamlToStruct, err := parser.ParseToStructActivity(INVALID_string_testDataOneValue)
				Expect(err).To(BeNil())
				err = parser.ValidateYamlDataTypeActivity(parsedYamlToStruct)
				Expect(err).ToNot(BeNil())
			})

			It("given two data, valid DataType 'STRING', and InvalidData 'boolean'", func() {
				parsedYamlToStruct, err := parser.ParseToStructActivity(INVALID_string_testDataTwoValue)
				Expect(err).To(BeNil())
				err = parser.ValidateYamlDataTypeActivity(parsedYamlToStruct)
				log.Println(err)
				Expect(err).ToNot(BeNil())
			})

			It("given nested data, valid DataType 'RECORD' but nested InvalidData 'bool'", func() {
				parsedYamlToStruct, err := parser.ParseToStructActivity(INVALID_bool_testDataNestedRecord)
				Expect(err).To(BeNil())
				err = parser.ValidateYamlDataTypeActivity(parsedYamlToStruct)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
