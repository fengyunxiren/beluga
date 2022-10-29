package server

import (
	v1 "beluga/application/api/v1"
	"beluga/server/common/apirouter"
)

var ServerAPIRouters apirouter.IAPIRouters = apirouter.IAPIRouters{
	v1.TestAPI,
}
