<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>文件上传</font>
</div>

>###### 使用方法

    namespace api;
    
    use service\Upload;
    
    class UploadAPI extends API {
    
        public function requestClass(): Request {
            return new UploadAPIRequest();
        }
        
        public function doRun(): Response {
            $request = UploadAPIRequest::fromAPI($this);
            $response = new UploadAPIResponse();
            //uploadFile方法会返回文件上传后的访问地址
            $response->fileUrl = Upload::uploadFile($request->file);
            return $response;
        }      
    }
<div></div>  

    namespace api;
    
    class UploadAPIRequest extends Request {
        /**
         * @var file 
         * @uses required
         */
        public $file;
    }
<div></div>    

    namespace api;
    
    class UploadAPIResponse extends Response {
        /**
         * @var string 
         */
        public $fileUrl;
    }
