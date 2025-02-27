import subprocess
import sys

def run_command(command):
    print(f"Running '{command}'")
    process = subprocess.Popen(command, stdout=subprocess.PIPE, shell=True, )
    for c in iter(lambda: process.stdout.read(1), b""):
        sys.stdout.buffer.write(c)

    returncode = process.wait()
    if returncode != 0:
        print(f"Failed to run command '{command}': {returncode}")
        sys.exit(returncode)

run_command('git init --initial-branch main')
run_command('make all')
run_command('make chores')

# Remove Artifactory GitHub Action if not needed
include_artifactory_gha = '{{ cookiecutter.publish_to_artifactory }}' == 'y'
if not include_artifactory_gha:
    run_command('rm .github/workflows/artifactory.yml')
