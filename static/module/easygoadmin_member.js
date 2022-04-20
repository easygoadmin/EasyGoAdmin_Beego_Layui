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
 * 会员管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func', 'form'], function () {
    var func = layui.func;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'Username', width: 120, title: '用户名', align: 'center'}
            , {field: 'Nickname', width: 120, title: '用户昵称', align: 'center'}
            , {
                field: 'Gender', width: 60, title: '性别', align: 'center', templet(d) {
                    var cls = "";
                    if (d.Gender == 1) {
                        // 男
                        cls = "layui-btn-normal";
                    } else if (d.Gender == 2) {
                        // 女
                        cls = "layui-btn-danger";
                    } else if (d.Gender == 3) {
                        // 保密
                        cls = "layui-btn-warm";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.GenderName + '</span>';
                }
            }
            , {field: 'Avatar', width: 90, title: '用户头像', align: 'center', templet: function (d) {
                    var avatarStr = "";
                    if (d.Avatar) {
                        avatarStr = '<a href="' + d.Avatar + '" target="_blank"><img src="' + d.Avatar + '" height="26" /></a>';
                    }
                    return avatarStr;
                }
            }
            , {
                field: 'Status', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '<input type="checkbox" name="status" value="' + d.Id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" ' + (d.Status == 1 ? 'checked' : '') + '>';
                }
            }
            , {field: 'CityName', width: 250, title: '所在地区', align: 'center'}
            , {field: 'Device', width: 100, title: '设备类型', align: 'center', templet(d) {
                    var cls = "";
                    if (d.Device == 1) {
                        // 苹果
                        cls = "layui-btn-normal";
                    } else if (d.Device == 2) {
                        // 安卓
                        cls = "layui-btn-danger";
                    } else if (d.Device == 3) {
                        // WAP站
                        cls = "layui-btn-warm";
                    } else if (d.Device == 4) {
                        // PC站
                        cls = "layui-btn-primary";
                    } else if (d.device == 5) {
                        // 微信小程序
                        cls = "layui-btn-disabled";
                    }

                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.DeviceName+'</span>';
                }}
            , {field: 'Source', width: 100, title: '用户来源', align: 'center', templet(d) {
                    var cls = "";
                    if (d.Source == 1) {
                        // 注册会员
                        cls = "layui-btn-normal";
                    } else if (d.Source == 2) {
                        // 马甲会员
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.SourceName+'</span>';
                }}
            , {field: 'CreateTime', width: 180, title: '注册时间', align: 'center',templet:'<div>{{ layui.util.toDateString(d.CreateTime*1000, "yyyy-MM-dd HH:mm:ss") }}</div>'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("会员用户");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
