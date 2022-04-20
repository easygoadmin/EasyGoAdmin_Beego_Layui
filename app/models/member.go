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

type Member struct {
	Id           int       `orm:"column(id);auto"`
	Openid       string    `orm:"column(openid);size(50);null" description:"用户唯一标识"`
	Username     string    `orm:"column(username);size(30)" description:"用户名"`
	Password     string    `orm:"column(password);size(32);null" description:"密码"`
	MemberLevel  int       `orm:"column(member_level)" description:"会员等级"`
	Realname     string    `orm:"column(realname);size(50);null" description:"真实姓名"`
	Nickname     string    `orm:"column(nickname);size(50);null" description:"用户昵称"`
	Gender       int       `orm:"column(gender)" description:"性别（1男 2女 3未知）"`
	Avatar       string    `orm:"column(avatar);size(180);null" description:"用户头像"`
	Birthday     time.Time `orm:"column(birthday);type(date);null" description:"出生日期"`
	ProvinceCode string    `orm:"column(province_code);size(30);null" description:"户籍省份编号"`
	CityCode     string    `orm:"column(city_code);size(30);null" description:"户籍城市编号"`
	DistrictCode string    `orm:"column(district_code);size(30);null" description:"户籍区/县编号"`
	Address      string    `orm:"column(address);size(255);null" description:"详细地址"`
	Intro        string    `orm:"column(intro);null" description:"个人简介"`
	Signature    string    `orm:"column(signature);size(30);null" description:"个性签名"`
	Device       int       `orm:"column(device)" description:"设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加"`
	DeviceCode   string    `orm:"column(device_code);size(40);null" description:"推送的别名"`
	PushAlias    string    `orm:"column(push_alias);size(40);null" description:"推送的别名"`
	Source       int       `orm:"column(source)" description:"来源：1、APP注册；2、后台添加；"`
	Status       int       `orm:"column(status)" description:"是否启用 1、启用  2、停用"`
	AppVersion   string    `orm:"column(app_version);size(30);null" description:"客户端版本号"`
	Code         string    `orm:"column(code);size(10);null" description:"我的推广码"`
	LoginIp      string    `orm:"column(login_ip);size(30);null" description:"最近登录IP"`
	LoginTime    time.Time `orm:"column(login_time);type(datetime);null" description:"登录时间"`
	LoginRegion  string    `orm:"column(login_region);size(20);null" description:"上次登录地点"`
	LoginCount   int       `orm:"column(login_count)" description:"登录总次数"`
	CreateUser   int       `orm:"column(create_user)" description:"添加人"`
	CreateTime   time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateUser   int       `orm:"column(update_user)" description:"修改人"`
	UpdateTime   time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
	Mark         int       `orm:"column(mark)" description:"有效标识：1正常 0删除"`
}

func (t *Member) TableName() string {
	return "ums_member"
}

func init() {
	orm.RegisterModel(new(Member))
}

// 根据条件查询单条数据
func (t *Member) Get() error {
	err := orm.NewOrm().QueryTable(new(Member)).Filter("id", t.Id).One(t)
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
func (t *Member) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Member) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Member) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}
