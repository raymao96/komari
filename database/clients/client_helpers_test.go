package clients

import (
	"time"

	"github.com/raymao96/komari/database/billing"
	"gorm.io/gorm"
)

func saveClient(db *gorm.DB, updates map[string]interface{}) error {
	return saveClientWithSource(db, updates, billing.PriceSourceClientEdit)
}

func saveClientInfo(db *gorm.DB, update map[string]interface{}) error {
	return saveClientInfoWithAutoOrder(db, update, true)
}

func currentTrafficCycle(day *int, now time.Time) string {
	return currentTrafficCycleAt(day, "", "", now)
}
