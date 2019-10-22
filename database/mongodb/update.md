<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>更新数据</font>
</div>
    
删除数据使用 `update` 方法
    
    User::set('age','=',18)->update($upsert,$multi);
    
在使用`update`方法时，你需要部分`链式操作`配合使用。`$upsert`这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
`$multi`默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。