<div align="center" style="height:50px">
    <font face="Microsoft YaHei UI" size=5>链式操作</font>
</div>

>###### where
   
    User::where('age','>',18)
        ->where('name','=','tiny')
        ->find();
     
<div style="margin-top:50px"></div>   

>###### in

    User::in('age',[18,19])
        ->find();
        
<div style="margin-top:50px"></div>   
    
>###### or
    
    //查询age=18 or name='tiny'的数据
    User::or('age','=',18)
        ->or('name','=','tiny')
        ->find();
    
  
<div style="margin-top:50px"></div>
    
>###### limit
    
    //指定返回满足条件的5条数据
    User::limit(5)
        ->find();
    
    
<div style="margin-top:50px"></div> 
 
>###### field
    
    //指定返回的字段
    User::field('_id','name')
        ->find();
 
<div style="margin-top:50px"></div>    
    
>###### sort
    
    //指定字段进行排序(1升序-1降序)
    User::sort('age',1)
        ->find();
     