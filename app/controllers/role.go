// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package controllers

import (
	"easygoadmin/app/dto"
	"easygoadmin/app/models"
	"easygoadmin/app/services"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/gookit/validate"
)

var Role = new(RoleController)

type RoleController struct {
	BaseController
}

func (ctl *RoleController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "role/index.html"
}

func (ctl *RoleController) List() {
	// 参数对象
	var req dto.RolePageReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用查询列表方法
	list, count, err := services.Role.GetList(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "查询成功",
		Count: count,
	})
}
func (ctl *RoleController) Edit() {
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.Role{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 渲染模板
		ctl.Data["info"] = info
	}
	// 渲染模板
	ctl.Layout = "public/form.html"
	ctl.TplName = "role/edit.html"
}

func (ctl *RoleController) Add() {
	// 参数对象
	var req dto.RoleAddReq
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
	// 调用添加方法
	rows, err := services.Role.Add(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 添加成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (ctl *RoleController) Update() {
	// 参数对象
	var req dto.RoleUpdateReq
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
	// 调用更新方法
	rows, err := services.Role.Update(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 更新成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (ctl *RoleController) Delete() {
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.Role.Delete(ids)
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 删除成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

func (ctl *RoleController) Status() {
	// 参数对象
	var req dto.RoleStatusReq
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
	// 调用设置状态方法
	rows, err := services.Role.Status(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 设置成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}
