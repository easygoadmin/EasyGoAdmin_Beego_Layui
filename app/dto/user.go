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

import (
	"github.com/gookit/validate"
)

// 用户分页查询条件
type UserPageReq struct {
	Realname string `form:"realname"` // 用户名
	Gender   int    `form:"gender"`   // 性别
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加用户
type UserAddReq struct {
	Realname     string `form:"realname" validate:"required"`
	Nickname     string `form:"nickname" validate:"required"`
	Gender       int    `form:"gender" validate:"int"`
	Avatar       string `form:"avatar" validate:"required"`
	Mobile       string `form:"mobile" validate:"required"`
	Email        string `form:"email" validate:"required"`
	Birthday     string `form:"birthday" validate:"required"`
	DeptId       int    `form:"deptId" validate:"int"`
	LevelId      int    `form:"levelId" validate:"int"`
	PositionId   int    `form:"positionId" validate:"int"`
	ProvinceCode string `form:"provinceCode"` // 省份编号
	CityCode     string `form:"cityCode"`     // 市区编号
	DistrictCode string `form:"districtCode"` // 区县编号
	Address      string `form:"address"`
	Username     string `form:"username" validate:"required"`
	Password     string `form:"password"`
	Intro        string `form:"intro"`
	Status       int    `form:"status" validate:"required"`
	Note         string `form:"note"`
	Sort         int    `form:"sort" validate:"required"`
	RoleIds      string `form:"roleIds"` // 用户角色
}

// 添加用户表单验证
func (a UserAddReq) Messages() map[string]string {
	return validate.MS{
		"Realname.required": "用户名称不能为空.",
		"Nickname.required": "用户昵称不能为空.",
		"Gender.int":        "请选择用户性别.",
		"Avatar.required":   "请上传头像.",
		"Mobile.required":   "手机号码不能为空.",
		"Email.required":    "电子邮件不能为空.",
		"Birthday.required": "请选择出生日期.",
		"DeptId.int":        "请选择所属部门.",
		"LevelId.int":       "请选择职级.",
		"PositionId.int":    "请选择用户.",
		"Username.required": "用户名不能为空.",
		"Status.int":        "请选择用户状态.",
		"Sort.int":          "排序不能为空.",
	}
}

// 添加用户
type UserUpdateReq struct {
	Id           int    `form:"id" validate:"int"`
	Realname     string `form:"realname" validate:"required"`
	Nickname     string `form:"nickname" validate:"required"`
	Gender       int    `form:"gender" validate:"int"`
	Avatar       string `form:"avatar" validate:"required"`
	Mobile       string `form:"mobile" validate:"required"`
	Email        string `form:"email" validate:"required"`
	Birthday     string `form:"birthday" validate:"required"`
	DeptId       int    `form:"deptId" validate:"int"`
	LevelId      int    `form:"levelId" validate:"int"`
	PositionId   int    `form:"positionId" validate:"int"`
	ProvinceCode string `form:"provinceCode"` // 省份编号
	CityCode     string `form:"cityCode"`     // 市区编号
	DistrictCode string `form:"districtCode"` // 区县编号
	Address      string `form:"address"`
	Username     string `form:"username" validate:"required"`
	Password     string `form:"password"`
	Intro        string `form:"intro"`
	Status       int    `form:"status" validate:"required"`
	Note         string `form:"note"`
	Sort         int    `form:"sort" validate:"required"`
	RoleIds      string `form:"roleIds"` // 用户角色
}

// 更新用户表单验证
func (u UserUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":            "用户ID不能为空.",
		"Realname.required": "用户名称不能为空.",
		"Nickname.required": "用户昵称不能为空.",
		"Gender.int":        "请选择用户性别.",
		"Avatar.required":   "请上传头像.",
		"Mobile.required":   "手机号码不能为空.",
		"Email.required":    "电子邮件不能为空.",
		"Birthday.required": "请选择出生日期.",
		"DeptId.int":        "请选择所属部门.",
		"LevelId.int":       "请选择职级.",
		"PositionId.int":    "请选择用户.",
		"Username.required": "用户名不能为空.",
		"Status.int":        "请选择用户状态.",
		"Sort.int":          "排序不能为空.",
	}
}

// 设置状态
type UserStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status"    validate:"int"`
}

func (s UserStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "用户ID不能为空.",
		"Status.int": "请选择用户状态.",
	}
}

// 用户中心
type UserInfoReq struct {
	Avatar   string `form:"avatar"`                       // 头像
	Realname string `form:"realname" validate:"required"` // 真实姓名
	Nickname string `form:"nickname" validate:"required"` // 昵称
	Gender   string `form:"gender" validate:"required"`   // 性别:1男 2女 3保密
	Mobile   string `form:"mobile" validate:"required"`   // 手机号码
	Email    string `form:"email" validate:"required"`    // 邮箱地址
	Address  string `form:"address"`                      // 详细地址
	Intro    string `form:"intro"`                        // 个人简介
}

// 更新个人中心表单验证
func (u UserInfoReq) Messages() map[string]string {
	return validate.MS{
		"Realname.required": "用户名称不能为空.",
		"Nickname.required": "用户昵称不能为空.",
		"Gender.required":   "请选择用户性别.",
		"Mobile.required":   "手机号码不能为空.",
		"Email.required":    "电子邮件不能为空.",
		"Address.required":  "用户名不能为空.",
	}
}

// 更新密码
type UpdatePwd struct {
	OldPassword string `form:"oldPassword" validate:"required"` // 旧密码
	NewPassword string `form:"newPassword" validate:"required"` // 新密码
	RePassword  string `form:"rePassword" validate:"required"`  // 确认密码
}

// 更新密码表单验证
func (u UpdatePwd) Messages() map[string]string {
	return validate.MS{
		"OldPassword.required": "旧密码不能为空.",
		"NewPassword.required": "新密码不能为空.",
		"RePassword.required":  "确认密码不能为空.",
	}
}

// 重置密码
type UserResetPwdReq struct {
	Id int `form:"id" validate:"int"`
}

// 更新密码表单验证
func (r UserResetPwdReq) Messages() map[string]string {
	return validate.MS{
		"Id.int": "用户ID不能为空.",
	}
}
