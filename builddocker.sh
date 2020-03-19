#!/usr/bin/env bash

docker build -t  registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e733496b028e1584608406 ./;

docker push registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e733496b028e1584608406;

