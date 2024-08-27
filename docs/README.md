# Documentation

**Nota bene:** This documentation is hypothetical! It serves to articulate the user requirements and the design of the application.

## Usage

### Add a new configuration or new configuration file
```shell
$ docker config-app add-config [<config-file>]
```

If no `<config-file>` is provided, you will be prompted to enter a file name.

If the file already exists, new configs will be appended to the file.

Config-file-extension is optional in the arg, defaults to `.env`.

### Set configuration values
```shell
$ docker config-app set-config [<config-file>]
```

Config-file-extension is optional and uses the template to determine the file extension.

If no config-file is provided, all config files for which a template exists will be updated.

### Options

* include the export keyword
* template framework: default to go templates.
* config language: default to shell.
* template location: default to `./conf`.
* config location: default to `./conf`.
* default config-file-extension: default to `.env`.
