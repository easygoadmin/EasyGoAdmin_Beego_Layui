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

package services

import (
	"easygoadmin/app/models"
	"easygoadmin/conf"
	"easygoadmin/utils"
	"easygoadmin/utils/gstr"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"time"
)

var Login = new(loginService)

type loginService struct{}

// 系统登录
func (s *loginService) UserLogin(username, password string, ctx *context.Context) (*models.User, error) {
	// 查询用户
	var user models.User
	err := orm.NewOrm().QueryTable(new(models.User)).Filter("username", username).Filter("mark", 1).One(&user)
	if err != nil {
		return nil, errors.New("用户名或者密码不正确")
	}
	// 密码校验
	pwd, _ := utils.Md5(password + user.Username)
	if user.Password != pwd {
		return nil, errors.New("密码不正确")
	}
	// 判断当前用户状态
	if user.Status != 1 {
		return nil, errors.New("您的账号已被禁用,请联系管理员")
	}

	// 更新登录时间、登录IP
	o := orm.NewOrm()
	entity := models.User{Id: user.Id}
	entity.LoginTime = time.Now()
	entity.LoginIp = ctx.Input.IP()
	entity.UpdateTime = time.Now()
	o.Update(&entity, "LoginTime", "LoginIp", "UpdateTime")

	// 结果
	return &user, nil
}

// 获取个人信息
func (s *loginService) GetProfile(userId int) *models.User {
	user := &models.User{Id: userId}
	err := user.Get()
	if err != nil {
		return nil
	}
	// 头像
	if user.Avatar != "" && !gstr.Contains(user.Avatar, conf.CONFIG.EGAdmin.Image) {
		user.Avatar = utils.GetImageUrl(user.Avatar)
	}
	return user
}
