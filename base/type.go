package base

import (
	v1 "k8s.io/api/core/v1"
)

// 完整的容器信息
type ResourceContainer struct {
	ResourceInfo string         `yaml:"info"`
	Containers   []v1.Container `yaml:"containers,omitempty"`
}

// deployment、statefulset、daemonset 三种应用的基础属性
type DeployStsDsBase struct {
	ResourceInfo       string            `yaml:"resourceInfo"`
	Replicas           int32             `yaml:"replicas,omitempty"`
	Volumes            []v1.Volume       `yaml:"volumes,omitempty"`
	RestartPolicy      string            `yaml:"restartPolicy,omitempty"`
	DnsPolicy          string            `yaml:"dnsPolicy,omitempty"`
	NodeSelector       map[string]string `yaml:"nodeSelector,omitempty"`
	ServiceAccountName string            `yaml:"serviceAccountName,omitempty"`
	Affinity           *v1.Affinity      `yaml:"affinity,omitempty"`
	Tolerations        []v1.Toleration   `yaml:"tolerations,omitempty"`
}

// 数据卷
type ResourceVolume []string
type AllVolume map[string]ResourceVolume

// 镜像
type ResourceImage map[string]string
type AllImage map[string]ResourceImage

// 副本数
type AllReplica map[string]string

// 容器对应的资源规格
type AllContainerResources map[string]RContainerResource

// 单个应用（deploy、sts、ds）对应的容器资源
type RContainerResource struct {
	RContainerResource []ContainerResource `yaml:"ContainerResource"`
}

// 单个容器资源规格
type ContainerResource struct {
	ContainerName string   `yaml:"containerName"`
	ReqLim        string   `yaml:"reqLim,omitempty"`
	Claims        []string `yaml:"claims,omitempty"`
}

// pvc(按照命名空间统计)
type AllPvc map[string]Pvc

// 单个应用（deploy、sts、ds）对应的pvc, 统计Name,Request和limit
type Pvc struct {
	PvcInfo []string `yaml:"pvcInfo"`
}

// 所有configmap
type AllConfigmap map[string]ResourceConfigmap

// 单个应用（deploy、sts、ds）对应的configmap
type ResourceConfigmap struct {
	ResourceConfigmap []Configmap `yaml:"resourceConfigmap,omitempty"`
}

// 单个configmap
type Configmap struct {
	CmName string            `yaml:"cmName,omitempty"`
	CmData map[string]string `yaml:"cmData,omitempty"`
}

// 亲和、反亲和
type AllAffinity map[string]*v1.Affinity

// 节点选择器
type AllNodeSelector map[string]ResourceNodeSelector
type ResourceNodeSelector map[string]string

// 所有应用的容忍(key,value,effect,TolerationSeconds)
type AllToleration map[string]ResourceToleration
type ResourceToleration []string

// 所有env
type AllEnv map[string]ResourceEnv

// 单个应用（deploy、sts、ds）对应的环境变量
type ResourceEnv []ContainerEnv

// 单个容器的环境变量
type ContainerEnv struct {
	ContainerName string             `yaml:"containerName"`
	Env           []v1.EnvVar        `yaml:"env,omitempty"`
	EnvFrom       []v1.EnvFromSource `yaml:"envFrom,omitempty"`
}

// 所有探针
type AllProbe map[string]ResourceProbe

// 单个应用包含的探针
type ResourceProbe []ContainerProbe

// 单个容器包含的探针
type ContainerProbe struct {
	ContainerName  string    `yaml:"containerName"`
	LivenessProbe  ProbeRule `yaml:"livenessProbe,omitempty"`
	ReadinessProbe ProbeRule `yaml:"redinessProbe,omitempty"`
	StartUpProbe   ProbeRule `yaml:"startUpProbe,omitempty"`
}

// 单个探针的内容
type ProbeRule struct {
	HTTPGet       string `yaml:"httpGet,omitempty"`
	TcpSocket     string `yaml:"tcpSocket,omitempty"`
	Command       string `yaml:"command,omitempty"`
	ProbeStretegy string `yaml:"probeStretegy,omitempty"`
}

// 存储卷挂载
type AllVolumeMount map[string]ResourceVolumeMount
type ResourceVolumeMount []ContainerVolumeMount
type ContainerVolumeMount map[string][]VolumeMount
type VolumeMount v1.VolumeMount

// 所有service
type AllService map[string][]Service

// 单个service
type Service struct {
	ServiceInfo string            `yaml:"serviceInfo"`
	Selector    map[string]string `yaml:"selector,omitempty"`
	Potrs       []string          `yaml:"potrs,omitempty"`
}

// 所有secret
type AllSecret map[string]ResourceSecret

// 单个应用（deploy、sts、ds）对应的secret;  列表为secret的name列表
type ResourceSecret []string
