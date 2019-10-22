<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>更新数据</font>
</div>
    
删除数据使用 `update` 方法
    
    User::set('age',18)->update($upsert,$multi);
    
在使用`update`方法时，你需要部分`链式操作`配合使用。`$upsert`这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
`$multi`默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。

<div style="margin-top:50px"></div>

`set`用来指定一个键的值。如果这个键不存在，则创建它。
    
    User::where('age','=',18)->set('name','tiny')->update();
    
<div style="margin-top:50px"></div>

`inc`用来增加已有键的值，或者在键不存在时创建一个键，只能用于整数、长整数或双精度浮点数。
    
    User::where('name','=','tiny')->inc('age',2)->update();

<div style="margin-top:50px"></div>

`unset`用来移除指定的键。

    User::where('name','=','tiny')->unset('age',1)->update();

<div style="margin-top:50px"></div>

`push`用来追加数组元素。若没有该列，则会自动增加。(不检查追加的元素是否重复)
    
    User::where('name','=','tiny')->push('sports','soccer')->update()

<div style="margin-top:50px"></div>
  
`addToSet`用来追加数组元素。若没有该列，则会自动增加。(检查追加的元素是否重复)
    
    User::where('name','=','tiny')->push('sports','soccer')->update()

<div style="margin-top:50px"></div>
    
通过`pop`可以从数组中移除数据。此时数据被看成是一个队列。
    
    //value=1队列开头移除 value=-1队列末尾移除
    User::where('name','=','tiny')->pop('sports',1)->update() 
 
<div style="margin-top:50px"></div>
    
`pull`可以根据元素值进行匹配，将匹配上的元素全部删除。

    User::where('name','=','tiny')->pull('sports','soccer')->update() 