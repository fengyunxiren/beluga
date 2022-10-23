package server

import (
	v1 "beluga/application/api/v1"
	"beluga/server/common/app"
)

var ServerApps app.IApps = app.IApps{
	v1.TestApi,
}
