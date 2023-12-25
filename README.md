# Lindir

**Lindir** is a CLI tool for **managing hard links across multiple directories**.

It is designed to solve the limitations of traditional hard links and symbolic links. While hard links allow multiple access points to a single file, they cannot be used for directories. Symbolic links, on the other hand, can link directories but may not always be compatible with certain programs.

Lindir allows you to mimic directory linking through a **Git-inspired approach**. Lindir tracks and manages hard links by creating a `.lindir` directory within your chosen directory. This approach is particularly beneficial in scenarios such as managing dot files in a Git repository.

> [!WARNING]
> In most cases, symbolic link is better choice.

## Installation

**Prerequisites**

- Git (optional)
- [Go](https://go.dev/dl/) 1.21.4 or later
  - It may work with earlier versions, but haven't tested it.

### Step 1: Clone this repository

```shell
git clone <this repository>
```

### Step 2: Install Lindir

```shell
go install

# Use this command to get a smaller binary
go install -ldflags "-s -w"
```

### Step 3: Check installation

```shell
lindir -h
```

## Before using Lindir

Before using Lindir, you need to understand the following concepts.

### Initialization

You need to initialize a directory before using Lindir.

Initialized directory could be connected(hard linked) to other directories. Connected directories will share the same files and directories that are hard linked.

To know which directories are connected, Lindir will create a `.lindir` directory within the initialized directory. This directory contains a `connector` file that tracks all the paths where this directory is hard linked.

To know what files are hard linked, Lindir will also create a `tracker` file that tracks what files are hard linked between directories.

> [!WARNING]
> Do not modify the `.lindir` directory manually. It may cause unexpected behavior.

### Directory hard linking(connecting)

Lindir uses a Git-inspired approach to mimic directory hard linking. Since it is not possible to hard link directories directly, we say that a directory is **connected** to another directory.

When directories are connected, Lindir will track what files are currently hard linked between directories.

### Git-inspired approach

As mentioned above, Lindir "tracks" what files are hard linked between connected directories. Lindir does not automatically create/delete hard links. Instead, it just tracks the changes(adding/deleting a file) you make in the connected directories.

This is where the Git-inspired approach comes in. Since Lindir only tracks the changes you make, you need to manually `push` or `sync` the changes.

So the whole process is as follows:

1. Initialize a directory
2. Modify ignore files (`.lindirignore`, `.gitignore`)
3. Link directories (= Connect directories)
4. Make changes (add/delete files)
5. Check status
6. Push or sync changes

## Commands

### `init`

Initialize a directory.

```shell
lindir init [<directory>]
```

This command will create a `.lindir` directory within the current directory. This directory contains a `connector` file that tracks all the paths where this directory is hard linked. It also contains a `tracker` file that tracks what files are hard linked between directories.

Specifying a directory will initialize the specified directory instead of the current directory.

### `link`

Link directories (= Connect directories).

```shell
lindir link [<from>] <to>
```

This command will **connect** the `<from>` directory to the `<to>` directory. The `<from>` directory must be initialized before linking.

If the `<from>` directory is not specified, the current directory will be used.

### `status`

Check what files are newly added or deleted.

```shell
lindir status [<directory>]
```

This command will check what files are newly added or deleted in the specified directory. If the directory is not specified, the current directory will be used.

### `push`

Push changes to connected directories.

```shell
lindir push [<directory>]
```

This command will push the changes to the connected directories. If the directory is not specified, the current directory will be used.

'Changes' means the files that are newly added or deleted (= files that are not yet hard linked across connected directories). You can check the changes using the `status` command.

### `sync`

Sync changes with connected directories.

```shell
lindir sync [<directory>]
```

This command will sync the changes with the connected directories. Syncing is same as running `push` in all connected directories.

If the directory is not specified, the current directory will be used.

### `unlink`

Unlink directories (= Disconnect directories).

```shell
lindir unlink [<directory>]
```

This command will **disconnect** the specified directory from all connected directories. By disconnecting, hard linked files will be a whole new copy of the original file. Also, the `.lindir` directory will be deleted.

If the directory is not specified, the current directory will be used.

## Ignore files

You can ignore files by creating a `.lindirignore` file in the directory. This file uses the pattern syntax of Go language's [filepath.Match](https://pkg.go.dev/path/filepath#Match).

> [!TIP]
> Add `.git` directory to `.lindirignore` file to ignore Git files.
