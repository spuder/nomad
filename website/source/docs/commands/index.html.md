---
layout: "docs"
page_title: "Commands (CLI)"
sidebar_current: "docs-commands"
description: >
  Nomad can be controlled via a command-line interface. This page documents all
  the commands Nomad accepts.
---

# Nomad Commands (CLI)

Nomad is controlled via a very easy to use command-line interface (CLI).
Nomad is only a single command-line application: `nomad`, which
takes a subcommand such as "agent" or "status". The complete list of
subcommands is in the navigation to the left.

The Nomad CLI is a well-behaved command line application. In erroneous cases,
a non-zero exit status will be returned. It also responds to `-h` and `--help`
as you would most likely expect.

To view a list of the available commands at any time, just run Nomad
with no arguments. To get help for any specific subcommand, run the subcommand
with the `-h` argument.

Each command has been conveniently documented on this website. Links to each
command can be found on the left.

## Autocomplete

Nomad's CLI supports command autocomplete. Autocomplete can be installed or
uninstalled by running the following on bash or zsh shells:

```shell
$ nomad -autocomplete-install
$ nomad -autocomplete-uninstall
```

## Command Contexts

Nomad's CLI commands have implied contexts in their naming convention. Because
the CLI is most commonly used to manipulate or query jobs, you can assume that
any given command is working in that context unless the command name implies
otherwise.

For example, the `nomad job run` command is used to run a new job, the `nomad
status` command queries information about existing jobs, etc. Conversely,
commands with a prefix in their name likely operate in a different context.
Examples include the `nomad agent-info` or `nomad node drain` commands,
which operate in the agent or node contexts respectively.

### Remote Usage

The Nomad CLI may be used to interact with a remote Nomad cluster, even when the
local machine does not have a running Nomad agent. To do so, set the
`NOMAD_ADDR` environment variable or use the `-address=<addr>` flag when running
commands.

```shell
$ NOMAD_ADDR=https://remote-address:4646 nomad status
$ nomad status -address=https://remote-address:4646
```

The provided address must be reachable from your local machine. There are a
variety of ways to accomplish this (VPN, SSH Tunnel, etc). If the port is
exposed to the public internet it is highly recommended to configure TLS.
