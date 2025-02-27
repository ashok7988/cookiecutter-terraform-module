package terratest

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceCreationAndDeletion(t *testing.T) {

	// Create a terraform workspace with randomly generated id
	workspaceName := "basic_usage_" + random.UniqueId()

	// Allow the TerraformBinary option to be changed by users.
	// This makes it possible to switch to tofu easily.
	terraformBinary := os.Getenv("TF_BINARY")
	if len(terraformBinary) <= 0 {
		terraformBinary = "terraform"
	}

	tfOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// Point this at the specific module or example to test.
		TerraformDir: "../examples/basic_usage",

		// Switch between Terraform binaries.
		TerraformBinary: terraformBinary,

		Vars: map[string]interface{}{
			// Input Variables go here.
			// "test_input": testInput,
		},
	})

	// Create the workspace
	terraform.Init(t, tfOptions)
	terraform.RunTerraformCommand(t, tfOptions, "workspace", "new", workspaceName)

	// Validate if the workspace is created
	workspaceList, err := terraform.RunTerraformCommandAndGetStdoutE(t, tfOptions, "workspace", "list", "-no-color")
	if err != nil {
		t.Fatalf("Failed to get workspace list: %s", err)
	}
	assert.Contains(t, workspaceList, workspaceName)

	// Clean up the resources and workspace at the end of the test
	defer func() {
		terraform.Destroy(t, tfOptions)
		terraform.RunTerraformCommand(t, tfOptions, "workspace", "select", "default")
		terraform.RunTerraformCommand(t, tfOptions, "workspace", "delete", workspaceName)
	}()

	// Run Terraform init and apply to create resources within the workspace
	terraform.InitAndApply(t, tfOptions)

	//Validate the resources being created
	validate(t, tfOptions)
}

func validate(t *testing.T, opts *terraform.Options) {
	id := terraform.Output(t, opts, "null_resource_id")
	assert.NotEqual(t, "", id, "id is not null")
}
