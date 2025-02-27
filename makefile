GITHUB_ORGANIZATION = "comcast-terraform-community-modules"
GITHUB_REPOSITORY = "cookiecutter-terraform-module"

.PHONY: tests
tests:
	@echo "Running tests by building a project using the cookiecutter template."
	cookiecutter --overwrite-if-exists --config-file tests/cookiecutter-input.yaml --no-input .
	rm -rf terraform-aws-test-cookiecutter # clean up the test project.
