package routers

import (
	"github.com/astaxie/beego"
	"material/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/index.html", &controllers.IndexController{})
	beego.Router("/page-login.html", &controllers.LoginController{})
	beego.Router("/page-palette-amber.html", &controllers.AmberController{})
	beego.Router("/page-palette-brand-accent.html", &controllers.AccentController{})
	beego.Router("/page-palette-brand.html", &controllers.BrandController{})
	beego.Router("/page-palette-green.html", &controllers.GreenController{})
	beego.Router("/page-palette-red.html", &controllers.RedController{})
	beego.Router("/page-palette.html", &controllers.PaletteController{})
	beego.Router("/ui-button.html", &controllers.ButtonController{})
	beego.Router("/ui-card.html", &controllers.UiCardController{})
	beego.Router("/ui-collapse.html", &controllers.UiCollapseController{})
	beego.Router("/ui-dropdown.html", &controllers.UiDropdownController{})
	beego.Router("/ui-form-adv.html", &controllers.UiFormAdvController{})
	beego.Router("/ui-form-default.html", &controllers.UiFormDefaultController{})
	beego.Router("/ui-form.html", &controllers.UiFormController{})
	beego.Router("/ui-icon.html", &controllers.UiIconController{})
	beego.Router("/ui-modal.html", &controllers.UiModalController{})
	beego.Router("/ui-nav.html", &controllers.UiNavController{})
	beego.Router("/ui-progress.html", &controllers.UiProgressController{})
	beego.Router("/ui-tab.html", &controllers.UiTabController{})
	beego.Router("/ui-table.html", &controllers.UiTableController{})
	beego.Router("/ui-tile.html", &controllers.UiTitleController{})
}
