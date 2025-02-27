import re
import sys

MODULE_REGEX = r"^terraform-\w*-\S*$"

module_name = "{{ cookiecutter.project_name }}"
provider_primary = "{{ cookiecutter.provider_primary }}"
provider_vendor = "{{ cookiecutter.provider_vendor }}"
provider_version = "{{ cookiecutter.provider_version }}"

# Collect all the errors in case there are multiple errors in the config.
# It's really annoying to get only one error, try again, and get an error
# that was already present.
errors = []

if not re.match(MODULE_REGEX, module_name):
    errors.append("ERROR: %s is not a valid Terraform module name!" % module_name)
    errors.append("Module names should follow the format `terraform-<PROVIDER>-<NAME>`")
    errors.append("See the Terraform Documentation for more details: https://developer.hashicorp.com/terraform/tutorials/modules/module#module-best-practices")

if not "." in provider_version:
    errors.append("ERROR: Provider version should include at least the MAJOR and MINOR versions.")
    errors.append("For example, `4.0` instead of `4`")

if not provider_primary:
    errors.append("ERROR: Primary terraform provider name is empty.")
    errors.append("See the Terraform Documentation for list of available providers: https://registry.terraform.io/browse/providers")

if len(errors):
    for error in errors:
        print(error)
    sys.exit(1)
