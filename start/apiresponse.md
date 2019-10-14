<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>APIResponse</font>
</div>

<div align="left" style="margin-top:10px">
    <font face="Microsoft YaHei UI" size=3>
    开发者可以在此类中注明该API的返回参数，并在业务逻辑中进行赋值。尽管目前Tiny不支持Response的参数验证，但为了开发规范，作者还是建议注明每一个返回参数的规则。
    </font>
</div>
<div align="left" style="margin-top:10px">
    <font face="Microsoft YaHei UI" size=3>
    如果返回参数是一个二维数组，那么我们可以这样定义该参数：
    </font>
</div>
            
     class IndexResponse extends Response {
        /**
         * @var IndexResponseItem[]
         */
        public $result;
     }
<div></div>   
 
     class IndexResponseItem {
         /**
          * @var int
          */
         public $id;
         
         /**
          * @var string
          */
         public $name;
     }