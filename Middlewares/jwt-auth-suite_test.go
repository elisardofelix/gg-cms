package Middlewares

import (
"testing"

. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"
)

func TestJWTMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JWT-Auth Middleware Test Suite")
}
