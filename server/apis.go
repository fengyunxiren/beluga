package server

import (
	"beluga/application/api/admin"
	v1 "beluga/application/api/v1"
	"beluga/server/common/apirouter"
)

var ServerAPIRouters apirouter.IAPIRouters = apirouter.IAPIRouters{
	v1.APIV1,
	admin.Admin,
}
