// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package controllers

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/services"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gookit/validate"
	"net/http"
)

var Index = new(IndexController)

type IndexController struct {
	BaseController
}

func (ctl *IndexController) Index() {
	// 获取用户信息
	userInfo := services.Login.GetProfile(utils.Uid(ctl.Ctx))
	// 获取菜单列表
	menuList := services.Menu.GetPermissionMenuList(userInfo.Id)
	ctl.Data["userInfo"] = userInfo
	ctl.Data["menuList"] = menuList
	// 渲染模板
	ctl.TplName = "index.html"
}

func (ctl *IndexController) Main() {
	// 渲染模板
	ctl.TplName = "welcome.html"
}

func (ctl *IndexController) UserInfo() {
	if ctl.Ctx.Input.IsPost() {
		// 参数验证
		var req dto.UserInfoReq
		// 参数绑定
		if err := ctl.ParseForm(&req); err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 参数校验
		v := validate.Struct(req)
		if !v.Validate() {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  v.Errors.One(),
			})
		}
		// 更新信息
		_, err := services.User.UpdateUserInfo(req, utils.Uid(ctl.Ctx))
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 返回结果
		ctl.JSON(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})
	}
	// 获取用户信息
	userInfo := services.Login.GetProfile(utils.Uid(ctl.Ctx))
	// 渲染模板
	ctl.Data["userInfo"] = userInfo
	ctl.TplName = "user/user_info.html"
}

func (ctl *IndexController) UpdatePwd() {
	// 更新密码对象
	var req dto.UpdatePwd
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 参数校验
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}

	// 调用更新密码方法
	rows, err := services.User.UpdatePwd(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "更新密码成功",
	})
}

// 注销系统
func (ctl *IndexController) Logout() {
	// 删除登录Session
	ctl.DelSession("userId")
	// 跳转登录页,方式：301(永久移动),308(永久重定向),307(临时重定向)
	ctl.Redirect("/login", http.StatusFound)
}
