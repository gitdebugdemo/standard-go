#!/usr/bin/env bash

docker build -t  registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e689d00e6b1c1583914240 ./;

docker push registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e689d00e6b1c1583914240;

