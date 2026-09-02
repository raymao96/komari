package notification

import (
	"context"

	"github.com/raymao96/komari/database/notificationdefaults"
)

func init() {
	notificationdefaults.RegisterTrafficReportRetentionReconciler(func() error {
		return EnsureTrafficReportMetricRetention(context.Background())
	})
}
