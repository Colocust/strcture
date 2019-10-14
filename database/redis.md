<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>Redis</font>
</div>
<div align="left" style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>
    关于redis的安装，可以参考<a style="font-weight: 500;color:#34495e;" href="https://www.jianshu.com/p/36e95b0a6350?utm_source=desktop&utm_medium=timeline">这篇文章</a>。(安装成功后不要忘记填写config中的redis配置项)
    </font>
</div>
<div style="margin-top:40px">
    <font face="Microsoft YaHei UI" size=3>你可以在Tiny中使用以下两种方式进行redis操作:</font>
</div>

    class A {
        public function redis() {
            $redis = new \tiny\Redis();
            $redis->redis()->functionName();
        }       
    }
    
    class B extends \tiny\Redis {
        public function redis() {            
            $this->redis()->functionName();
        } 
    }