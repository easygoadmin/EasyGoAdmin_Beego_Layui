// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
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

type ItemCate struct {
	Id         int       `orm:"column(id);auto" description:"ID"`
	Name       string    `orm:"column(name);size(150);null" description:"栏目名称"`
	Pid        int       `orm:"column(pid);null" description:"父级ID"`
	ItemId     int       `orm:"column(item_id);null" description:"站点ID"`
	Pinyin     string    `orm:"column(pinyin);size(50);null" description:"拼音(全)"`
	Code       string    `orm:"column(code);size(10);null" description:"拼音(简)"`
	IsCover    int       `orm:"column(is_cover)" description:"是否有封面：1是 2否"`
	Cover      string    `orm:"column(cover);size(50);null" description:"封面"`
	Status     int       `orm:"column(status);null" description:"状态：1启用 2停用"`
	Note       string    `orm:"column(note);size(200);null" description:"备注"`
	Sort       int       `orm:"column(sort);null" description:"排序"`
	CreateUser int       `orm:"column(create_user);null" description:"添加人"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"添加时间"`
	UpdateUser int       `orm:"column(update_user);null" description:"更新人"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
	Mark       int       `orm:"column(mark)" description:"有效标识(1正常 0删除)"`
}

func (t *ItemCate) TableName() string {
	return "sys_item_cate"
}

func init() {
	orm.RegisterModel(new(ItemCate))
}

// 根据条件查询单条数据
func (t *ItemCate) Get() error {
	err := orm.NewOrm().QueryTable(new(ItemCate)).Filter("id", t.Id).One(t)
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
func (t *ItemCate) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *ItemCate) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *ItemCate) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}
