<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>API</font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    打开app目录下的API.php，我们会发现该抽象类中存在以下方法:
    </font>
</div>

    abstract public function requestClass(): Request; 
    
    abstract public function doRun(): Response;
    
    protected function beforeRun(\tiny\Request $request, \tiny\Response $response): bool {...}
    
    public function run(\tiny\Request $request, \tiny\Response $response) {...}
    
    protected function afterRun(\tiny\Request $request, \tiny\Response $response) {...}
    
    protected function needToken(): bool {...}
    
    protected function needValidate(): bool {...}
    
    private function checkToken(\tiny\Request $request): bool {...}
    
    public function getRequest(): Request {...}
    
    protected function getNet(): Net {...}
    
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    现在我们逐一讲解每个方法
    </font>
</div>

>###### requestClass
    获取业务API的APIRequest类，在参数验证中我们会通过这个方法来拿到APIRequesst的注释，从而分析出此API每个
    请求参数的数据类型。开发者可以不用关心框架是如何拿到验证规则的，只需要在每个业务API中重写这个抽象方法以
    及在APIRequest类中注明每个参数的类型即可。(注明方法会在后续APIRequest章节讲解)
    
>###### doRun
    执行业务API的逻辑并返回APIResponse
    
>###### beforeRun
    通过名字我们便可以知道这是在执行Run之前的方法。此方法主要是为了进行参数的验证，当参数验证未通过时便不会
    执行后面的逻辑，打开storage/error.log，我们可以查看详细的参数验证错误原因。
    
>###### run
    第一步我们会进行token的验证，验证通过后便会执行doRun方法，注意此doRun方法是业务API重写的doRun方法，也
    就是进入业务API的逻辑。
    
>###### afterRun
    目前afterRun方法没有任何逻辑，后续版本中该方法会进行返回参数的验证。开发者也可以根据自己业务的需要修改此
    方法。
    
>###### needToken
    判断业务API是否需要进行Token验证，如果不需要执行此操作，可在业务API中重写此方法并返回false。
    
>###### needValidate
    判断业务API是否需要进行参数验证，如果不需要执行此操作，可在业务API中重写此方法并返回false。
    
>###### checkToken
    校验Token的合法性，如果验证通过，API会创建一个网络层并将Token存储的信息放入网络层以供业务API使用。
        
>###### getRequest
    将请求参数赋值给APIRequest类中定义的公有成员变量。
    
>###### getNet
    业务API可调用此方法获取网络层中Token存储的信息。
                