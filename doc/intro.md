# 先说思想

1.为什么开头都叫Api？叫别的行不行？

受到Golang的限制，首字母大写挎包能调用到，那么前面几个关键字就很重要了，
所以在设计数据库层的时候，这里使用的是单例模式，每个对数据的调用都是API的形式，
所以开头叫Api，你也可以改成别的

2.为什么又蛇形又大小写？

可恶心到很多人，其实蛇形的目的是为了把动作分开，

例如通过“username字段找1个用户”，Java或者PHP开发喜欢写GetUser，Py开发喜欢get_user这也没问题，
很多朋友是大神喜欢直接写目的，不过后续2开让人接手就很吐了，或者过1年看自己代码就难受了，我一开始开发也是这样的，
后来才开始使用“过程”代替“目的”的命名方式


给个示例：
~~~
例如Api_find_byQqandPassword
~~~
1.这种命名方式请问他是数据库方法还是一个逻辑方法

2.他是增删改查中的哪一种？

3.他返回的是单条数据还是多条数据，我是使用map[string]interface{}来处理还是[]map[string]interface{}来处理？

4.这个方法主要需要哪两种数据类型才能完成查询？

好的如上问题请在看下面的代码前回答



·

·

·

·


```go
func Api_find_byQqandPassword(qq, password interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq":       qq,
		"password": password,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
```

好的，你在对照下Model代码，你看过程式命名方法是不是给后面接手你代码的人带来了很大的便利

写代码就像玩魂斗罗，要么一人跳太快给另一人给拖死，要么后面的太慢了把前面的拖死，如果你还没有决定项目规范，GorosePro的这个方案可以作为你的备选呢！
