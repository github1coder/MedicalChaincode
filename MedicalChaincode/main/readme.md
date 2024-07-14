## 环境配置与流程说明
1. 环境配置依照官方文档，也可参考csdn博客https://blog.csdn.net/sbsy123456/article/details/131294186以及原作者其他相关文档
2. 本项目代码的配置已经封装到文件config.sh中，该文件完成本项目测试时默认的网络身份配置和网络启动与链码部署，注意替换路径。文件附有注释说明。
3. 文件command.sh是测试时常用的测试命令，主要为 mimic数据集（https://mimic.mit.edu）的批量读入。
4. log.csv文件是当前测试时所采用的mimic数据集。

## 代码结构
./MedicalChaincode 文件夹中为智能合约代码
./application-gateway-go 文件夹中为后端代码

## 执行
```
cd path/to/back-end
source config.sh
source command.sh
```

## 结果
结果采用couchdb存储，进入这个网址，是couchdb的本地端口 

http://localhost:5984/_utils   Org1

http://localhost:7984/_utils   Org2

账户：admin

密码：adminpw

cloud文件夹中存放医疗数据读写的结果,可以保证当前操作的结果最新,但不保证所有文件最新; LOG文件记录事务

back-end默认设置为Org1操作,因此当添加私有数据后在5984端口查看可以看到私有数据集合,切换端口7984后就看不到了

如果是上云数据,查询摘要链可以看到url,也可以通过get函数直接下拉数据到本地文件,注意url的有效期默认为24h

可以记录事务,目前就是都写在一个LOG日志文件中,从上到下由新到旧


## 架构

![alt text](image.png)
