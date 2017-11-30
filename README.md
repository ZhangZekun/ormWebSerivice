# xorm VS database/sql 

## 编程效率比较

xorm只需要连接好数据库，设置结构体和数据库中表的映射关系，就可以很方便地对数据库进行操作。想插入数据的时候可以直接Engine.Insert()。函数会根据映射关系将参数插入到数据库的特定表中。编程效率比较高。

## 程序结构比较

database/sql 的比较优雅，分层结构。但可能是因为我刚接触xorm，也许它也有更优雅的写法呢？但就目前我接触到的知识来看，还是database/sql的结构更加清晰，容易调试、增加功能。

## 服务性能

我用ab -n1000 -c100 http://localhost:8080/service/userinfo?userid=1和ab -n5000 -c500 http://localhost:8080/service/userinfo?userid=1去测试性能，发现就整体而言，xorm的性能更差一点，每条请求的处理时间较长。当请求总数增加时，差距更加明显。

## 问题思考

* orm 是否就是实现了 dao 的自动化？

  暂时还没有研究orm的源码，但我感觉是的。

* 使用 ab 测试性能

  使用ab -n1000 -c100 http://localhost:8080/service/userinfo?userid=1：

  database/sql:

  ![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw5_ormWeb/1000_100.png)

  orm:

![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw5_ormWeb/1000_100_orm.png)



​	使用ab -n5000 -c500 http://localhost:8080/service/userinfo?userid=1：

​	database/sql

![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw5_ormWeb/5000_500.png)

​	orm:

![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw5_ormWeb/5000_500_orm.png)