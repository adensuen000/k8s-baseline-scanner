/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-baseline-scanner/pkg"
	"k8s-baseline-scanner/pkg/service"
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
		pkg.K8sInit.GetClientSet()
		pull()
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().StringVar(&pullArg, "rce", "",
		"输入你想拉取的资源名称,如下: \n"+
			"all\n"+
			"pvc\n"+
			"volume\n"+
			"cm\n"+
			"replicas\n"+
			"affinity\n"+
			"nodeSelector\n"+
			"toleration\n"+
			"env\n"+
			"image\n"+
			"probe\n"+
			"volumeMount\n")

}

func pull() {
	switch pullArg {
	case "all":
		service.StartAllAction()
	case "pvc":
		service.Resource.GetAllPvc()
	case "volume":
		service.Resource.GetAllVolume()
	case "cm":
		service.Resource.GetAllConfigmap()
	case "replicas":
		service.Resource.GetAllReplicas()
	case "affinity":
		service.Resource.GetAllAffinity()
	case "nodeSelector":
		service.Resource.GetAllNodeSelector()
	case "toleration":
		service.Resource.GetAllToleration()
	case "env":
		service.Container.GetAllEnv()
	case "image":
		service.Container.GetAllImage()
	case "probe":
		service.Container.GetAllProbe()
	case "volumeMount":
		service.Container.GetAllVolumeMount()
	}
}
