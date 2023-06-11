/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"k8s-baseline-scanner-v2/service"
	"os"

	"github.com/spf13/cobra"
)

var compareArg string

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "get difference between current baseline and standard baseline.",
	Long:  `get difference between current baseline and standard baseline.`,
	Run: func(cmd *cobra.Command, args []string) {
		if compareArg == "" {
			fmt.Println("pls use compare -h for help.")
			os.Exit(0)
		}
		compare()
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
	compareCmd.Flags().StringVar(&compareArg, "rce", "",
		"输入你想比对的资源名称,列表如下，并附有解释: \n"+
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

func compare() {
	switch compareArg {
	case "all":
		service.CompareAll()
	case "service":
		service.GetServiceCompareRes()
	case "secret":
		service.GetSecretCompareRes()
	case "pvc":
		service.GetPvcCompareRes()
	case "volume":
		service.GetvolumeCompareRes()
	case "cm":
		service.GetConfigmapCompareRes()
	case "replicas":
		service.GetReplicasCompareRes()
	case "affinity":
		service.GetAffinityCompareRes()
	case "container_resource":
		service.GetContainerResourceCompareRes()
	case "nodeSelector":
		service.GetNodeSelectorCompareRes()
	case "toleration":
		service.GetTolerationCompareRes()
	case "env":
		service.GetEnvCompareRes()
	case "image":
		service.GetImageCompareRes()
	case "probe":
		service.GetProbeCompareRes()
	case "volumeMount":
		service.GetVolumeMountCompareRes()
	default:
		fmt.Println("参数有误,请检查输入的参数.")
		os.Exit(1)
	}
}
