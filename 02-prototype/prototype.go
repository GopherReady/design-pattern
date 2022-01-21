package prototype

import "fmt"

/**
 * @Description: 简历类，里面包含简历的基本信息
 */
type Resume struct {
	name       string
	age        int64
	sex        string
	company    string
	experience string
}

/**
 * @Description: 设置简历个人信息
 * @receiver r
 * @param name
 * @param age
 * @param sex
 */
func (r *Resume) setPersonInfo(name string, age int64, sex string) {
	r.name = name
	r.age = age
	r.sex = sex
}

/**
 * @Description: 设置工作经验
 * @receiver r
 * @param company
 * @param experience
 */
func (r *Resume) setWorkExperience(company string, experience string) {
	r.company = company
	r.experience = experience
}

/**
 * @Description: 显示简历内容
 * @receiver r
 */
func (r *Resume) display() {
	fmt.Printf("我的名字是%s，性别%s，今年%d岁，在%s工作，工作经验为:%s \n", r.name, r.sex, r.age, r.company, r.experience)
}

/**
 * @Description: 深拷贝，原型模式的核心
 * @receiver r
 * @return *Resume
 */
func (r *Resume) clone() *Resume {
	return &Resume{
		name:       r.name,
		sex:        r.sex,
		age:        r.age,
		company:    r.company,
		experience: r.experience,
	}
}
