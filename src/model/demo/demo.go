package demo

//这个文件可以根据后台的配置文件生成:类壳
//实现的代码部分需要自己手写

/**
*
 */
type Demo struct {
	//接口
	Demo_interface
	//复用类
	Demo_trait_execute
	Num int64
}

func (this Demo) Say_world() string {
	return this.Demo_trait_execute.Say_world()
}

func (this *Demo) Invoke() {
	this.Num = this.Num + 1
}
