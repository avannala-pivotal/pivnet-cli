package release_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCommands(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Release suite")
}
