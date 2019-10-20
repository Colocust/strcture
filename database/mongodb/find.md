<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>查询数据</font>
</div>

查询数据使用 `find` 方法：
    
    //User为连接数据库章节创建的类
    User::find();



指定条件查询使用`where`方法：
    
    User::where('age','=',18)->find();
 
 
    
指定返回字段使用`field`方法：
    
    //_id会默认返回
    User::where('age','=',18)->field('age','name')->find();
    
        
指定返回条数使用`limit`方法：
    
    User::where('age','=',18)->field('age','name')->limit(5)->find();
    