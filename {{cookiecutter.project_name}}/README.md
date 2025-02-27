# {{cookiecutter.project_name}}

## Description

A curated Terraform module for creating a(n) {{ (" ").join(cookiecutter.project_name.split('-')[1:]) }}.

## Features

This {{cookiecutter.project_name}} module supports the following:

-
-
-

## Usage

#### Configuration

To integrate the module into your Terraform configuration, add a module block specifying the registry path for the module along with its semantic version to be used:

```hcl
module "test" {
  source  = "artifactory.comcast.com/terraform-community-module__community/{{ ("-").join(cookiecutter.project_name.split('-')[2:]) }}/{{ cookiecutter.project_name.split('-')[1] }}"
  version = "~> 1.0"
}
```

#### Example

See `examples/basic_usage` for an example of this module in use.

#### Best Practices

As always, please consult Comcast's resources like OneCloud to make sure you are implementing our best practices. Some relevant things to remember are:

- The [Terraform Community Modules Best Practices](https://github.com/comcast-terraform-community-modules/best-practices/)
- The [naming strategy for IAM policies](https://onecloud.comcast.net/docs/using-aws/account-access/iam-policies/)
- The resource [tagging strategy](https://onecloud.comcast.net/docs/using-aws/tagging/)

<!-- BEGIN_TF_DOCS -->
<!-- END_TF_DOCS -->

## Development

- To automatically format code and update documentation, run `make chores`.
- To run formatting and linting checks as well as Terratests locally, run `make tests`.
- Pre-commit hooks are triggered with each commit. To bypass these hooks, add the `--no-verify` flag. Exercise caution when overriding these checks, as they are designed to enhance code quality and identify issues.

## CI with GitHub Actions

Continuous Integration for this module is configured using GitHub Actions. For more details on the setup, refer to the [GitHub Actions setup guide](./.github/README.md).
