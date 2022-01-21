## 单一职责原则（SRP）

表明一个类有且只有一个职责。一个类就像容器一样，它能添加任意数量的属性、方法等。然而，如果你试图让一个类实现太多，很快这个类就会变得笨重。任意小的改变都将导致这个单一类的变化。当你改了这个类，你将需要重新测试一遍。如果你遵守 SRP，你的类将变得简洁和灵活。每一个类将负责单一的问题、任务或者它关注的点，这种方式你只需要改变相应的类，只有这个类需要再次测试。SRP 核心是把整个问题分为小部分，并且每个小部分都将通过一个单独的类负责。

假设你在构建一个应用程序，其中有个模块是根据条件搜索顾客并以Excel形式导出。随着业务的发展，搜索条件会不断增加，导出数据的分类也会不断增加。如果此时将搜索与数据导出功能放在同一个类中，势必会变的笨重起来，即使是微小的改动，也可能影响其他功能。所以根据单一职责原则，一个类只有一个职责，故创建两个单独的类，分别处理搜索以及导出数据。

![img](https://images.cnblogs.com/cnblogs_com/OceanEyes/836627/o_srp.png)

## 单一职责应用

单一职责原则适用的范围有接口、方法、类。接口和方法必须保证单一职责，类就不必刻意保证，只要符合业务就行。

#### 场景设计

现在有一个场景, 需要修改用户的用户名和密码. 就针对这个功能我们可以有多种实现. 第一种:

```go
type UserOperate interface {
	UpdateUserInfo(job string)
}

type UserOperator struct {
	updateUserName string
	updatePassword string
}

func (u UserOperator) UpdateUserInfo(job string) {
	if job == u.updatePassword {
		// 修改密码逻辑
	} else if job == u.updateUserName {
		// 修改用户名逻辑
	}
}
```

第二种

```go
type UserInfo struct {
	name     string
	password string
}
type UserOperate interface {
	UpdateUserPassword(u *UserInfo)
	UpdateUsername(u *UserInfo)
}
}
func (UserInfo)UpdateUserPassword(u *UserInfo) {
	u.password = "kuri"
	fmt.Println(u.name)
}
func (UserInfo)UpdateUsername(u *UserInfo) {
	u.name = "codePassWord"
	fmt.Println(u.password)
}
```

这两种实现的区别: 第一种实现是根据操作类型进行区分, 不同类型执行不同的逻辑. 把修改用户名和修改密码这两件事耦合在一起了. 如果客户端在操作的时候传错了类型, 那么就会发生错误. 第二种实现是我们推荐的实现方式. 修改用户名和修改密码逻辑分开. 各自执行各自的职责, 互不干扰. 功能清晰明了.

由此可见, 第二种设计是符合单一职责原则的. 这是在方法层面实现单一职责原则.

#### 单一职责原则的应用

比如, 我们在网站首页可以注册, 登录, 微信登录.注册登录等操作. 我们通常的做法是:

```go
type UserInfo struct {
	name     string
	password string
}
type Operate interface {
	Login(userinfo UserInfo)
	Register(userinfo UserInfo)
	Logout(userinfo UserInfo)
}

type User struct {
}

func (u User) Login(userinfo UserInfo) {
	panic("implement me")
}

func (u User) Register(userinfo UserInfo) {
	panic("implement me")
}

func (u User) Logout(userinfo UserInfo) {
	panic("implement me")
}
```



