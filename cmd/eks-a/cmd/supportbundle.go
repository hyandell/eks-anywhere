package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/logger"
	support "github.com/aws/eks-anywhere/pkg/support"
	"github.com/aws/eks-anywhere/pkg/validations"
	"github.com/aws/eks-anywhere/pkg/version"
)

type createSupportBundleOptions struct {
	fileName     string
	wConfig      string
	since        string
	sinceTime    string
	bundleConfig string
}

func (csbo *createSupportBundleOptions) kubeConfig(clusterName string) string {
	if csbo.wConfig == "" {
		return filepath.Join(clusterName, fmt.Sprintf(kubeconfigPattern, clusterName))
	}
	return csbo.wConfig
}

var csbo = &createSupportBundleOptions{}

var supportbundleCmd = &cobra.Command{
	Use:          "support-bundle -f my-cluster.yaml",
	Short:        "Generate a support bundle",
	Long:         "This command is used to create a support bundle to troubleshoot a cluster",
	PreRunE:      preRunSupportBundle,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := csbo.validate(cmd.Context()); err != nil {
			return err
		}
		if err := csbo.createBundle(csbo.since, csbo.sinceTime, csbo.bundleConfig); err != nil {
			return fmt.Errorf("failed to create support bundle: %v", err)
		}
		return nil
	},
}

func init() {
	generateCmd.AddCommand(supportbundleCmd)
	supportbundleCmd.Flags().StringVarP(&csbo.sinceTime, "since-time", "", "", "Collect pod logs after a specific datetime(RFC3339) like 2021-06-28T15:04:05Z")
	supportbundleCmd.Flags().StringVarP(&csbo.since, "since", "", "", "Collect pod logs in the latest duration like 5s, 2m, or 3h.")
	supportbundleCmd.Flags().StringVarP(&csbo.bundleConfig, "bundle-config", "", "", "Bundle Config file to use when generating support bundle")
	supportbundleCmd.Flags().StringVarP(&csbo.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	supportbundleCmd.Flags().StringVarP(&csbo.wConfig, "w-config", "w", "", "Kubeconfig file to use when creating support bundle for a workload cluster")
	err := supportbundleCmd.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
}

func (csbo *createSupportBundleOptions) validate(ctx context.Context) error {
	clusterConfig, err := commonValidation(ctx, csbo.fileName)
	if err != nil {
		return err
	}
	if !validations.KubeConfigExists(clusterConfig.Name, clusterConfig.Name, csbo.wConfig, kubeconfigPattern) {
		return fmt.Errorf("KubeConfig doesn't exists for cluster %s", clusterConfig.Name)
	}
	return nil
}

func preRunSupportBundle(cmd *cobra.Command, args []string) error {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		err := viper.BindPFlag(flag.Name, flag)
		if err != nil {
			log.Fatalf("Error initializing flags: %v", err)
		}
	})
	return nil
}

func (csbo *createSupportBundleOptions) createBundle(since, sinceTime, bundleConfig string) error {
	clusterSpec, err := cluster.NewSpec(csbo.fileName, version.Get())
	if err != nil {
		return fmt.Errorf("unable to get cluster config from file: %v", err)
	}
	os.Setenv("KUBECONFIG", csbo.kubeConfig(clusterSpec.Name))
	supportBundle, err := support.ParseBundleFromDoc(bundleConfig)
	if err != nil {
		return fmt.Errorf("failed to parse collector: %v", err)
	}

	var sinceTimeValue *time.Time
	sinceTimeValue, err = support.ParseTimeOptions(since, sinceTime)
	if err != nil {
		return fmt.Errorf("failed parse since time: %v", err)
	}

	archivePath, err := support.CollectBundleFromSpec(sinceTimeValue, &supportBundle.Spec)
	if err != nil {
		return fmt.Errorf("run collectors: %v", err)
	}

	logger.Info("\r \033[36mAnalyzing support bundle\033[m")
	err = support.AnalyzeBundle(&supportBundle.Spec, archivePath)
	if err != nil {
		return fmt.Errorf("there is an error when analyzing: %v", err)
	}

	logger.Info("a support bundle has been created in the current directory:", "path", archivePath)
	return nil
}