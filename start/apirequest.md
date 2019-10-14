<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>APIRequest</font>
</div>
    
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    在APIRequest类中，我们需要注明每一个请求参数以及它的规则。这样框架便可以正常的进行请求参数验证，业务逻辑也可以正常的使用该请求参数。
    </font>
</div>
<div align="left" style="margin-top:10px">
    <font face="Microsoft YaHei UI" size=3>
    打开之前创建的IndexRequest，假如目前我们需要的请求参数为id，它的数据类型是int并且为必传参数，那么我们可以这样定义:
    </font>
</div>

    namespace api;
       
    class IndexRequest extends Request {
        /**
         * @var int
         * @uses required 
         */
        public $id;
    }
>###### var
    目前Tiny支持int,float,string,array,telephone,email,idNumber,file的规则验证(后续会不断完善)，
    每个参数必须注明它的规则。
    
>###### uses
    如果这不是一个必传参数，那么我们可以不加这行代码。示例如下：      
    class IndexRequest extends Request {
        /**
         * @var int
         */
        public $id;
    }
<div align="left" style="margin-top:10px">
    <font face="Microsoft YaHei UI" size=3>
    如果开发者发现业务API出现了http_response_code等于415的错误，那么可以打开error.log错误日志，里面会有详细的错误原因。
    </font>
</div>