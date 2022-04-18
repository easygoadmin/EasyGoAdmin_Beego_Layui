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

package dto

import "github.com/gookit/validate"

// 字典项列表查询条件
type ConfigDataPageReq struct {
	ConfigId int    `form:"configId"` // 字典ID
	Title    string `form:"name"`     // 配置标题
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加字典项
type ConfigDataAddReq struct {
	Title    string `form:"Title" validate:"required"` // 配置标题
	Code     string `form:"Code" validate:"required"`  // 配置编码
	Value    string `form:"Value"`                     // 配置值
	Options  string `form:"Options"`                   // 配置项
	ConfigId int    `form:"ConfigId" validate:"int"`   // 配置ID
	Type     string `form:"Type" validate:"required"`  // 配置类型
	Sort     int    `form:"Sort" validate:"int"`       // 排序
	Note     string `form:"Note"`                      // 配置说明
}

// 添加配置项表单验证
func (a ConfigDataAddReq) Messages() map[string]string {
	return validate.MS{
		"Title.required": "配置项名称不能为空.",
		"Code.required":  "配置项编码不能为空.",
		"ConfigId.int":   "配置ID不能为空.",
		"Sort.int":       "排序不能为空.",
	}
}

// 更新字典项
type ConfigDataUpdateReq struct {
	Id       int    `form:"Id" validate:"int"`
	Title    string `form:"Title" validate:"required"` // 配置标题
	Code     string `form:"Code" validate:"required"`  // 配置编码
	Value    string `form:"Value"`                     // 配置值
	Options  string `form:"Options"`                   // 配置项
	ConfigId int    `form:"ConfigId" validate:"int"`   // 配置ID
	Type     string `form:"Type" validate:"required"`  // 配置类型
	Sort     int    `form:"Sort" validate:"int"`       // 排序
	Note     string `form:"Note"`                      // 配置说明
}

// 更新配置项表单验证
func (u ConfigDataUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":         "配置项ID不能为空.",
		"Title.required": "配置项名称不能为空.",
		"Code.required":  "配置项编码不能为空.",
		"ConfigId.int":   "配置ID不能为空.",
		"Sort.int":       "排序不能为空.",
	}
}

// 设置状态
type ConfigDataStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}