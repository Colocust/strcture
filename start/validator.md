<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>参数验证</font>
</div>

<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    在APIRequest类以及APIResponse类中，我们需要注明每一个参数的规则。这样框架便可以正常的进行参数验证，业务逻辑也可以正常的使用该参数。
    </font>
</div>
<div align="left" style="margin-top:10px">
    <font face="Microsoft YaHei UI" size=3>
    打开之前创建的IndexRequest，假如目前我们需要的请求参数为id，它的数据类型是int并且为必要参数，那么我们可以这样定义:
    </font>
</div>

    namespace api;

    class IndexRequest extends Request {
        /**
         * @var int
         * @uses required
         */
        public $id;
    }
>###### var
    目前Tiny支持int,float,string,int[],float[],string[],*class,*class[],telephone,email,
    idNumber,file的规则。

>###### uses
    如果这不是一个必要参数，那么我们可以不加这行代码。示例如下：
    class IndexRequest extends Request {
        /**
         * @var int
         */
        public $id;
    }
<div align="left" style="margin-top:10px">
    <font face="Microsoft YaHei UI" size=3>
    如果开发者发现业务API出现了http_response_code等于415的错误，那么可以打开error.log错误日志，里面会有详细的错误原因。
    </font>
</div>

>###### 关于*class的使用方法
     场景:假设目前的请求参数是一个用户详情，它包含了id，name，age等字段。那么我们可以在请求类中这样定义：
     class GetUserInfoRequest extends Request {
        /**
         * @var UserInfoItem
         * @uses required
         */
        public $userInfo;
     }
     class UserInfoItem {
         /**
          * @var int
          * @uses required
          */
        public $id;
         /**
          * @var string
          * @uses required
          */
        public $name;
         /**
          * @var int
          * @uses required
          */
        public $age;
     }

>###### 关于*class[ ]的使用方法
     场景和*class一样，我们只需要在$userInfo字段的声明中加上数组的定义格式即可
     class GetUserInfoRequest extends Request {
        /**
         * @var UserInfoItem[]
         * @uses required
         */
        public $userInfo;
     }
     class UserInfoItem {
         /**
          * @var int
          * @uses required
          */
         public $id;
         /**
          * @var string
          * @uses required
          */
         public $name;
         /**
           * @var int
           * @uses required
           */
         public $age;
     }
     
>###### 建议
    目前Tiny仅支持GET、POST两种网络请求方式，而底层在接受这两种请求方式的数据时会使用$_GET、$_POST变量。这
    样就可能会导致接收到的数据格式和你传入的不一致(例如int变为string)，在参数验证时便会出现415的网络错误。
    所以作者建议前端在请求时将CONTENT_TYPE修改为application/json，这样Tiny就会用另一种方式去接收数据，从
    而避免以上错误。
    
>###### 注意
    @var以及@uses后面只能写规则，不能写规则以外的任何代码(包括注释)!