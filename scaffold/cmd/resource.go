package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hashicorp/terraform-provider-aws/scaffold/resource"
)

var (
	snakeName string
	comments  bool
	name      string
)

var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Create scaffolding for a resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		return resource.Create(name, snakeName, comments)
	},
}

func init() {
	rootCmd.AddCommand(resourceCmd)
	resourceCmd.Flags().StringVarP(&snakeName, "snakename", "s", "", "If scaffold doesn't get it right, explicitly give name in snake case (e.g., db_vpc_instance)")
	resourceCmd.Flags().BoolVarP(&comments, "comments", "c", false, "Include instructional comments in source")
	resourceCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the entity")
}
