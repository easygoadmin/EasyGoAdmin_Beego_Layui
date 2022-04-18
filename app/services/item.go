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
	"easygoadmin/app/constant"
	"easygoadmin/app/dto"
	"easygoadmin/app/models"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"time"
)

var Item = new(itemService)

type itemService struct{}

func (ctl *itemService) GetList(req dto.ItemPageReq) ([]vo.ItemInfoVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Item)).Filter("mark", 1)
	// 站点名称查询
	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}
	// 站点类型查询
	if req.Type > 0 {
		query = query.Filter("type", req.Type)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询总数
	count, _ := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.Item, 0)
	query.All(&lists)

	// 数据处理
	var result = make([]vo.ItemInfoVo, 0)
	for _, v := range lists {
		item := vo.ItemInfoVo{}
		item.Item = v
		// 站点类型
		typeName, ok := constant.ITEM_TYPE_LIST[v.Type]
		if ok {
			item.TypeName = typeName
		}
		// 站点图片
		if v.Image != "" {
			item.Image = utils.GetImageUrl(v.Image)
		}
		result = append(result, item)
	}
	return result, count, nil
}

func (ctl *itemService) Add(req dto.ItemAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.Item
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Status
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1

	// 图片处理
	if req.Image != "" {
		image, err := utils.SaveImage(req.Image, "item")
		if err != nil {
			return 0, err
		}
		entity.Image = image
	}

	// 插入数据
	return entity.Insert()
}

func (ctl *itemService) Update(req dto.ItemUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.Item{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}

	// 设置对象
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 图片处理
	if req.Image != "" {
		image, err := utils.SaveImage(req.Image, "item")
		if err != nil {
			return 0, err
		}
		entity.Image = image
	}

	// 更新记录
	return entity.Update()
}

func (ctl *itemService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Item{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.Item{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (ctl *itemService) Status(req dto.ItemStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.Item{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}
