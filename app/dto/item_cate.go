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

// 栏目查询条件
type ItemCateQueryReq struct {
	Name string `form:"name"` // 栏目名称
}

// 添加栏目
type ItemCateAddReq struct {
	Name    string `form:"name" validate:"required"`   // 栏目名称
	Pid     int    `form:"pid" validate:"int"`         // 父级ID
	ItemId  int    `form:"itemId" validate:"int"`      // 栏目ID
	Pinyin  string `form:"pinyin" validate:"required"` // 拼音(全)
	Code    string `form:"code" validate:"required"`   // 拼音(简)
	IsCover int    `form:"isCover" validate:"int"`     // 是否有封面：1是 2否
	Cover   string `form:"cover"`                      // 封面
	Status  int    `form:"status" validate:"int"`      // 状态：1启用 2停用
	Note    string `form:"note"`                       // 备注
	Sort    int    `form:"sort" validate:"int"`        // 排序
}

// 添加栏目表单验证
func (a ItemCateAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required":   "栏目名称不能为空.",
		"Pid.int":         "请选择上级栏目.",
		"ItemId.int":      "请选择栏目ID.",
		"Pinyin.required": "拼音全拼不能为空.",
		"Code.required":   "拼音简拼不能为空.",
		"IsCover.int":     "请选择是否有封面.",
		"Status.int":      "请选择栏目状态.",
		"Sort.int":        "排序不能为空.",
	}
}

// 修改栏目
type ItemCateUpdateReq struct {
	Id      int    `form:"id" validate:"int"`
	Name    string `form:"name" validate:"required"`   // 栏目名称
	Pid     int    `form:"pid" validate:"int"`         // 父级ID
	ItemId  int    `form:"itemId" validate:"int"`      // 栏目ID
	Pinyin  string `form:"pinyin" validate:"required"` // 拼音(全)
	Code    string `form:"code" validate:"required"`   // 拼音(简)
	IsCover int    `form:"isCover" validate:"int"`     // 是否有封面：1是 2否
	Cover   string `form:"cover"`                      // 封面
	Status  int    `form:"status" validate:"int"`      // 状态：1启用 2停用
	Note    string `form:"note"`                       // 备注
	Sort    int    `form:"sort" validate:"int"`        // 排序
}

// 更新栏目表单验证
func (a ItemCateUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":          "栏目ID不能为空.",
		"Name.required":   "栏目名称不能为空.",
		"Pid.int":         "请选择上级栏目.",
		"ItemId.int":      "请选择栏目ID.",
		"Pinyin.required": "拼音全拼不能为空.",
		"Code.required":   "拼音简拼不能为空.",
		"IsCover.int":     "请选择是否有封面.",
		"Status.int":      "请选择栏目状态.",
		"Sort.int":        "排序不能为空.",
	}
}
