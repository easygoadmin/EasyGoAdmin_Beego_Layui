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

package dto

import "github.com/gookit/validate"

// 列表查询条件
type MenuQueryReq struct {
	Name string `form:"name"` // 菜单标题
}

// 添加菜单
type MenuAddReq struct {
	Name       string `form:"name" validate:"required"` // 菜单标题
	Icon       string `form:"icon" `                    // 图标
	Url        string `form:"url"`                      // URL地址
	Pid        int    `form:"pid" validate:"int"`       // 上级ID
	Type       int    `form:"type" validate:"int"`      // 类型：1模块 2导航 3菜单 4节点
	Permission string `form:"permission"`               // 权限标识
	Status     int    `form:"status" validate:"int"`    // 状态：1正常 2禁用
	Target     int    `form:"target"`                   // 打开方式：1内部打开 2外部打开
	Note       string `form:"note"`                     // 菜单备注
	Sort       int    `form:"sort" validate:"int"`      // 显示顺序
	Func       string `form:"func"`                     // 权限节点
}

// 添加菜单表单验证
func (a MenuAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "菜单名称不能为空.",
		"Pid.int":       "请选择上级菜单.",
		"Type.int":      "请选择菜单类型.",
		"Status.int":    "请选择菜单状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新菜单
type MenuUpdateReq struct {
	Id         int    `form:"id" validate:"int"`
	Name       string `form:"name" validate:"required"` // 菜单标题
	Icon       string `form:"icon" `                    // 图标
	Url        string `form:"url"`                      // URL地址
	Pid        int    `form:"pid" validate:"int"`       // 上级ID
	Type       int    `form:"type" validate:"int"`      // 类型：1模块 2导航 3菜单 4节点
	Permission string `form:"permission"`               // 权限标识
	Status     int    `form:"status" validate:"int"`    // 状态：1正常 2禁用
	Target     int    `form:"target"`                   // 打开方式：1内部打开 2外部打开
	Note       string `form:"note"`                     // 菜单备注
	Sort       int    `form:"sort" validate:"int"`      // 显示顺序
	Func       string `form:"func"`                     // 权限节点
}

// 添加菜单表单验证
func (a MenuUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "菜单ID不能为空.",
		"Name.required": "菜单名称不能为空.",
		"Pid.int":       "请选择上级菜单.",
		"Type.int":      "请选择菜单类型.",
		"Status.int":    "请选择菜单状态.",
		"Sort.int":      "排序不能为空.",
	}
}
