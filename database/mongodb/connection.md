<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>连接MongoDB</font>
</div>

>###### 安装

    1、下载MongoDB(可自行前往官网下载并安装)。
    2、安装 MongoDB PHP 扩展。
    3、打开php.ini文件，添加mongodb配置。
    4、创建管理员账号。
    
>###### 连接
        
    打开config/db.php，填写mongodb项的配置参数。
    
    db_name  => 你需要连接的数据库(需自行创建)
    address  => 连接地址(默认为mongodb://127.0.0.1:27017)
    user     => 连接用户(需自行配置)
    password => 连接密码(需自行配置)
    
>###### 创建MongoDB类
      
    进入app/db路径，创建一个用户类，并继承tiny\MongoDB。
    
    class User extends tiny\MongoDB {
        //填写此类对应的集合(非常关键！)
        protected $table = 'user';
    }
    