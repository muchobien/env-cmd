<p align="center">
   <a aria-label="NPM Package" href="https://www.npmjs.com/package/@muchobien/env-cmd" target="_blank">
      <img alt="npm" src="https://img.shields.io/npm/v/@muchobien/env-cmd?color=success&logo=npm&style=flat-square">
   </a>
</P>

# env-cmd

A simple program for executing commands using an environment from an env file.

## 💾 Install

`npm install @muchobien/env-cmd -D` or `yarn add @muchobien/env-cmd -D`

## ⌨️ Basic Usage

**Environment file `./.env`**

```text
# This is a comment
ENV1=THANKS
ENV2=FOR ALL
ENV3=THE FISH
```

**Package.json**

```json
{
  "scripts": {
    "test": "env-cmd jest"
  }
}
```

**Terminal**

```sh
yarn env-cmd node index.js
```

```sh
npx env-cmd node index.js
```

### Using custom env file path

To use a custom env filename or path, pass the `-f` flag.

**Terminal**

```sh
yarn env-cmd -f ./custom/path/.env -f ./other/path/.env node index.js
```

```sh
npx env-cmd -f ./custom/path/.env -f ./other/path/.env node index.js
```

## 📜 Help

```text
NAME:
   env-cmd - Load environment variables from .env file and execute commands

USAGE:
   env-cmd [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   list, l  List environment variables
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --env value, -e value [ --env value, -e value ]    Additional environment variables
   --file value, -f value [ --file value, -f value ]  Paths to env files (default: ".env")
   --help, -h                                         show help (default: false)
   --interpolate, -i                                  Interpolate environment variables in command arguments (default: false)
   --override, -o                                     Override existing environment variables with new ones (default: true)
   --prefix value, -p value                           Prefix for environment variables
   --silent, -s                                       Ignore errors if .env file is not found (default: false)
   --version, -v                                      print the version (default: false)
   --watch, -w                                        Watch for changes in .env files and reload them (default: false)
```

## 🧬 Related Projects

[`toddbluhm/env-cmd`](https://github.com/toddbluhm/env-cmd) - Orginal project that inspired this one.
