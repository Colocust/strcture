<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>小程序登录</font>
</div>

>###### 流程

    1.调用wx.login()获取登录凭证code；
    2.调用wx.request()将code传到服务器；
    3.服务器端调用微信登录校验接口（appid+appsecret+code），获得session_key+openid；
    4.服务器端根据openid获取用户id(无则创之，有则取之)；
    5.通过用户id生成token，并存储tokenInfo；
    6.返回token，客户端将token存入本地缓存。
    
>###### 使用方法
          
    namespace api;
     
    use service\wechat\MiniProgram;
     
    class MiniProgramLogin extends API {
     
        public function requestClass(): Request {
            return new MiniProgramLoginRequest();
        }
         
        public function doRun(): Response {
            $request = MiniProgramLoginRequest::fromAPI($this);
            $response = new MiniProgramLoginResponse();
            
            $res = MiniProgram::getCode2Session($request->code);                      
            
            ***(数据库操作)
            
            ***(token操作)
                        
            $response->token = $token;
            
            return $response;
        }      
    }
<div></div>  

    namespace api;
    
    class MiniProgramLoginRequest extends Request {
        /**
         * @var string 
         * @uses required
         */
        public $code;
    }
<div></div>    

    namespace api;
    
    class MiniProgramLoginResponse extends Response {
        /**
         * @var string 
         * @uses required
         */
        public $token;
    }
    
>###### 注意

    你需要在config/wechat.php中填写mini_program的配置值