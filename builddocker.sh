#!/usr/bin/env bash

docker build -t  registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e689d4a2cae71583914314 ./;

docker push registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e689d4a2cae71583914314;

