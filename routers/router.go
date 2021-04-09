package routers

import (
	"monolit/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "GET:Home")
	beego.Router("/add", &controllers.MainController{}, "GET:Add")
	beego.Router("/edit/:id", &controllers.MainController{}, "GET:Edit")

	beego.Router("/api/list-data-pegawai", &controllers.MainController{}, "GET:ReadDataPegawai")
	beego.Router("/api/hapus-data-pegawai/:id", &controllers.MainController{}, "DELETE:DeleteDataPegawai")
	beego.Router("/api/buat-data-pegawai", &controllers.MainController{}, "POST:CreateDataPegawai")
	beego.Router("/api/ubah-data-pegawai", &controllers.MainController{}, "PUT:UpdateDataPegawai")
	beego.Router("/api/list-data-atasan", &controllers.MainController{}, "GET:ReadDataAtasan")
}
