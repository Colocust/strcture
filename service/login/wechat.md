<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>微信第三方登录</font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    使用之前，请你详细阅读<a style="font-weight: 500;color:#34495e;" href="https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html">官方文档</a>，并在开放平台注册开发者帐号，详细流程可参考<a style="font-weight: 500;color:#34495e;" href="https://www.cnblogs.com/yanbigfeg/p/9224756.html">这篇文档</a>。
    </font>
</div>

>###### 使用方法
    
    Factory::weChatLogin()->login($code);
    注意此Factory类的命名空间为service\login\Factory

>###### 注意
   
    使用之前，你需要在config/wechat.php文件中填写access_token以及third_party_login的配置值。
    
>###### 备注

    Subtle只会完成微信第三方登录服务，但service/login路径还提供了qq以及新浪微博三方登录类，开发者可根据自己
    的业务需要完善代码。
    