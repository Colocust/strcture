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
            //只需一行代码，返回值为文件的存储地址(完整的url地址需要开发者自行拼接)。
            $response->fileName = Upload::uploadFile($request->file);
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
        public $fileName;
    }
