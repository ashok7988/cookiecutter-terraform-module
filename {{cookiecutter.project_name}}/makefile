# Select the engine: opentofu or terraform.
# make validate TF_ENGINE=opentofu
# This will switch the engine for all actions, including testing.
TF_ENGINE:=terraform
ifeq ($(TF_ENGINE), terraform)
	TF_BINARY:=terraform
else ifeq ($(TF_ENGINE), opentofu)
	TF_BINARY:=tofu
endif

TERRAFORM_DOCS_VERSION := v0.16.0
GO_VERSION := 1.20.6
GITHUB_ORGANIZATION := {{ cookiecutter.github_org }}
GITHUB_REPOSITORY := {{ cookiecutter.project_name }}

# Detect whether we're on Linux or MacOS and alter install accordingly.
# @TODO: Windows support

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
		INSTALL_SYSTEM += install_linux
		SYSTEM=linux
endif
ifeq ($(UNAME_S),Darwin)
		INSTALL_SYSTEM += install_macos
		SYSTEM=darwin
endif

# Detect Architecture- amd64 or arm/arm64
UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),x86_64)
		ARCHITECTURE:=amd64
else ifeq ($(UNAME_M),aarch64)
	ARCHITECTURE := arm64
else
		ARCHITECTURE:=$(UNAME_M)
endif

# Assign sudo if installed, else use empty string
SUDO := $(shell which sudo)

#
# Installation
#

.PHONY: all
all: install configure

.PHONY: install
install: $(INSTALL_SYSTEM)

.PHONY: install_macos
install_macos: install_macos_terraform_docs install_macos_tflint install_macos_tenv install_macos_go install_macos_checkov install_macos_pre-commit

.PHONY: install_macos_terraform_docs
install_macos_terraform_docs:
	brew install terraform-docs

.PHONY: install_macos_tflint
install_macos_tflint:
	brew install tflint

.PHONY: install_macos_tenv
install_macos_tenv:
	brew install tenv

.PHONY: install_macos_go
install_macos_go:
	brew install go

.PHONY: install_macos_checkov
install_macos_checkov:
	brew install checkov

.PHONY: install_macos_pre-commit
install_macos_pre-commit:
	brew install pre-commit

.PHONY: install_linux
install_linux: install_linux_terraform_docs install_linux_tflint install_linux_checkov install_linux_pre-commit

.PHONY: install_linux_terraform_docs
install_linux_terraform_docs: /usr/local/bin/terraform-docs

/usr/local/bin/terraform-docs:
	curl -sSLo ./terraform-docs.tar.gz https://github.com/terraform-docs/terraform-docs/releases/download/$(TERRAFORM_DOCS_VERSION)/terraform-docs-$(TERRAFORM_DOCS_VERSION)-$(SYSTEM)-$(ARCHITECTURE).tar.gz
	tar -xzf terraform-docs.tar.gz
	chmod +x terraform-docs
	${SUDO} mv terraform-docs /usr/local/bin/terraform-docs

.PHONY: install_linux_tflint
install_linux_tflint:
	curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash

.PHONY: install_linux_go
install_linux_go:
	curl -sSLo ./go-binary.tar.gz https://go.dev/dl/go$(GO_VERSION).$(SYSTEM)-$(ARCHITECTURE).tar.gz
	rm -rf /usr/local/go && tar -C /usr/local -xzf go-binary.tar.gz
	export PATH=$PATH:/usr/local/go/bin
	cd terratest && go mod init github.com/{{ cookiecutter.github_org }}/{{ cookiecutter.project_name }}/tests

.PHONY: install_linux_checkov
install_linux_checkov:
	python -m pip install checkov

.PHONY: install_linux_pre-commit
install_linux_pre-commit:
	${SUDO} apt-get update
	${SUDO} apt-get -y install pre-commit

.PHONY: configure
configure: configure_pre-commit_hooks

.PHONY: configure_pre-commit_hooks
configure_pre-commit_hooks:
	pre-commit install

#
# Actions
#

.PHONY: chores
chores: formatting documentation_build

.PHONY: formatting
formatting:
	$(TF_BINARY) fmt

.PHONY: documentation_build
documentation_build:
	terraform-docs .

.PHONY: linting_fix
linting_fix:
	tflint --init
	tflint --fix

.PHONY: publish
publish:
	@echo "Publishing not currently implemented."


#
# Tests
#

.PHONY: tests
tests: validation terratest formatting_test linting_test documentation_test security_test custom_policy_check

.PHONY: validation
validation:
	$(TF_BINARY) init -backend=false
	$(TF_BINARY) validate
	rm -Rf .terraform

terratest/go.mod:
	cd terratest && \
	go mod init github.com/{{ cookiecutter.github_org }}/{{ cookiecutter.project_name }}/tests

.PHONY: terratest
terratest: terratest/go.mod
	cd terratest && \
	go mod tidy && \
	go test -timeout 60m

.PHONY: formatting_test
formatting_test:
	$(TF_BINARY) fmt -check

.PHONY: linting_test
linting_test:
	tflint --init
	tflint

# Please see .checkov.yml for optional parameters and update it as needed
# Run module/project specific checks
# Please add any custom checks under checkov_custom_checks folder in the project root and update checkov.yml to skip any checks
# Display only failed checks
.PHONY: security_test
security_test:
	checkov --quiet --framework terraform

# Run custom checks from an external git repository https://github.com/comcast-terraform-community-modules/checkov-custom-policies.git
# Checks if the resources have OneCloud recommended tags
# Display only failed checks
.PHONY: custom_policy_check
custom_policy_check:
	checkov --external-checks-git git@github.com:comcast-terraform-community-modules/checkov-custom-policies.git --quiet

.PHONY: documentation_test
documentation_test:
	terraform-docs . --output-check
