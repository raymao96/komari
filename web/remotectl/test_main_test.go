package remotectl

import (
	"os"
	"testing"

	"github.com/raymao96/komari/utils/instancekey"
)

func TestMain(m *testing.M) {
	cleanup := instancekey.SetupTempFileForTest()
	code := m.Run()
	cleanup()
	os.Exit(code)
}
