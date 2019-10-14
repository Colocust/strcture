<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>Token</font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    如果小伙伴不清楚token的涵义的话，可以参考<a style="font-weight: 500;color:#34495e;" href="https://my.oschina.net/jamesfancy/blog/1613994">这篇文章</a>。
    </font>
</div>

>###### 生成Token
    在我们的登录接口中，如果账号密码等都验证通过，那么我们就可以开始生成token。
    db路径下存在MyToken以及MyTokenInfo两个类，其中MyTokenInfo中存在newToken的静态方法可生成token。而在
    调用此方法时需要传入uid以及channel两个参数。uid是该登录用户记录在数据库中的主键id，而channel是用于区分
    此用户的登录渠道，开发者可根据自己的业务需要进行调整。
    
>###### 设置Token
    在生成token后我们便可以将此token作为key存入redis，而它的value便是MytokenInfo中的三个成员变量。
    
>###### 获取Token
    严格来说开发者是不需要自行获取token的，因为Tiny在run方法中进行了token的验证，并以该token创建了网络层。
    开发者只需要在业务API中调用网络层的方法便可以直接获取token中存储的信息。
    
>###### 完整的Token流程
    //生成token
    $token = MyTokenInfo::newToken($uid,$channel);
    //将需要存储的信息进行赋值
    $tokenInfo = new MyTokenInfo();
    $tokenInfo->token = $token;
    $tokenInfo->uid = $uid;
    $tokenInfo->channel = $channel;
    //设置token
    $myToken = new MyToken($token);
    $myToken->setToken($tokenInfo);   
    //获取token中存储的信息(此类必须继承/api/API)
    $uid = $this->getNet()->getUID();
    $channel = $this->getNet()->getChannel();
    
