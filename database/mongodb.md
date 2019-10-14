<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>MongoDB</font>
</div>
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    关于MongoDB的安装开发者可自行查找资料。(安装成功后不要忘记填写config中的MongoDB配置项)
    </font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    在db路径下新建类并继承\tiny\MongoDB，并告诉父类子类的表名，这样你就完成了一个MongoDB类的创建。
    </font>
</div>
     
    例如我这里新建一个用户表
    class User extends MongoDB {
        protected $table = 'User';       
    }
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    创建完成后我们便可以进行MongoDB操作，目前Tiny只提供了CURD的方法。因为非关系型数据库的特性，这四个方法也就能满足我们业务的需要。
    </font>
</div>

    //查询($filter为查询条件，$options为指定的查询字段)
    protected function find(array $filter, array $options = []): array {...}
    
    //新增($values为新增记录的值)
    protected function create(array $values): bool {...}
       
    //修改($filter为更新条件，$operate为具体操作)
    protected function update(array $filter, array $operate): bool {...}
    
    //删除($filter为删除条件)
    protected function delete(array $filter): bool {...}