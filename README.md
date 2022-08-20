# env-cmd

A simple program for executing commands using an environment from an env file.

## 💾 Install

`npm install env-cmd -D` or `yarn add env-cmd -D`

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
yarn env-cmd -f ./custom/path/.env,./other/path/.env node index.js
```

```sh
npx env-cmd -f ./custom/path/.env,./other/path/.env node index.js
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
   --env value, -e value   Additional environment variables      (accepts multiple inputs)
   --file value, -f value  Paths to env files (default: ".env")  (accepts multiple inputs)
   --help, -h              show help (default: false)
   --version, -v           print the version (default: false)
   --watch, -w             Watch for changes in .env files and reload them (default: false)
```
