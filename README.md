# Watcher - Develop your programs easily

Watcher watches all the files present in the directory it is run from of the directory that is specified while running it and whenever a file is changed or a file is created/deleted from the directory, it runs the command specified while running it or a default command that it recognises from the contents of the current directory.

## Features

- [x] Choose which directory to watch for file changes
- [x] Specify amount of time to wait before running the command after a file change occurs
- [x] List the files that are being watched
- [x] Choose whether to run the command on startup
- [x] Set the maximum number of files that can be watched
- [x] Supports never ending processes such as dev servers as the cmd
- [ ] Have a default ignore list
- [ ] Press keys for force refreshing even when file change is not detected
- [ ] Automatically recognize commands for popular project structure

## Arguments

#### Command

The command you want to run when any file changes

```sh
watcher --cmd '<cmd>'
```

#### Directory

The directory that you want watcher to watch for file changes. _(default: ".")_

```sh
watcher --dir '<dir>'
```

#### Ignore

The files that you want to ignore. It also supports file globbing.

```sh
watcher --ignore '<files>'
```

#### Hidden

Whether the hidden files should also be watched for file changes. _(default: true)_

```sh
watcher --hidden false
```

#### Delay

The amount of time to wait before running the specified command after a file change occurs. _(default: 500ms)_

```sh
watcher --delay 1000ms
```

#### Run on Start

Whether the specified command should run when watcher has first started. _(default: true)_

```sh
watcher --run-cmd-on-start false
or
watcher -r false
```

#### List on Start

Whether the list of files being watched should be printed when watcher has first started. _(default: false)_

```sh
watcher --list-on-start true
```

#### Limit

The maximum number of files that can be watched. _(default: 10000)_

```sh
watcher --limit 50000
or
watcher -l 50000
```

> ⚠️ Every system has a maximum value which cannot be exceeded. To find it look at:
>
> - **Linux**: `/proc/sys/fs/inotify/max_user_watches` contains the limit, reaching this limit results in a "no space left on device" error.
> - **BSD / OSX**: `sysctl` variables `kern.maxfiles` and `kern.maxfilesperproc`, reaching these limits results in a "too many open files" error.

## Examples

- **Basic** example (use _default_ options):
  ```sh
  watcher
  ```
- **Intermediate** example:
  ```sh
  watcher --cmd 'go run .' --delay 1s
  ```
- **Advanced** example:
  ```sh
  watcher --cmd 'go run .' --dir '~/Projects/watcher' --hidden false -l true
  ```
