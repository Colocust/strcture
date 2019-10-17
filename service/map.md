<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>地图</font>
</div>

>###### 获取某个地址的经纬度坐标

在config/map.php中提供了3个获取经纬度的url地址，你可以使用其中的任意一个。我们这里以高德地图为例:
    
    $url = sprintf(config('map.amap.url.geocode_url'),$address,config('map.amap.key'));
    $res = Map::getLatAndLng($url);
    
>###### getDistance
    
    计算出两对经纬度坐标的距离。
    
>###### calcScope
    
    计算某个坐标指定距离的经纬度范围
    
    