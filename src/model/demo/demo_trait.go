package demo

//这个文件可以根据后台的配置文件生成
//生成各个属性的get/set
type Demo_trait struct {
	Name string
}

//根据配置生成的get方法
func (this *Demo_trait) GetName() string {
	return this.Name
}

//根据配置生成的set方法
func (this *Demo_trait) SetName(Name string) *Demo_trait {
	this.Name = Name
	return this
}
