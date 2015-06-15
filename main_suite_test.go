package main_test

import (
. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"
"testing"
)

func TestMainApp(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Main App Suite")
}
