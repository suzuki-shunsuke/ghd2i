# Usage

<!-- This is generated by scripts/generate-usage.sh. Don't edit this file directly. -->

```console
$ ghd2i help
NAME:
   ghd2i - A new cli application

USAGE:
   ghd2i [global options] command [command options]

VERSION:
   0.1.0 (6bebd9eb3658e151ced8ecd9939994e0b21273d5)

COMMANDS:
   version         Show version
   run             Create GitHub Issues from GitHub Discussions
   create-config   Create a configuration file
   get-discussion  Get discussion and output the data
   completion      Output shell completion script for bash, zsh, or fish
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level value  log level
   --log-color value  Log color. One of 'auto' (default), 'always', 'never'
   --help, -h         show help
   --version, -v      print the version
```

## ghd2i run

```console
$ ghd2i help run
NAME:
   ghd2i run - Create GitHub Issues from GitHub Discussions

USAGE:
   ghd2i run [command options]

DESCRIPTION:
   Create GitHub Issues from GitHub Discussions

   $ ghd2i run https://github.com/suzuki-shunsuke/test-github-action/discussions/55


OPTIONS:
   --config value, -c value  configuration file path. Configuration file is optional. If \.ghd2i.yaml exists, it's used as the configuration file by default
   --data value              data file path. If data file path is set, the data is read from the file instead of calling GitHub API
   --lock value              Whether created issues are locked. One of 'auto', 'always', 'never'. Auto means that the issue is locked if the discussion is locked (default: "auto")
   --close value             Whether created issues are closed. One of 'auto', 'always', 'never'. Auto means that the issue is closed if the discussion is closed (default: "auto")
   --repo-owner value        Repository owner where issues are created. By default, issues are created in the repository of each discussion
   --repo-name value         Repository name where issues are created. By default, issues are created in the repository of each discussion
   --dry-run                 Instead of creating issues, output issue body and comment bodies (default: false)
   --help, -h                show help
```

## ghd2i create-config

```console
$ ghd2i help create-config
NAME:
   ghd2i create-config - Create a configuration file

USAGE:
   ghd2i create-config [command options]

DESCRIPTION:
   Create a configuration file.

   $ ghd2i create-config


OPTIONS:
   --help, -h  show help
```

## ghd2i get-discussion

```console
$ ghd2i help get-discussion
NAME:
   ghd2i get-discussion - Get discussion and output the data

USAGE:
   ghd2i get-discussion [command options]

DESCRIPTION:
   Get discussion and output the data

   $ ghd2i get-discussion <discussion-url> [<discussion-url> ...]


OPTIONS:
   --help, -h  show help
```

## ghd2i version

```console
$ ghd2i help version
NAME:
   ghd2i version - Show version

USAGE:
   ghd2i version [command options]

OPTIONS:
   --json      (default: false)
   --help, -h  show help
```

## ghd2i completion

```console
$ ghd2i help completion
NAME:
   ghd2i completion - Output shell completion script for bash, zsh, or fish

USAGE:
   ghd2i completion command [command options]

DESCRIPTION:
   Output shell completion script for bash, zsh, or fish.
   Source the output to enable completion.

   e.g.

   .bash_profile

   source <(ghd2i completion bash)

   .zprofile

   source <(ghd2i completion zsh)

   fish

   ghd2i completion fish > ~/.config/fish/completions/ghd2i.fish


```
