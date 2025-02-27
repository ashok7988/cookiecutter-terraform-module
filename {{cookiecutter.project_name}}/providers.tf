terraform {
  
  required_version = ">= {{cookiecutter.minimum_terraform_version}}"
  required_providers {
    {{cookiecutter.provider_primary}} = {
      source  = "{{cookiecutter.provider_vendor}}/{{cookiecutter.provider_primary}}"
      version = ">= {{cookiecutter.provider_version}}"
    }
  }
}
