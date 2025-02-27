{% set supported_providers = { 'aws': '0.24.1', 'google': '0.24.0', 'azurerm': '0.24.0' } %}

plugin "terraform" {
    enabled = true
}

{% if cookiecutter.provider_primary in ['aws', 'google', 'azurerm'] %}
plugin "{{ cookiecutter.provider_primary}}" {
    enabled = true
    version = "{{supported_providers[cookiecutter.provider_primary]}}"
    source  = "github.com/terraform-linters/tflint-ruleset-{{ cookiecutter.provider_primary }}"
}
{% endif %}
