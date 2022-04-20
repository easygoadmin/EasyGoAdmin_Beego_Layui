// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id           int       `orm:"column(id);auto" description:"主键ID"`
	Realname     string    `orm:"column(realname);size(150);null" description:"真实姓名"`
	Nickname     string    `orm:"column(nickname);size(150);null" description:"昵称"`
	Gender       int       `orm:"column(gender);null" description:"性别:1男 2女 3保密"`
	Avatar       string    `orm:"column(avatar);size(150);null" description:"头像"`
	Mobile       string    `orm:"column(mobile);size(11);null" description:"手机号码"`
	Email        string    `orm:"column(email);size(30);null" description:"邮箱地址"`
	Birthday     time.Time `orm:"column(birthday);type(date);null" description:"出生日期"`
	DeptId       int       `orm:"column(dept_id);null" description:"部门ID"`
	LevelId      int       `orm:"column(level_id);null" description:"职级ID"`
	PositionId   int       `orm:"column(position_id);null" description:"岗位ID"`
	ProvinceCode string    `orm:"column(province_code);size(50);null" description:"省份编号"`
	CityCode     string    `orm:"column(city_code);size(50);null" description:"市区编号"`
	DistrictCode string    `orm:"column(district_code);size(50);null" description:"区县编号"`
	Address      string    `orm:"column(address);size(255);null" description:"详细地址"`
	CityName     string    `orm:"column(city_name);size(150);null" description:"所属城市"`
	Username     string    `orm:"column(username);size(50);null" description:"登录用户名"`
	Password     string    `orm:"column(password);size(150);null" description:"登录密码"`
	Salt         string    `orm:"column(salt);size(30);null" description:"盐加密"`
	Intro        string    `orm:"column(intro);size(500);null" description:"个人简介"`
	Status       int       `orm:"column(status);null" description:"状态：1正常 2禁用"`
	Note         string    `orm:"column(note);size(500);null" description:"备注"`
	Sort         int       `orm:"column(sort);null" description:"排序号"`
	LoginNum     int       `orm:"column(login_num);null" description:"登录次数"`
	LoginIp      string    `orm:"column(login_ip);size(20);null" description:"最近登录IP"`
	LoginTime    time.Time `orm:"column(login_time);type(datetime);null" description:"最近登录时间"`
	CreateUser   int       `orm:"column(create_user);null" description:"添加人"`
	CreateTime   time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateUser   int       `orm:"column(update_user);null" description:"更新人"`
	UpdateTime   time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
	Mark         int       `orm:"column(mark)" description:"有效标识(1正常 0删除)"`
}

func (t *User) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User))
}

// 根据条件查询单条数据
func (t *User) Get() error {
	err := orm.NewOrm().QueryTable(new(User)).Filter("id", t.Id).One(t)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return errors.New("查询到了多条记录")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return errors.New("未查询到记录")
	}
	return nil
}

// 插入数据
func (t *User) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *User) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *User) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}
