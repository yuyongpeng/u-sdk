package testu

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

////////////////////////////////////////////////////////
/////////////////   testify 模拟对象    /////////////////
///////////////////////////////////////////////////////

type MyInterface interface {
	Method() string
}

type MockObject struct {
	mock.Mock
}

func (m *MockObject) Method() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

// 等待测试的函数，需要传递一个对象(接口)，并且调用对象的方法
func MyFunction(obj MyInterface) string {
	return obj.Method()
}

func TestMyFunction(t *testing.T) {
	mockObj := new(MockObject)
	mockCall := mockObj.On("Method").Return("mocked data")
	mockCall2 := mockObj.On("DoSomething", 1).Return(true, nil) // 调用DoSomething方法，并返回一个布尔值和一个错误值)
	// 调用需要测试的方法
	result := MyFunction(mockObj)

	assert.Equal(t, "mocked data", result)
	// 清除后，可以再次测试别的变量
	mockCall.Unset()
	mockCall2.Unset()

	mockObj.On("Method").Return("mocked data")
	mockObj.On("DoSomething", 123).Return(true, nil) // 调用DoSomething方法，并返回一个布尔值和一个错误值)
	// 调用需要测试的方法
	result2 := MyFunction(mockObj)

	assert.Equal(t, "mocked data", result2)
	mockObj.AssertExpectations(t)
}
