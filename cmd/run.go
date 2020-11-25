// Copyright (c) Inlets Author(s) 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cmd

import (
	"context"
	"fmt"
	"os"
	"path"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the query",
	Long: `
  Run the query`,
	RunE: rundeploylister,
}

func rundeploylister(cmd *cobra.Command, args []string) error {

	kubeconfig := path.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return fmt.Errorf("Error building config: %v", err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("Error building kubernetes client: %v", err)
	}

	deployments, err := client.AppsV1().Deployments("").List(
		context.Background(), metav1.ListOptions{})

	if err != nil {
		return err
	}

	for _, deployment := range deployments.Items {
		for _, cs := range deployment.Spec.Template.Spec.Containers {
			fmt.Println(cs.Name, cs.Image)
		}
	}

	return nil
}
