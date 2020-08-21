package user

import (
	"awesomeProject1/GoMockDemo/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)    //1.mock 控制器
	defer ctl.Finish()    //5.进行mock 用例的期望值断言

	var id int64 = 1
	mockMale := mock.NewMockMale(ctl)  //2. 创建一个新的mock 实例
	gomock.InOrder(mockMale.EXPECT().Get(id).Return(nil))  //3. mock对象的行为注入

	user := NewUser(mockMale)   //4. mock 对象的注入：创建User 实例
	err := user.GetUserInfo(id)  // user.GetUserInfo(id) 调用的是事先模拟好的mock 方法
	if err != nil {
		t.Errorf("user.GetUserInfo err :%v",err)
	}
}
