package testu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// //////////////////////////////////////////////////////
// ///////////////   testify 辅助函数    /////////////////
// /////////////////////////////////////////////////////
type MyTestSuite struct {
	suite.Suite       // testify的suite 测试套件
	data        []int // 准备的测试数据
}

// 在测试套件运行之前设置测试环境。
func (suite *MyTestSuite) SetupSuite() {
	suite.data = []int{1, 2, 3}
	fmt.Println("SetupSuite")
}

// 在每个测试函数运行之前设置测试环境。
func (suite *MyTestSuite) SetupTest() {
	fmt.Println("SetupTest")
}

// 在每个测试函数运行之后清理测试环境。
func (suite *MyTestSuite) TearDownTest() {
	fmt.Println("TearDownTest")
}

// 在测试套件运行之后清理测试环境。
func (suite *MyTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

// 测试前需要执行的内容
func (suite *MyTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

// 测试后需要执行的内容
func (suite *MyTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

// 测试函数1
func (suite *MyTestSuite) TestSum() {
	sum := 0
	for _, num := range suite.data {
		sum += num
	}
	suite.Equal(6, sum)
}

// 测试函数2
func (suite *MyTestSuite) TestMod() {
	var ax interface{} = nil
	assert.Nil(suite.T(), ax) // suite 包中断言的方法
	assert.FileExists(suite.T(), "/Users/yuyongpeng/git/u-sdk/pkg/testu/example_test2.go")
	assertion := assert.New(suite.T())
	assertion.Equal(1, 1)
}

// 运行MyTestSuite下的所有测试用例
func TestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}
