package iaccount

import (
	"os"
	"testing"

	itest "github.com/davidtrse/go-clean-arch/integration_test"
)

func TestMain(m *testing.M) {
	itest.Setup()

	ret := m.Run()

	os.Exit(ret)
}
