package cmd

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/onspaceship/agent/pkg/config"
	"github.com/onspaceship/ship/pkg/client"
	"github.com/onspaceship/ship/pkg/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var setupDeliveryCmd = &cobra.Command{
	Use:   "setup-delivery DEPLOYMENT_NAME",
	Short: "Sets up delivery of an app to your current Kubernetes cluster",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("a deployment name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		deploymentName := args[0]

		ctx := context.Background()

		k8s := utils.NewKubernetesClient()
		namespace, _, err := utils.LoadKubeConfig().Namespace()
		if err != nil {
			log.Fatal("No namespace selected in kubeconfig")
		}

		deployment, err := k8s.AppsV1().Deployments(namespace).Get(ctx, deploymentName, metav1.GetOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				color.HiRed("Unable to find deployment %s", deploymentName)
				os.Exit(1)
			} else {
				log.Fatal(err)
			}
		}

		app, err := client.NewClient().GetApp(viper.GetString("current_team"), viper.GetString("current_app"))
		if err != nil {
			log.Fatal(err)
		}

		if _, exists := deployment.Labels[config.AppIdLabel]; exists {
			color.HiRed("This deployment already has a delivery label. If you wish to update it, please remove the label first.")
			os.Exit(1)
		}

		deployment.Labels[config.AppIdLabel] = app.ID

		_, err = k8s.AppsV1().Deployments(namespace).Update(ctx, deployment, metav1.UpdateOptions{})
		if err != nil {
			log.Fatal(err)
		}

		color.HiBlue("Delivery label added to deployment! The Spaceship Agent will use this deployment for future deliveries.")
	},
}

func init() {
	viper.SetDefault("registry_host", defaultRegistryHost)

	rootCmd.AddCommand(setupDeliveryCmd)
}
