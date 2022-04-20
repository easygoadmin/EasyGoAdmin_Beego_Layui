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
 * 网站配置
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {
    var $ = layui.jquery;
    var form = layui.form;
    var func = layui.func;

    /**
     * 提交表单
     */
    form.on('submit(submitForm2)', function (data) {
        // 网络请求
        var loadIndex = layer.load(2);
        $.ajax({
            type: "POST",
            url: '/configweb/index',
            data: JSON.stringify(data.field),
            contentType: "application/json",
            dataType: "json",
            beforeSend: function () {
                // TODO...
            },
            success: function (res) {
                layer.close(loadIndex);
                if (res.code == 0) {
                    //0.5秒后关闭
                    layer.msg(res.msg, {icon: 1, time: 500});
                } else {
                    layer.msg(res.msg, {icon: 5});
                    return false;
                }
            },
            error: function () {
                layer.msg("AJAX请求异常");
            }
        });
        return false;
    });
});
