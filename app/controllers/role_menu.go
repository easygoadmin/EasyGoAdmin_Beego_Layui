// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨
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
	"easygoadmin/utils/common"
	"github.com/gookit/validate"
)

var RoleMenu = new(RoleMenuController)

type RoleMenuController struct {
	BaseController
}

func (ctl *RoleMenuController) Index() {
	// 角色ID
	roleId, _ := ctl.GetInt("roleId", 0)
	if roleId <= 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "角色ID不能为空",
		})
	}

	// 获取角色菜单权限列表
	list, err := services.RoleMenu.GetRoleMenuList(roleId)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Data: list,
		Msg:  "查询成功",
	})
}

func (ctl *RoleMenuController) Save() {
	// 参数对象
	var req dto.RoleMenuSaveReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 参数验证
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}
	// 调用保存角色菜单数据
	err := services.RoleMenu.Save(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 保存成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "保存成功",
	})
}
