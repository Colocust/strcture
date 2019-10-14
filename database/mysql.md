<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>MySQL</font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    目前Tiny对MySQL的支持还不完善，需要开发者自行完善sql语句并调用\tiny\MySQL中的CURD方法进行操作。
    </font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    在db路径下新建类并继承\tiny\MySQL,这样你就完成了一个MySQL类的创建。
    </font>
</div>

    namespace db;
    
    class MySQL extends \tiny\MySQL {...}
    
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    在\tiny\MySQL类中，我们可以发现四个基础的方法。
    </font>
</div>

    protected function select(string $sql): array {...}
    
    protected function create(string $sql): bool {...}
    
    protected function delete(string $sql): bool {...}
    
    protected function update(string $sql): bool {...}
   
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    开发者可自行完成sql语句并根据业务的需要调用这其中的一个方法。
    </font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    Tiny也提供了多库连接的方法，首先你需要在db的配置项中完成这个数据库的配置。
    </font>
</div>

    'mysql' => [
        //集群地址
        'host' => 'localhost',
    
        //端口号
        'port' => 3306,
    
        //数据库名
        'database' => 'test',
    
        //用户名
        'username' => 'root',
    
        //密码
        'password' => ''
      ],
      'mysql_wechat' => [
          //集群地址
          'host' => 'localhost',
      
          //端口号
          'port' => 3306,
      
          //数据库名
          'database' => 'test2',
      
          //用户名
          'username' => 'root2',
      
          //密码
          'password' => ''
        ]
        
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    其次，你需要在MySQL子类中注明此类的连接对象(默认连接对象是mysql)。
    </font>
</div>

    namespace db;
   
    class Mysql extends \tiny\MySQL {  
        //protected $connection = 'mysql'; 
        protected $connection = 'mysql_wechat';
    }
    
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    这样，父类就知道此时需要去连接配置为mysql_wechat的数据库了。
    </font>
</div>