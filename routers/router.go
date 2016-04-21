// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"net/http"

	"github.com/ysqi/splitimage/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/image/split", &controllers.ImageController{}, "*:Split")
	beego.ErrorHandler("404", func(rw http.ResponseWriter, req *http.Request) {
		http.NotFound(rw, req)
	})
}
