package Services

import (
	"testing"

	 . "github.com/onsi/ginkgo"
	 . "github.com/onsi/gomega"
)

func TestJWTService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JWT Service Test Suite")
}
