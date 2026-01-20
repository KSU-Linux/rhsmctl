# rhsmctl

A command-line tool to control resources using Red Hat Subscription Manager API.

## API Token

In order to use `rhsmctl` you need to obtain an API token from the [customer portal](https://access.redhat.com/management/api).

## Usage

`rhsmctl` uses sub-commands to interact with RHSM API.

```console
rhsmctl --api-token ${API_TOKEN} <sub-command> <options>
```

The current supported sub-commands are:

* __account__: Show the details of your account.
* __allocation__: Manage allocations.
* __allocation-entitlements__: Manage allocation entitlements.
* __cloud-access__: Manage cloud access.
* __cloud-access-accounts__: Manage cloud access accounts.
* __errata__: Manage errata.
* __images__: Manage images.
* __organization__: Show the details of your organization.
* __packages__: Manage packages.
* __subscription__: Manage subscriptions.
* __system__: Manage systems.
* __system-entitlements__: Manage entitlements of a system.

To get a sub-command's detailed description, options and usage, use the `-h` or `--help` option.

```console
rhsmctl <sub-command> --help
```

## Config File

To avoid having to use the `--api-token` option on every command, you can save it to a config file.

`rhsmctl` checks the following files for your API token:

1. `$HOME/.config/rhsmctl/config.yml`

The config file should be in YAML format.

```yml
api-token: "${API_TOKEN}"
```
