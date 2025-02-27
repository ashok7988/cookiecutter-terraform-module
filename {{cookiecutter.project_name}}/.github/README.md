# CI/CD with GitHub Actions

Enhance your infrastructure development process with integrated GitHub Actions Workflows, designed to streamline continuous integration (CI) for 
Terraform modules. These workflows facilitate testing and releasing Terraform modules, and are configured to
execute against multiple versions of Terraform and OpenTofu, where applicable.

1. Up-to-Date README Check:
A preliminary step ensures your README documentation stays up-to-date, reflecting the latest changes and enhancements.

2. Formatting Verification:
A formatting check to maintain consistent and organized code, aligning with best practices.

3. Terraform Validation:
Terraform validation is automatically performed, guaranteeing the integrity and correctness of your infrastructure code.

4. Terraform Lint Tests:
Terraform linting tests to identify potential issues and adhere to industry-established coding standards.

5. Checkov Security Tests:
Checkov security tests identify vulnerabilities and ensure a robust infrastructure foundation.

6. Terratests Execution:
Terratests enable comprehensive testing and verification of your infrastructure code against real environments.

7. Publish to Artifactory:
Furthermore, this pipeline is structured to enable easy integration with Artifactory Terraform repositories, facilitating module 
publication when necessary.

## Set up for GitHub Actions

GitHub Actions is enabled by default after the initial project skeleton is checked into `main` branch. To run Terratests and Publish the module to Artifactory, a set of GitHub Actions variables and secrets must be configured. Refer the [GitHub Actions Documentation](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/store-information-in-variables) Learn more about setting variables and secrets.

### Required Variables and Secrets

|Name| Type | Scope | Description |
|---|--|--|--|
| ARTIFACTORY_URL | `variable` | Organization | URL of the Artifactory instance where the module is published. Should be `https://artifactory.comcast.com`|
| ARTIFACTORY_REPOSITORY | `variable` | Organization | The name of the Artifactory Repository where the modules would be published. |
| ARTIFACTORY_API_KEY | `secret` | Organization | The API Key with write access to the Artifactory Repository where the module would be published. |
| ARTIFACTORY_MODULE_PATH | `variable` | Repository | The full module path, including name space in Artifactory where the module would be published. For example, if your namespace is `community`, and the module name is `terraform-aws-artifactory-upload`, the path would be `community/artifactory-upload/aws`. |
| Cloud Service Credentials | `secret` | Organization | The credentials that would be required by Terratests and examples to set up infrastructure and execute tests against them. |

**Note:** The `Scope` being organization is only a recommended best practice for where the variable or secret should be set. If there's a need to override a secret at the repository level, it can be done.

## References
* https://github.com/terraform-docs/terraform-docs
* https://developer.hashicorp.com/terraform/cli/commands/fmt
* https://developer.hashicorp.com/terraform/cli/commands/validate
* https://github.com/terraform-linters/tflint
* https://www.checkov.io/1.Welcome/What%20is%20Checkov.html
* https://github.com/gruntwork-io/terratest
* https://jfrog.com/help/r/jfrog-artifactory-documentation/terraform-repositories 
