// gotest project doc.go

/*
gotest document
*/
package gotest

/*
单元测试文件，记住下面的这些原则：

文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码
你必须import testing这个包
所有的测试用例函数必须是Test开头
测试用例会按照源代码中写的顺序依次执行
测试函数TestXxx()的参数是testing.T，我们可以使用该类型来记录错误或者是测试状态
测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。
*/

/*
压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似,此处不再赘述，但需要注意以下几点：

压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母

func BenchmarkXXX(b *testing.B) { ... }
go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数

在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行
文件名也必须以_test.go结尾
*/
