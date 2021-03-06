<h1 align="center"> 类库[goarray] </h1>
<p align="center"> 数组切片基本操作方法集合。主要是配合数据库ORM常见的需求编写的。</p>

![](https://github.com/newgolibs/goarray/workflows/单元测试/badge.svg)

## 代码编写原则
- 遵循代码轮廓由后台配置生成,保证不会掺杂人为造成的错误
- 每一个特性都有对应的测试代码
- 向对象编程.每个类只实现一个invoke函数
- 向对象编程,参数全采用set/get操作

## 安装、引入代码 Installation

```
$ go get -u https://github.com/newgolibs/goarray
```


## 单页测试 - 也是使用案例

>[本类库所有单元测试案例](./test/)

### 读懂单元测试目录结构

1.测试入口文件

test\类、结构体名称\要测试的函数test.go<br>
`run` 函数接受参数为：要处理的内容，以及测试函数可能会需要的函数参数（面向对象写法决定：大部分函数是不接受参数）

2.测试数据源表

test\类、结构体名称\要测试的函数test\要测试的数据源Provider.go<br>
`DataProvider` 函数提供一个测试表，为2维数组 ：[]interface{}{...[]interface{}}<br>

3.数据源数组的结构有4个索引

* 要处理的输入
* 期望得到的结果
* 函数的参数
* 测试案例本身的说明

## License
MIT