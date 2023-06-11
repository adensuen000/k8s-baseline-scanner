/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-baseline-scanner-v2/base"
	"k8s-baseline-scanner-v2/service"
	"os"
)

var pullArg string

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pull baseline of current cluster.",
	Long:  `pull baseline of current cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		if pullArg == "" {
			fmt.Println("pls use pull -h for help.")
			os.Exit(0)
		}
		//初始化k8s客户端
		base.K8sInit.GetClientSet()
		pull()
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().StringVar(&pullArg, "rce", "",
		"输入你想拉取的资源名称,如下: \n"+
			"all\n"+
			"pvc\n"+
			"service\n"+
			"secret\n"+
			"volume\n"+
			"cm\n"+
			"replicas\n"+
			"affinity\n"+
			"nodeSelector\n"+
			"toleration\n"+
			"env\n"+
			"image\n"+
			"probe\n"+
			"volumeMount\n"+
			"container_resource\n")
}

func pull() {
	switch pullArg {
	case "all":
		service.PullAllResource()
	case "service":
		service.PullAllService()
	case "secret":
		service.PullAllSecret()
	case "pvc":
		service.PullAllPvc()
	case "volume":
		service.PullAllVolume()
	case "cm":
		service.PullAllConfigmap()
	case "replicas":
		service.PullAllReplica()
	case "affinity":
		service.PullAllAffinity()
	case "container_resource":
		service.PullAllContainerResources()
	case "nodeSelector":
		service.PullAllNodeSelector()
	case "toleration":
		service.PullAllToleration()
	case "env":
		service.PullAllEnv()
	case "image":
		service.PullAllImage()
	case "probe":
		service.PullAllProbe()
	case "volumeMount":
		service.PullAllVolumeMount()
	default:
		fmt.Println("参数有误,请检查输入的参数.")
		os.Exit(1)
	}
}
