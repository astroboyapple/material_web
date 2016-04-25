package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.TplNames = "index.html"
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "page-login.html"
}

func (this *LoginController) Post() {
	// jsoninfo := this.GetString("jsoninfo")
	// if jsoninfo == "" {
	// 	this.Ctx.WriteString("jsoninfo is empty")
	// 	return
	// }
	this.TplNames = "index.html"
}

type AmberController struct {
	beego.Controller
}

func (this *AmberController) Get() {
	this.TplNames = "page-palette-amber.html"
}

type AccentController struct {
	beego.Controller
}

func (this *AccentController) Get() {
	this.TplNames = "page-palette-brand-accent.html"
}

type BrandController struct {
	beego.Controller
}

func (this *BrandController) Get() {
	this.TplNames = "page-palette-brand.html"
}

type GreenController struct {
	beego.Controller
}

func (this *GreenController) Get() {
	this.TplNames = "page-palette-green.html"
}

type RedController struct {
	beego.Controller
}

func (this *RedController) Get() {
	this.TplNames = "page-palette-red.html"
}

type PaletteController struct {
	beego.Controller
}

func (this *PaletteController) Get() {
	this.TplNames = "page-palette.html"
}

type ButtonController struct {
	beego.Controller
}

func (this *ButtonController) Get() {
	this.TplNames = "ui-button.html"
}

type UiCardController struct {
	beego.Controller
}

func (this *UiCardController) Get() {
	this.TplNames = "ui-card.html"
}

type UiCollapseController struct {
	beego.Controller
}

func (this *UiCollapseController) Get() {
	this.TplNames = "ui-collapse.html"
}

type UiDropdownController struct {
	beego.Controller
}

func (this *UiDropdownController) Get() {
	this.TplNames = "ui-dropdown.html"
}

type UiFormAdvController struct {
	beego.Controller
}

func (this *UiFormAdvController) Get() {
	this.TplNames = "ui-form-adv.html"
}

type UiFormDefaultController struct {
	beego.Controller
}

func (this *UiFormDefaultController) Get() {
	this.TplNames = "ui-form-default.html"
}

type UiFormController struct {
	beego.Controller
}

func (this *UiFormController) Get() {
	this.TplNames = "ui-form.html"
}

type UiIconController struct {
	beego.Controller
}

func (this *UiIconController) Get() {
	this.TplNames = "ui-icon.html"
}

type UiModalController struct {
	beego.Controller
}

func (this *UiModalController) Get() {
	this.TplNames = "ui-modal.html"
}

type UiNavController struct {
	beego.Controller
}

func (this *UiNavController) Get() {
	this.TplNames = "ui-nav.html"
}

type UiProgressController struct {
	beego.Controller
}

func (this *UiProgressController) Get() {
	this.TplNames = "ui-progress.html"
}

type UiTabController struct {
	beego.Controller
}

func (this *UiTabController) Get() {
	this.TplNames = "ui-tab.html"
}

type UiTableController struct {
	beego.Controller
}

func (this *UiTableController) Get() {
	this.TplNames = "ui-table.html"
}

type UiTitleController struct {
	beego.Controller
}

func (this *UiTitleController) Get() {
	this.TplNames = "ui-tile.html"
}
