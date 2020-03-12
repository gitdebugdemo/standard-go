package demo

//这个文件可以根据后台的配置文件生成:类壳
//实现的代码部分需要自己手写

/**
*
 */
type Demo struct {
	//复用类
	Demo_trait
	Demo_trait_execute
	Num int64
}

func (this *Demo) SetNum(num int64) *Demo {
	this.Num = num
	return this
}

func (this *Demo) GetNum() int64 {
	return this.Num
}

func (this Demo) Say_world() string {
	return this.Demo_trait_execute.Say_world()
}

func (this *Demo) Invoke() {
	this.Num = this.Num + 1
}
