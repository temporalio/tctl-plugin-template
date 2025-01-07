**:warning: Deprecation Notice :warning:**

The `tctl` CLI is now deprecated in favor of Temporal CLI. <br />
This repository is no longer maintained. <br />
Please use the new utility for all future development. <br />

* New [Temporal CLI repository](https://github.com/temporalio/cli).
* [Temporal CLI Documentation site](https://docs.temporal.io/cli).

# tctl-plugin-template

> **Nota bene**: The tctl plugin feature is still experimental and is not documented or officially supported. There are going to be some further changes & UX improvement

Use this template to quickly create a new CLI or [tctl](https://github.com/temporalio/tctl) plugin

## Quick Start

Run `make build` to build the project

The template shows the examples of:
 
- configuring a value 

``` 
$ tctl-my-plugin set-hello --value Buzz
```

- reading and outputting from a config
```
$ tctl-my-plugin hello
Hello Buzz
```

In addition, it provides an easy access to
 - Temporal SDK `temporalClient`
 - CLI kit https://github.com/temporalio/tctl-kit
 - Makefile to build and test your project

## Running as a tctl plugin

Place the resulting binary anywhere in $PATH and tctl will pick it up as a plugin:

```
$ tctl my-plugin hello
```
