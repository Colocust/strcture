<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>全局配置</font>
</div>

<div>
    <font face="Microsoft YaHei UI" size=3>目前Tiny中存在5个配置文件，开发者只需要关注app.php以及db.php，其余三个配置文件与具体业务有关。</font>
</div>

>###### app.php
    app_debug                是否开启调试模式。On开启，Off关闭。
    log_file                 错误日志的默认路径，开发者可自行修改。
    upload_file_folder       文件上传后的默认路径，开发者可自行修改。
 
>###### db.php
    mongodb                  mongodb的配置参数
    redis                    redis的配置参数
    mysql                    mysql数据库的配置参数