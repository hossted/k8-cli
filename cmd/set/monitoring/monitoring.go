package monitoring

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// infoCmd represents the info command
var SetMonitoring = &cobra.Command{
	Use:   "monitoring",
	Short: `[m] hossted set monitoring true - Allow to send monitoring information about cpu, memory, network usage and logs to the hossted Dashboard.`,
	Long: ` [m] hossted set monitoring true - Allow to send monitoring information about cpu, memory, network usage and logs to the hossted Dashboard. 
	so it can be displayed within the hossted dashboard and recommend the course of action to secure your hossted application.`,
	Aliases: []string{"m"},
	Example: `
  hossted set monitoring true
  hossted set monitoring false
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}

		// Parse input
		var replica int32
		input := strings.ToLower(args[0])
		if input == "true" {
			replica = 1
		} else if input == "false" {
			replica = 0
		} else {
			fmt.Printf("\033[1;36m Only true/false is supported.\033[0m")
			fmt.Printf(" Input - %s\n\033[0m", input)
			os.Exit(0)
		}

		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}

		namespace, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
		if err != nil {
			panic(err.Error())
		}

		deployClient := clientset.AppsV1().Deployments(string(namespace))
		deploy, err := deployClient.Get(context.TODO(), "hossted-wrapper-develop-grafana-agent", metav1.GetOptions{})
		if err != nil {
			panic(err)
		}

		*deploy.Spec.Replicas = replica
		_, updateErr := deployClient.Update(context.TODO(), deploy, metav1.UpdateOptions{})
		if updateErr != nil {
			panic(updateErr)
		}

		fmt.Printf("Scaled deployment %s to %d replicas", "grafana-agent", *deploy.Spec.Replicas)

		fmt.Printf("\033[1;34m set monitoring \033[0m")
		fmt.Println(replica)
	},
}

func init() {
	//fmt.Printf("There are pods in the cluster\n")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
