package public

import (
	"os"
	"testing"

	"github.com/raymao96/komari/cmd/flags"
	"github.com/raymao96/komari/database/dbcore"
	"github.com/raymao96/komari/utils/instancekey"
)

func TestMain(m *testing.M) {
	cleanup := instancekey.SetupTempFileForTest()
	flags.DatabaseType = flags.DatabaseTypeSQLite
	flags.DatabaseFile = "file:web_api_public_test?mode=memory&cache=shared"

	db := dbcore.GetDBInstance()
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxOpenConns(1)
	}

	code := m.Run()
	cleanup()
	os.Exit(code)
}
