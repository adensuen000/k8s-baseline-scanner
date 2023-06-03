# README.md
---

##  功能简介
*   k8s集群基线管理工具
*   扫描集群的配置信息并输出到文件。
*   拿当前集群的数据和标准基线数据比较并输出比较结果。

## 附加介绍
* 基线信息包含：statefulset、deployment、deamonset的如下信息：probe、env、image、containerresource、volume、volumeMount、nodeSelector、toleration、affinity、configmap、secret、service、pvc、replicas。
* 拉取的配置以json方式存入./data/baseline_current文件夹。
---
## 使用方式
- 下载解压后，进入./bin目录，以./xxxx方式执行二进制文件即可。
- 主要命令如下
> 查看子命令（主要为pull、compare）
```html
[root@10 bin]# ./k8s-baseline-scanner -h
scan current baseline and get the difference between them by compare them.

Usage:
  k8s-baseline-scanner [command]

Available Commands:
  compare     get difference between current baseline and standard baseline
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  pull        pull baseline of current cluster.

Flags:
      --config string   config file (default is $HOME/.k8s-baseline-scanner.yaml)
  -h, --help            help for k8s-baseline-scanner
  -t, --toggle          Help message for toggle

Use "k8s-baseline-scanner [command] --help" for more information about a command.
```
> 查看子命令支持的参数（如pull，conpare同理）

```
[root@10 bin]# ./k8s-baseline-scanner pull -h
pull baseline of current cluster.

Usage:
  k8s-baseline-scanner pull [flags]

Flags:
  -h, --help         help for pull
      --rce string   输入你想拉取的资源名称,如下: 
                     all
                     pvc
                     volume
                     cm
                     replicas
                     affinity
                     nodeSelector
                     toleration
                     env
                     image
                     probe
                     volumeMount

Global Flags:
      --config string   config file (default is $HOME/.k8s-baseline-scanner.yaml)
```
> 执行拉取资源，比如拉取pvc的资源（compare同理）

```
[root@10 bin]# ./k8s-baseline-scanner pull --rce pvc
拉取pvc成功. 请查看data目录.
```

> 拉取所有资源

```
[root@10 bin]# ./k8s-baseline-scanner pull --rce all
拉取image成功. 请查看data目录.
拉取volume-mount成功. 请查看data目录.
拉取probe成功. 请查看data目录.
拉取volume成功. 请查看data目录.
拉取configmap成功. 请查看data目录.
拉取toleration成功. 请查看data目录.
拉取node-selector成功. 请查看data目录.
拉取affinity成功. 请查看data目录.
拉取pvc成功. 请查看data目录.
拉取replicas成功. 请查看data目录.
拉取service成功. 请查看data目录.
拉取secret成功. 请查看data目录.
拉取env成功. 请查看data目录.
拉取container-resource成功. 请查看data目录.
```

## 初使用compare功能
 下载tar.gz包后，拉取某个环境的基线，复制一份到baseline-standard文件夹下当做标准基线，才可使用compare功能。

## 配置文件config.ini
 - namespace: 填写你想要扫描的命名空间。
 - kubeconfig: 填写k8s集群认证文件的路径。
