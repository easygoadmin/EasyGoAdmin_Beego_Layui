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
	"easygoadmin/app/constant"
	"easygoadmin/app/dto"
	"easygoadmin/app/models"
	"easygoadmin/app/services"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gookit/validate"
)

var Menu = new(MenuController)

type MenuController struct {
	BaseController
}

func (ctl *MenuController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "menu/index.html"
}

func (ctl *MenuController) List() {
	// 参数对象
	var req dto.MenuQueryReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用查询方法
	list, err := services.Menu.GetList(req)
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

func (ctl *MenuController) Edit() {
	// 获取菜单列表
	menuTreeList, _ := services.Menu.GetTreeList()
	// 数据源转换
	menuList := services.Menu.MakeList(menuTreeList)
	// 记录ID
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.Menu{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 获取节点
		funcList := make([]models.Menu, 0)
		orm.NewOrm().QueryTable(new(models.Menu)).
			Filter("pid", id).
			Filter("type", 1).
			Filter("mark", 1).
			All(&funcList)
		sortList := make([]interface{}, 0)
		for _, v := range funcList {
			sortList = append(sortList, v.Sort)
		}

		// 渲染模板
		ctl.Data["info"] = info
		ctl.Data["funcList"] = sortList
	} else {
		// 添加
		pid, _ := ctl.GetInt("pid", 0)
		var info models.Menu
		info.Pid = pid
		info.Status = 1
		info.Target = 1
		ctl.Data["info"] = info
		ctl.Data["funcList"] = make([]interface{}, 0)
	}
	// 渲染模板
	ctl.Data["menuList"] = menuList
	ctl.Data["typeList"] = constant.MENU_TYPE_LIST
	ctl.Layout = "public/form.html"
	ctl.TplName = "menu/edit.html"
}

func (ctl *MenuController) Add() {
	// 参数对象
	var req dto.MenuAddReq
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
	rows, err := services.Menu.Add(req, utils.Uid(ctl.Ctx))
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

func (ctl *MenuController) Update() {
	// 参数对象
	var req dto.MenuUpdateReq
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
	rows, err := services.Menu.Update(req, utils.Uid(ctl.Ctx))
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

func (ctl *MenuController) Delete() {
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.Menu.Delete(ids)
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
