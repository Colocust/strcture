<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>查询数据</font>
</div>

查询数据使用 `find` 方法
    
    User::find();

指定条件查询使用`where`方法
    
    User::where('age','=',18)->find();
        

    