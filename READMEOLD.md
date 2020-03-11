# 通用代码结构
本分支是为了测试某个逻辑存在<br>
逻辑代码肯定放在test目录,并且下面没几个文件,最多2个文件<br>

入口文件为: CMD.sh<br>
各语言的标准测试<br>

* php vendor/bin/phpunit
* go /usr/local/go/bin/go run -mod=vendor /root/src/main.go &

# 分支备注
一份标准类面向对象代码结构实现. 包含入口函数invoke接口文件[生成]+trait包含属性的流程文件[生成]+逻辑文件[手写];

# 懒人一键验证效果

[Dockerfile源码](./Dockerfile)

docker run --rm  -it   registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e689c8f994631583914127;


# 原项目的说明

[README.md](./READMEOLD.md)