<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>模板消息</font>
</div>

>###### 使用方法

    namespace api;
     
    use service\wechat\MiniProgram;
     
    class SendTemplateMessage extends API {
     
        public function requestClass(): Request {
            return new SendTemplateMessageRequest();
        }
         
        public function doRun(): Response {
            $request = SendTemplateMessageRequest::fromAPI($this);
            $response = new SendTemplateMessageResponse();           
            MiniProgram::sendTemplateMessage([
                'touser' => 'user-openid',
                'template_id' => 'template-id',
                'page' => 'index',
                'form_id' => 'form-id',
                'data' => [
                    'keyword1' => 'VALUE',
                    'keyword2' => 'VALUE2',
                    // ...
                ],
            ]);                                    
            return $response;
        }      
    }