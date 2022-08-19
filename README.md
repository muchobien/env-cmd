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
Usage: env-cmd [options] <command> [...args]

Options:
  -v, --version                       output the version number
  -f, --file [paths]                  Custom env file path (default paths: ./.env)
  -h, --help                          output usage information
```
