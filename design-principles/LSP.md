## 里氏替换原则（LSP）

里氏替换原则指出，派生的子类应该是可替换基类的，也就是说任何基类可以出现的地方，子类一定可以出现。值得注意的是，当你通过继承实现多态行为时，如果派生类没有遵守LSP，可能会让系统引发异常。所以请谨慎使用继承，只有确定是“is-a”的关系时才使用继承。

假设你在开发一个大的门户网站，并提供很多定制的功能给终端用户，根据用户的级别，系统提供了不同级别的设定。考虑到这个需求，设计如下类图：

![img](https://images.cnblogs.com/cnblogs_com/OceanEyes/836627/o_ISettings.png)

可以看到，ISettings 接口有 GlobalSettings、SectionSettings 以及 UserSettings 三个不同的实现。GlobalSettings 设置会影响整个应用程序,例如标题、主题等。SectionSettings 适用于门户的各个部分,如新闻、天气、体育等设置。UserSettings 为特定登录用户设置,如电子邮件和通知偏好。

