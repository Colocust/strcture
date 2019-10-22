<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>查询数据</font>
</div>

查询数据使用 `find` 方法
    
    User::find();

`where`条件    

    User::where('age','=',18)->find();
    
`in`条件  
   
    //查询age在[18,20]区间的数据
    User::in('age',[18,19,20])->find();     
    
`or`条件
    
    //查询age!=18 or name='tiny'的数据
    User::or('age','!=',18)->or('name','=','tiny')->find();
    
指定返回条数使用`limit`方法
    
    User::where('age','<',18)->limit(5)->find();
    
指定返回的字段使用`field`方法

    User::where('age','>',18)->field('name')->find();
    
指定字段进行排序使用`sort`方法
        
    //1升序-1降序
    User::sort('age',1)->find();
    