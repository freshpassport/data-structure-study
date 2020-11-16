## 冒泡排序算法 ##

**测试结果**

* 使用int类型

```
$ go test -v -test.run=TestSortInt
=== RUN   TestSortInt
[15 11 10 7 3 -1]
--- PASS: TestSortInt (0.00s)
PASS
ok  	sorting/bubble	0.005s
```

* 使用自定义类型

> 需要定义一个自定义类型的Slice并且：
> 实现Len(),Less(),Swap()三个函数

```
$ go test -v -test.run=TestSortUser
=== RUN   TestSortUser
[@(Name: 赵武 RegTime: 1605548113) @(Name: 李四 RegTime: 1600732800) @(Name: 张三 RegTime: 1592179200)]
--- PASS: TestSortUser (0.00s)
PASS
ok  	sorting/bubble	0.004s
```
