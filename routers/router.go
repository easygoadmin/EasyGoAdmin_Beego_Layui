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

package routers

import (
	"easygoadmin/app/controllers"
	"easygoadmin/app/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	// 登录验证中间件
	middleware.CheckLogin()
	// 初始化加载组件
	middleware.LoadWidget()

	// 系统登录
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/captcha", &controllers.LoginController{}, "get:Captcha")

	// 系统主页
	beego.Router("/index", &controllers.IndexController{}, "get:Index")
	beego.Router("/main", &controllers.IndexController{}, "get:Main")
	beego.Router("/userInfo", &controllers.IndexController{}, "get,post:UserInfo")
	beego.Router("/updatePwd", &controllers.IndexController{}, "get,post:UpdatePwd")
	beego.Router("/logout", &controllers.IndexController{}, "get:Logout")

	// 普通图片上传
	beego.Router("/upload/uploadImage", &controllers.UploadController{}, "post:UploadImage")
	// 富文本图片上传
	beego.Router("/upload/uploadEditImage", &controllers.UploadController{}, "post:UploadEditImage")

	// 职级管理
	beego.Router("/level/index", &controllers.LevelController{}, "get:Index")
	beego.Router("/level/list", &controllers.LevelController{}, "post:List")
	beego.Router("/level/edit", &controllers.LevelController{}, "get:Edit")
	beego.Router("/level/add", &controllers.LevelController{}, "post:Add")
	beego.Router("/level/update", &controllers.LevelController{}, "post:Update")
	beego.Router("/level/delete", &controllers.LevelController{}, "post:Delete")
	beego.Router("/level/setStatus", &controllers.LevelController{}, "post:Status")

	// 岗位管理
	beego.Router("/position/index", &controllers.PositionController{}, "get:Index")
	beego.Router("/position/list", &controllers.PositionController{}, "post:List")
	beego.Router("/position/edit", &controllers.PositionController{}, "get:Edit")
	beego.Router("/position/add", &controllers.PositionController{}, "post:Add")
	beego.Router("/position/update", &controllers.PositionController{}, "post:Update")
	beego.Router("/position/delete", &controllers.PositionController{}, "post:Delete")
	beego.Router("/position/setStatus", &controllers.PositionController{}, "post:Status")

	// 角色管理
	beego.Router("/role/index", &controllers.RoleController{}, "get:Index")
	beego.Router("/role/list", &controllers.RoleController{}, "post:List")
	beego.Router("/role/edit", &controllers.RoleController{}, "get:Edit")
	beego.Router("/role/add", &controllers.RoleController{}, "post:Add")
	beego.Router("/role/update", &controllers.RoleController{}, "post:Update")
	beego.Router("/role/delete", &controllers.RoleController{}, "post:Delete")
	beego.Router("/role/setStatus", &controllers.RoleController{}, "post:Status")

	// 角色菜单
	beego.Router("/rolemenu/index", &controllers.RoleMenuController{}, "get:Index")
	beego.Router("/rolemenu/save", &controllers.RoleMenuController{}, "post:Save")

	// 菜单管理
	beego.Router("/menu/index", &controllers.MenuController{}, "get:Index")
	beego.Router("/menu/list", &controllers.MenuController{}, "post:List")
	beego.Router("/menu/edit", &controllers.MenuController{}, "get:Edit")
	beego.Router("/menu/add", &controllers.MenuController{}, "post:Add")
	beego.Router("/menu/update", &controllers.MenuController{}, "post:Update")
	beego.Router("/menu/delete", &controllers.MenuController{}, "post:Delete")

	// 部门管理
	beego.Router("/dept/index", &controllers.DeptController{}, "get:Index")
	beego.Router("/dept/list", &controllers.DeptController{}, "post:List")
	beego.Router("/dept/edit", &controllers.DeptController{}, "get:Edit")
	beego.Router("/dept/add", &controllers.DeptController{}, "post:Add")
	beego.Router("/dept/update", &controllers.DeptController{}, "post:Update")
	beego.Router("/dept/delete", &controllers.DeptController{}, "post:Delete")

	// 用户管理
	beego.Router("/user/index", &controllers.UserController{}, "get:Index")
	beego.Router("/user/list", &controllers.UserController{}, "post:List")
	beego.Router("/user/edit", &controllers.UserController{}, "get:Edit")
	beego.Router("/user/add", &controllers.UserController{}, "post:Add")
	beego.Router("/user/update", &controllers.UserController{}, "post:Update")
	beego.Router("/user/delete", &controllers.UserController{}, "post:Delete")
	beego.Router("/user/resetPwd", &controllers.UserController{}, "post:ResetPwd")

	// 城市管理
	beego.Router("/city/index", &controllers.CityController{}, "get:Index")
	beego.Router("/city/list", &controllers.CityController{}, "post:List")
	beego.Router("/city/edit", &controllers.CityController{}, "get:Edit")
	beego.Router("/city/add", &controllers.CityController{}, "post:Add")
	beego.Router("/city/update", &controllers.CityController{}, "post:Update")
	beego.Router("/city/delete", &controllers.CityController{}, "post:Delete")
	beego.Router("/city/getChilds", &controllers.CityController{}, "post:GetChilds")

	// 会员等级管理
	beego.Router("/memberlevel/index", &controllers.MemberLevelController{}, "get:Index")
	beego.Router("/memberlevel/list", &controllers.MemberLevelController{}, "post:List")
	beego.Router("/memberlevel/edit", &controllers.MemberLevelController{}, "get:Edit")
	beego.Router("/memberlevel/add", &controllers.MemberLevelController{}, "post:Add")
	beego.Router("/memberlevel/update", &controllers.MemberLevelController{}, "post:Update")
	beego.Router("/memberlevel/delete", &controllers.MemberLevelController{}, "post:Delete")

	// 会员管理
	beego.Router("/member/index", &controllers.MemberController{}, "get:Index")
	beego.Router("/member/list", &controllers.MemberController{}, "post:List")
	beego.Router("/member/edit", &controllers.MemberController{}, "get:Edit")
	beego.Router("/member/add", &controllers.MemberController{}, "post:Add")
	beego.Router("/member/update", &controllers.MemberController{}, "post:Update")
	beego.Router("/member/delete", &controllers.MemberController{}, "post:Delete")
	beego.Router("/member/setStatus", &controllers.MemberController{}, "post:Status")

	// 友链管理
	beego.Router("/link/index", &controllers.LinkController{}, "get:Index")
	beego.Router("/link/list", &controllers.LinkController{}, "post:List")
	beego.Router("/link/edit", &controllers.LinkController{}, "get:Edit")
	beego.Router("/link/add", &controllers.LinkController{}, "post:Add")
	beego.Router("/link/update", &controllers.LinkController{}, "post:Update")
	beego.Router("/link/delete", &controllers.LinkController{}, "post:Delete")
	beego.Router("/link/setStatus", &controllers.LinkController{}, "post:Status")

	// 字典管理
	beego.Router("/dict/index", &controllers.DictController{}, "get:Index")
	beego.Router("/dict/list", &controllers.DictController{}, "post:List")
	beego.Router("/dict/add", &controllers.DictController{}, "post:Add")
	beego.Router("/dict/update", &controllers.DictController{}, "post:Update")
	beego.Router("/dict/delete", &controllers.DictController{}, "post:Delete")

	// 字典数据
	beego.Router("/dictdata/list", &controllers.DictDataController{}, "post:List")
	beego.Router("/dictdata/add", &controllers.DictDataController{}, "post:Add")
	beego.Router("/dictdata/update", &controllers.DictDataController{}, "post:Update")
	beego.Router("/dictdata/delete", &controllers.DictDataController{}, "post:Delete")

	// 配置管理
	beego.Router("/config/index", &controllers.ConfigController{}, "get:Index")
	beego.Router("/config/list", &controllers.ConfigController{}, "post:List")
	beego.Router("/config/add", &controllers.ConfigController{}, "post:Add")
	beego.Router("/config/update", &controllers.ConfigController{}, "post:Update")
	beego.Router("/config/delete", &controllers.ConfigController{}, "post:Delete")

	// 配置数据
	beego.Router("/configdata/list", &controllers.ConfigDataController{}, "post:List")
	beego.Router("/configdata/add", &controllers.ConfigDataController{}, "post:Add")
	beego.Router("/configdata/update", &controllers.ConfigDataController{}, "post:Update")
	beego.Router("/configdata/delete", &controllers.ConfigDataController{}, "post:Delete")
	beego.Router("/configdata/setStatus", &controllers.ConfigDataController{}, "post:Status")

	// 网站配置
	beego.Router("/configweb/index", &controllers.ConfigWebController{}, "get,post:Index")

	// 统计分析
	beego.Router("/analysis/index", &controllers.AnalysisController{}, "get:Index")

	// 通知公告管理
	beego.Router("/notice/index", &controllers.NoticeController{}, "get:Index")
	beego.Router("/notice/list", &controllers.NoticeController{}, "post:List")
	beego.Router("/notice/edit", &controllers.NoticeController{}, "get:Edit")
	beego.Router("/notice/add", &controllers.NoticeController{}, "post:Add")
	beego.Router("/notice/update", &controllers.NoticeController{}, "post:Update")
	beego.Router("/notice/delete", &controllers.NoticeController{}, "post:Delete")
	beego.Router("/notice/setStatus", &controllers.NoticeController{}, "post:Status")

	// 代码生成器
	beego.Router("/generate/index", &controllers.GenerateController{}, "get:Index")
	beego.Router("/generate/list", &controllers.GenerateController{}, "post:List")
	beego.Router("/generate/generate", &controllers.GenerateController{}, "post:Generate")
	beego.Router("/generate/batchGenerate", &controllers.GenerateController{}, "post:BatchGenerate")

	// 站点管理
	beego.Router("/item/index", &controllers.ItemController{}, "get:Index")
	beego.Router("/item/list", &controllers.ItemController{}, "post:List")
	beego.Router("/item/edit", &controllers.ItemController{}, "get:Edit")
	beego.Router("/item/add", &controllers.ItemController{}, "post:Add")
	beego.Router("/item/update", &controllers.ItemController{}, "post:Update")
	beego.Router("/item/delete", &controllers.ItemController{}, "post:Delete")
	beego.Router("/item/setStatus", &controllers.ItemController{}, "post:Status")

	// 栏目管理
	beego.Router("/itemcate/index", &controllers.ItemCateController{}, "get:Index")
	beego.Router("/itemcate/list", &controllers.ItemCateController{}, "post:List")
	beego.Router("/itemcate/edit", &controllers.ItemCateController{}, "get:Edit")
	beego.Router("/itemcate/add", &controllers.ItemCateController{}, "post:Add")
	beego.Router("/itemcate/update", &controllers.ItemCateController{}, "post:Update")
	beego.Router("/itemcate/delete", &controllers.ItemCateController{}, "post:Delete")
	beego.Router("/itemcate/getCateTreeList", &controllers.ItemCateController{}, "get:GetCateTreeList")

	// 广告位管理
	beego.Router("/adsort/index", &controllers.AdSortController{}, "get:Index")
	beego.Router("/adsort/list", &controllers.AdSortController{}, "post:List")
	beego.Router("/adsort/edit", &controllers.AdSortController{}, "get:Edit")
	beego.Router("/adsort/add", &controllers.AdSortController{}, "post:Add")
	beego.Router("/adsort/update", &controllers.AdSortController{}, "post:Update")
	beego.Router("/adsort/delete", &controllers.AdSortController{}, "post:Delete")

	// 广告管理
	beego.Router("/ad/index", &controllers.AdController{}, "get:Index")
	beego.Router("/ad/list", &controllers.AdController{}, "post:List")
	beego.Router("/ad/edit", &controllers.AdController{}, "get:Edit")
	beego.Router("/ad/add", &controllers.AdController{}, "post:Add")
	beego.Router("/ad/update", &controllers.AdController{}, "post:Update")
	beego.Router("/ad/delete", &controllers.AdController{}, "post:Delete")
	beego.Router("/ad/setStatus", &controllers.AdController{}, "post:Status")
}
