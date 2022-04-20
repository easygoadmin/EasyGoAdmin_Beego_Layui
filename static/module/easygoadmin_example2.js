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
 * 演示二管理
 * @author 半城风雨
 * @since 2022-04-15
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'Id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
        
             
			, {field: 'Name', width: 100, title: '演示名称', align: 'center'}
            
        
            
            , {field: 'Status', width: 100, title: '状态', align: 'center', templet: function (d) {
				return  '<input type="checkbox" name="Status" value="'+d.Id+'" lay-skin="switch" lay-text="正常|停用" lay-filter="Status" '+(d.Status==1 ? 'checked' : '')+'>';
            }}
            
        
             
			, {field: 'Sort', width: 100, title: '排序号', align: 'center'}
            
        
            , {field: 'CreateTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'UpdateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("演示二", 750, 450);

    
		
	
		
		//【设置状态】
        func.formSwitch('Status', null, function (data, res) {
            console.log("开关回调成功");
        });
		
	
		
	
    }
});
