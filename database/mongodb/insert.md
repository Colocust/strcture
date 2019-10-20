<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>添加数据</font>
</div>

使用`create`方法向数据库添加数据：
    
    User::create([
         '_id' => new ObjectId() . '',
         'age' => 22,
         'name' => 'tiny'
    ]);