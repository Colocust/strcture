<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>公共函数</font>
</div>

<div align="left" >
    <font face="Microsoft YaHei UI" size=3>
    公共函数位于common路径下，目前存在helper.php以及common.php两个文件。
    </font>
</div>

>###### helper.php
    function config(string $name = "") {...}                获取指定配置项的值
    function app($name) {...}                               获取容器中指定类的实例
    function curl_get(string $url, &$httpCode = 0) {...}    CURL GET请求
    function curl_post(string $url, $params) {...}          CURL POST请求
    function get_client_ip() {...}                          获取服务器ip地址
    
>###### common.php
    function millisecond() {...}                            获取毫秒级时间戳
    function millisecondToString() {...}                    获取时间格式,精确到毫秒级
    function toArray($obj) {...}                            对象转数组
    function toXML(array $array) {...}                      数组转化为XML格式的字符串     