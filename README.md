# Cookiecutter Terraform Module Template

This project is a [CookieCutter template](https://www.cookiecutter.io/) for creating Terraform Modules.

## Usage

Install CookieCutter. On MacOS this can be done with `brew install cookiecutter`.

Now run CookieCutter and specify this template as the source. Once launched you'll be asked a bunch of questions that'll be used to generate your project.

```bash
cookiecutter gh:comcast-terraform-community-modules/cookiecutter-terraform-module
```

## Development

### Helpful Commands
* To automatically format code and update documentation, run `make chores`. 
* To run formatting and linting checks as well as Terratests locally, run `make tests`.
* Pre-commit hooks are triggered with each commit. To bypass these hooks, add the `--no-verify` flag. Exercise caution when overriding these checks, as they are designed to enhance code quality and identify issues.

### CI with GitHub Actions

Enhance your infrastructure development process with integrated GitHub Actions Workflows, designed to streamline continuous integration (CI) for 
Terraform modules. These workflows facilitate testing and releasing Terraform modules, and are configured to 
execute against multiple versions of Terraform and OpenTofu, where applicable.

1. Up-to-Date README Check:
A preliminary step ensures your README documentation remains current, reflecting the latest changes and improvements.

2. Formatting Verification:
A formatting check ensures consistent and organized code, adhering to best practices.

3. Terraform Validation:
Automated Terraform validation ensures the integrity and accuracy of your infrastructure code, providing confidence in its correctness.

4. Terraform Lint Tests:
Terraform linting tests identify potential issues and ensure adherence to industry-established coding standards.

5. Checkov Security Tests:
Checkov is a powerful static code analysis tool designed to scan Infrastructure as Code (IaC) for security vulnerabilities for Terraform configurations.

6. Terratests Execution:
Terratests facilitate comprehensive testing and verification of your infrastructure code in real-world environments.

7. Publish to Artifactory:
Additionally, this pipeline is designed for seamless integration with Artifactory Terraform repositories, simplifying module publication as needed.

#### Set up for GitHub Actions

GitHub Actions is enabled by default after the initial project skeleton is checked into `main` branch. To run Terratests and Publish the module to Artifactory, a set of GitHub Actions variables and secrets must be configured. Refer the [GitHub Actions Documentation](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/store-information-in-variables) to learn more about setting variables and secrets.

##### Required Variables and Secrets

|Name| Type | Scope | Description |
|---|--|--|--|
| ARTIFACTORY_URL | `variable` | Organization | URL of the Artifactory instance where the module is published. Should be `https://artifactory.comcast.com`|
| ARTIFACTORY_REPOSITORY | `variable` | Organization | The name of the Artifactory Repository where the modules would be published. |
| ARTIFACTORY_API_KEY | `secret` | Organization | The API Key with write access to the Artifactory Repository where the module would be published. |
| ARTIFACTORY_MODULE_PATH | `variable` | Repository | The full module path, including name space in Artifactory where the module would be published. For example, if your namespace is `community`, and the module name is `terraform-aws-artifactory-upload`, the path would be `community/artifactory-upload/aws`. |
| Cloud Service Credentials | `secret` | Organization | You will need to make sure to add any other unique credentials that would be required by Terratests and examples to set up infrastructure and execute tests against them (eg. AWS credentials or VinylDNS credentials). |

**Note:** The `Scope` being organization is only a recommended best practice for where the variable or secret should be set. If there's a need to override a secret at the repository level, it can be done.

## References
* https://github.com/terraform-docs/terraform-docs
* https://developer.hashicorp.com/terraform/cli/commands/fmt
* https://developer.hashicorp.com/terraform/cli/commands/validate
* https://github.com/terraform-linters/tflint
* https://www.checkov.io/1.Welcome/What%20is%20Checkov.html
* https://github.com/gruntwork-io/terratest
* https://jfrog.com/help/r/jfrog-artifactory-documentation/terraform-repositories 
