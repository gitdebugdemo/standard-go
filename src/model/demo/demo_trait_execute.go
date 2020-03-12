package demo

//这个文件可以根据后台的配置文件生成:类壳
//实现的代码部分需要自己手写
type Demo_trait_execute struct {
	Demo_trait
}

func (this Demo_trait_execute) Say_world() string {
	return "hello,i am say.Name:" + this.Name
}