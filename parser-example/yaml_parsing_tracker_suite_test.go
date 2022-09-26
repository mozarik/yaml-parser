package parser_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestYamlParsingTracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "YamlParsingTracker Suite")
}
