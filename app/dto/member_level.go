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

// 查询会员等级
type MemberLevelPageReq struct {
	Name  string `form:"name"`  // 等级名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加会员等级
type MemberLevelAddReq struct {
	Name string `form:"name" validate:"required"` // 级别名称
	Sort int    `form:"sort" validate:"int"`      // 排序号
}

// 添加会员等级表单验证
func (a MemberLevelAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "会员等级名称不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新会员等级
type MemberLevelUpdateReq struct {
	Id   int    `form:"id" validate:"int"`
	Name string `form:"name" validate:"required"` // 级别名称
	Sort int    `form:"sort" validate:"int"`      // 排序号
}

// 添加会员等级表单验证
func (a MemberLevelUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "会员等级ID不能为空.",
		"Name.required": "会员等级名称不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}
