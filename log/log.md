<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>错误日志</font>
</div>

>###### 使用

    你可以调用Logger类中以下四种方法写入日志，内容可以是字符串，也可以是一个异常对象。
    Logger::getInstance()->info(string $message, ?\Throwable $throwable = null)
    Logger::getInstance()->warn(string $message, ?\Throwable $throwable = null)
    Logger::getInstance()->error(string $message, ?\Throwable $throwable = null)
    Logger::getInstance()->fatal(string $message, ?\Throwable $throwable = null)
    
>###### 查看

    打开终端并进入项目根目录，输入以下命令进行查看:
    tail -f storage/error.log
    
   