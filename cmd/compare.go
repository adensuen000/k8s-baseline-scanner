/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-baseline-scanner/pkg/service"
	"os"
)

var rceCompare string

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "get difference between current baseline and standard baseline",
	Long:  `get difference between current baseline and standard baseline`,
	Run: func(cmd *cobra.Command, args []string) {
		if rceCompare == "" {
			fmt.Println("pls use compare -h for help.")
			os.Exit(0)
		}
		service.Compare(rceCompare)
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
	compareCmd.Flags().StringVar(&rceCompare, "rce", "",
		"输入你想比对的资源名称,列表如下，并附有解释: \n"+
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
