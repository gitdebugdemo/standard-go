package demo

// 这个文件可以根据后台的配置文件生成
// 生成各个属性的get/set
type Demo_trait struct {
	Name string
}

// 根据配置生成的get方法
func (this *Demo_trait) GetName() string {
	return this.Name
}

// 根据配置生成的set方法
func (this *Demo_trait) SetName(Name string) *Demo_trait {
	this.Name = Name
	return this
}

func (this *Demo_trait) A流程1() string {
	return "[基类]Demo_trait-A流程1"
}

/**
通用的流程函数
这个时候,被流程函数调用的内部函数全部不能被重写,
要重写,基本全部函数都重写,再会
*/
func (this *Demo_trait) A我来定流程() string {
	return this.A流程1()
}
