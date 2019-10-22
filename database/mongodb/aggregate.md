<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>聚合</font>
</div>

MongoDB中聚合(aggregate)主要用于处理数据(诸如统计平均值,求和等)，并返回计算后的数据结果。

例如现在表中存在以下4条数据:
    
    {
        "_id" : "5daebc73e4520000ba004ed3","score" :32,"name":"tiny"
    }
    {
        "_id" : "5daebc76e4520000ba004ed4","score" :42,"name":"tiny"
    }
    {
        "_id" : "5daebc77e4520000ba004ed5","score" :52,"name":"tiny"
    }
    {
        "_id" : "5daed6ebe4520000ba004ed6","score" :52,"name":"locust"
    }
    
我需要通过 `name` 字段对数据进行分组，并计算 `name` 字段相同值的 `score` 总和。
    
    $pipeline = [
        [
            '$group' => [
                '_id' => '$name', 'scoreNum' => ['$sum' => '$score']
            ]
        ]
    ];
    $res = User::aggregate($pipeline);
    return $res;
    
返回值:
    
    {
        "_id": "locust",
        "scoreNum": 40
    },
    {
        "_id": "tiny",
        "scoreNum": 96
    }
    
如果需要计算score的平均值，你可以将`$sum`修改为`$avg`

    $pipeline = [
        [
            '$group' => [
                '_id' => '$name', 'scoreAvg' => ['$avg' => '$score']
            ]
        ]
    ];
    $res = User::aggregate($pipeline);
    return $res;

返回值:
    
    {
        "_id": "locust",
        "scoreAvg": 40
    },
    {
        "_id": "tiny",
        "scoreAvg": 32
    }
    
<div style="margin-top:50px"></div>

MongoDB提供了`$project`、`$match`、`$limit`、`$skip`、`$unwind`、`$group`、`$sort`、`$geoNear`操作符，开发者可在官方文档学习后自行进行实际操作。

