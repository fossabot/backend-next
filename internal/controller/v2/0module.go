package v2

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("controllers.v2", fx.Invoke(
		RegisterItem,
		RegisterZone,
		RegisterStage,
		RegisterNotice,
		RegisterResult,
		RegisterReport,
		RegisterAccount,
		RegisterFormula,
		RegisterPrivate,
		RegisterSiteStats,
		RegisterEventPeriod,
		RegisterShortURL,
	))
}
