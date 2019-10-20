<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>删除数据</font>
</div>

删除数据使用 `delete` 方法
    
    User::where('age','=',18)->delete();
    
如果不带任何条件调用`delete`方法会删除集合中所有的数据，谨慎使用！