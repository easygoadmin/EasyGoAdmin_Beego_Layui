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

/**
 * 演示一-路由
 * @author 半城风雨
 * @since 2022-04-15
 * @File : example
 */
package routers

import (
	"easygoadmin/app/controllers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	fmt.Println("模块路由初始化")

	/* 演示一 */
	beego.Router("/example/index", &controllers.ExampleController{}, "get:Index")
	beego.Router("/example/list", &controllers.ExampleController{}, "post:List")
	beego.Router("/example/edit", &controllers.ExampleController{}, "get:Edit")
	beego.Router("/example/add", &controllers.ExampleController{}, "post:Add")
	beego.Router("/example/update", &controllers.ExampleController{}, "post:Update")
	beego.Router("/example/delete", &controllers.ExampleController{}, "post:Delete")
	beego.Router("/example/setStatus", &controllers.ExampleController{}, "post:Status")

	beego.Router("/example/setStatus", &controllers.ExampleController{}, "post:Status")
	beego.Router("/example/setIsVip", &controllers.ExampleController{}, "post:IsVip")
}
