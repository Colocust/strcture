<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>开始第一个API</font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    首先在项目的app/api目录下创建一个Index类，该类必须继承api/API，由于父类是一个抽象类，所以你必须重写它的抽象方法(后续章节会讲解这两个抽象方法)。
    </font>
</div>

> Index
 
    namespace api;
    
    class Index extends API {
        public function requestClass(): Request {
            return new IndexRequest();
        }
        
        public function doRun(): Response {
            $request = IndexRequest::fromAPI($this);
            $response = new IndexResponse();
            return $response;
        }
    }   
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    其次，你需要在目录下创建两个请求参数类，一个是IndexRequest，还有一个是IndexResponse。这两个类需要分别继承api/Request以及api/Response。
    </font>
</div>

> IndexRequest 
 
    namespace api;
     
    class IndexRequest extends Request {
        
    }

> IndexResponse
    
    namespace api;
     
    class IndexResponse extends Response {
        
    }
这样你就完成了第一个API的开发，此时使用Postman以POST的方式请求localhost/tiny/public/api/Index，你会发现以下返回信息:

    {
        "errMsg": "token not set or error",
        "code": 401
    }

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    这是因为在api/API类中是默认验证了token。如果我们此时的场景不需要验证token，那么可以在业务API中重写父类的needToken方法:
    </font>
</div>

    protected function needToken(): bool {
        return false;
    }

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    这样再请求一次URL地址，便可以正常返回Response
    </font>
</div>