<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>内置业务</font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    在Service路径下存在着部分我自己已经封装好了的业务，下面我们将逐一讲解。
    </font>
</div>

>###### payment
    payment下存在着三个类。分别是微信支付、支付宝支付以及一个工厂类(在使用第三方支付前请详细阅读第三方支付
    官方文档)。目前微信支付已经完成，使用方式如下：   
    use service\payment\Factory;
    $params = array([
        'body' => '',
        'out_trade_no' => '',
        'total_fee' => ,    
        'notify_url' => '', // 支付结果通知网址，如果不设置则会使用配置里的默认地址
        'trade_type' => '', // 请对应换成你的支付方式对应的值类型
        'openid' => '', 
    ]);
    Factory::weChatPay()->pay($params);
    
>###### wechat
    wechat目前存在AccessToken以及MiniProgram类，而MiniProgram类中封装了常用的小程序接口，例如：发送模板
    消息，生成小程序二维码等。使用它们之前请在config中配置好相关参数并详细阅读官方文档。    

>###### Map
    public static function getLatAndLng(string $url) {...}                  //获取指定地址的经纬度，完整的url需要开发者自行拼接
    public static function getDistance($lat1, $lng1, $lat2, $lng2) {...}    //通过两对经纬度获取它们之间的距离
    public static function calcScope($lat, $lng, $radius) {...}             //指定经纬度，计算出以它为中心点，半径为$radius的圆的经纬度范围
    
>###### Upload
    public static function uploadFile(object $file) {...} //文件上传，$file为请求参数中是文件格式的参数