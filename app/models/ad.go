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

package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Ad struct {
	Id          int       `orm:"column(id);auto" description:"主键ID"`
	Title       string    `orm:"column(title);size(100)" description:"广告标题"`
	AdSortId    int       `orm:"column(ad_sort_id)" description:"广告位ID"`
	Cover       string    `orm:"column(cover);size(255);null" description:"广告图片"`
	Type        int       `orm:"column(type)" description:"广告格式：1图片 2文字 3视频 4推荐"`
	Description string    `orm:"column(description);size(150);null" description:"广告描述"`
	Content     string    `orm:"column(content);null" description:"广告内容"`
	Url         string    `orm:"column(url);null" description:"广告链接"`
	Width       int       `orm:"column(width)" description:"广告宽度"`
	Height      int       `orm:"column(height)" description:"广告高度"`
	StartTime   time.Time `orm:"column(start_time);type(datetime);null" description:"开始时间"`
	EndTime     time.Time `orm:"column(end_time);type(datetime);null" description:"结束时间"`
	ViewNum     int       `orm:"column(view_num)" description:"点击率"`
	Status      int       `orm:"column(status)" description:"状态：1在用 2停用"`
	Sort        int       `orm:"column(sort)" description:"排序"`
	Note        string    `orm:"column(note);size(255);null" description:"备注"`
	CreateUser  int       `orm:"column(create_user);null" description:"添加人"`
	CreateTime  time.Time `orm:"column(create_time);type(datetime);null" description:"添加时间"`
	UpdateUser  int       `orm:"column(update_user);null" description:"更新人"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
	Mark        int       `orm:"column(mark)" description:"有效标识(1正常 0删除)"`
}

func (t *Ad) TableName() string {
	return "sys_ad"
}

func init() {
	orm.RegisterModel(new(Ad))
}

// 根据条件查询单条数据
func (t *Ad) Get() error {
	err := orm.NewOrm().QueryTable(new(Ad)).Filter("id", t.Id).One(t)
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
func (t *Ad) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Ad) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Ad) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}
