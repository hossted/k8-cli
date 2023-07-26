package info

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// infoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "information about k8",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

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
		// get the current namespace from the service account's token
		namespace, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
		if err != nil {
			panic(err.Error())
		}

		pods, err := clientset.CoreV1().Pods(string(namespace)).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the namespace\n", len(pods.Items))

		// Print a list of all the pods.
		fmt.Println("Pods in namespace", string(namespace))

		type Pod struct {
			Instance_id  string
			Docker_id    string
			Image_id     string
			Date_created time.Time
			ports        string
			status       string
			size         string
			name         string
			mounts       string
			networks     string
		}

		for _, pod := range pods.Items {
			p := Pod{
				name:         pod.Name,
				Instance_id:  string(pod.UID),
				Docker_id:    string(pod.UID),
				Image_id:     pod.Spec.Hostname,
				Date_created: pod.CreationTimestamp.Time,
				status:       string(pod.Status.Phase),
			}
			fmt.Println(p)
		}
		return nil
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
