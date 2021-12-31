# Watcher - Develop your programs easily

Watcher watches all the files present in the directory it is run from of the directory that is specified while running it and whenever a file is changed or a file is created/deleted from the directory, it runs the command specified while running it or a default command that it recognises from the contents of the current directory.

## Features

- [ ] Choose which directory to watch for file changes
- [ ] Press keys for force refreshing even when file change is not detected
- [ ] Automatically recognize commands for popular project structure

## Arguments

#### Command

The command you want to run when any file changes

```sh
watcher --cmd '<cmd>'
```

#### Directory

The directory that you want watcher to watch for file changes. (default: ".")

```sh
watcher --dir '<dir>'
```

#### Ignore

The files that you want to ignore. It also supports file globbing.

```sh
watcher --ignore '<files>'
```

#### Hidden

Whether the hidden files should also be watched for file changes. (default: true)

```sh
watcher --hidden false
```

## Examples

- Basic example (use default options):
  ```sh
  watcher
  ```
- Intermediate example:
  ```sh
  watcher --cmd 'go run .'
  ```
- Advanced example:
  ```sh
  watcher --cmd 'go run .' --dir '~/Projects/watcher' --hidden false
  ```
