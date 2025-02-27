
terraform {
  required_providers {
    {{cookiecutter.provider_primary}} = {
      source  = "{{cookiecutter.provider_vendor}}/{{cookiecutter.provider_primary}}"
      version = "~> {{cookiecutter.provider_version}}"
    }
  }
}

provider "{{cookiecutter.provider_primary}}" {

}

module "simple_test" {
  source = "../../"
}

resource "null_resource" "test" {
}
