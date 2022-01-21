package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	fmt.Println("---------------------------原简历")
	resume := &Resume{
		name:       "王工作",
		sex:        "男",
		age:        10,
		company:    "光明顶无限责任公司",
		experience: "学武功和划水、摸鱼",
	}
	resume.display()
	fmt.Println("---------------------------简历特别好，抄")
	copyResume := resume.clone()
	copyResume.setPersonInfo("李工作", 21, "男")
	copyResume.display()
}
