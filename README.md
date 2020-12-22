# gatekeeper-cli

This is an example repo implementing the [gatekeeper](https://github.com/agrim123/gatekeeper) as a cli application.

## Usage

```bash
$ gatekeeper
Gatekeeper is an authentication and authorization oriented deployment and access managment tool.

Usage:
  gatekeeper [command]

Available Commands:
  help        Help about any command
  list        Lists all commands the user can run
  run-plan    Runs the specified plan with given option
  self        Gatekeeper management commands

Flags:
  -h, --help   help for gatekeeper

Use "gatekeeper [command] --help" for more information about a command.
```

The `run-plan` command runs the plan defined in `plans.json`.
