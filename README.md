# tctl-plugin-template

Use this template to quickly create a new CLI or [tctl](https://github.com/temporalio/tctl) plugin

## Build

Run `make build` to build the project

## Quick Start

The template shows an examples of:
 
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
