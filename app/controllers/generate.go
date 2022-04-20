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
	"easygoadmin/utils/gconv"
	"github.com/gookit/validate"
	"strconv"
	"strings"
)

var Generate = new(GenerateController)

type GenerateController struct {
	BaseController
}

func (ctl *GenerateController) Index() {
	// 渲染模板
	ctl.Layout = "public/layout.html"
	ctl.TplName = "generate/index.html"
}

func (ctl *GenerateController) List() {
	// 参数验证
	var req dto.GeneratePageReq
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用查询列表方法
	list, err := services.Generate.GetList(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  list,
		Count: gconv.Int64(len(list)),
	})
}

func (ctl *GenerateController) Generate() {
	// 生成对象
	var req dto.GenerateFileReq
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
	// 调用生成方法
	err := services.Generate.Generate(req, ctl.Ctx)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "生成成功",
	})
}

func (ctl *GenerateController) BatchGenerate() {
	// 生成对象
	var req dto.BatchGenerateFileReq
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
	// 参数分析
	tableList := strings.Split(req.Tables, ",")
	count := 0
	for _, item := range tableList {
		itemList := strings.Split(item, "|")
		// 组装参数对象
		var param dto.GenerateFileReq
		param.Name = itemList[0]
		param.Comment = itemList[1]
		// 调用生成方法
		err := services.Generate.Generate(param, ctl.Ctx)
		if err != nil {
			continue
		}
		count++
	}
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "本次共生成【" + strconv.Itoa(count) + "】个模块文件",
	})
}
