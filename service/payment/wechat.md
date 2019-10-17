<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>微信支付</font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    在使用微信支付时，你需要提前阅读官方文档，这样便能更得心应手地使用它。
    </font>
</div>

>###### 微信支付使用方法

    namespace api;
    
    use service\payment\Factory;
    
    class WeChatPay extends API {
    
        public function requestClass(): Request {
            return new WeChatPayRequest();
        }
        
        public function doRun(): Response {
            $request = WeChatPayRequest::fromAPI($this);
            $response = new WeChatPayResponse();
            
            $params = array([
                'body' => '',
                'out_trade_no' => '',
                'total_fee' => ,    
                'notify_url' => '', // 支付结果通知网址，如果不设置则会使用配置里的默认地址
                'trade_type' => '', // 请对应换成你的支付方式对应的值类型
                'openid' => '', 
            ]);
            
            Factory::weChatPay()->pay($params);
            
            return $response;
        }
            
    }
    
>###### 正确的返回结果

     $result : {
         "return_code": "SUCCESS",
         "return_msg": "OK",
         "appid": "",
         "mch_id": "",
         "nonce_str": "",
         "openid": "",
         "sign": "",
         "result_code": "SUCCESS",
         "prepay_id": "",
         "trade_type": "JSAPI"
     }
     
>###### 注意

    你需要在config/wechat.php中填写payment的配置值