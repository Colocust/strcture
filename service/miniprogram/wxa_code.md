<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>小程序码</font>
</div>

>###### 接口一

获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制。<a style="font-weight: 500;color:#34495e;" href="https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html">参考文档</a>
    
        MiniProgram::createQRCode([
            "path" => $path,
            "width" => 100
        ])
        
>###### 接口二

获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制。<a style="font-weight: 500;color:#34495e;" href="https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html">参考文档</a>

        MiniProgram::getWxaCode([
            "path" => $path,
            "width" => 100
        ])
                
>###### 接口三

获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。<a style="font-weight: 500;color:#34495e;" href="https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html">参考文档</a>
    
        MiniProgram::getWxaCodeUnlimit([
            'scene' => $scene,
            'page'  => $page,
            'width' => 100,
        ])