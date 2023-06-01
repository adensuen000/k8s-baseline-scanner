# k8s-baseline-scanner
## 功能简介
- k8s集群基线管理工具
- 扫描集群的配置信息并输出到文件。
- 拿当前集群的数据和标准基线数据比较并输出比较结果。

## 使用方式
### 1、编译和创建目录
- 拉取代码后，编译成二进制可执行文件
- 创建如下目录结构，并把config.ini放进config目录
> k8s-baseline-scanner
>  - bin
>  - config
>  - baseline-standard
>  - data

### 2、把config.ini放入config目录
### 3、执行二进制文件
- 进入到./bin目录下，然后使用./xxxx执行二进制文件
### 4、拉取的配置信息文件
- 从环境上拉取所有配置后，配置自动存入./data/baseline-current目录下，把所有生成的文件复制到baseline-standard当做标准基线，才能使用compare功能（也就是必须得先有标准基线，才能使用比对功能）。
### 5、config.ini说明
目前有两个参数
- namespace: namespace的值必须是逗号分隔
- kubeconfig: 是k8s集群认证文件的位置 
