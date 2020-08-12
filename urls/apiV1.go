package urls

import (
	"web-zjh/controllers"
	"web-zjh/mainStructs"
)

var apiV1 = map[string]mainStructs.Controller{
	"index": controllers.IndexController,
}
