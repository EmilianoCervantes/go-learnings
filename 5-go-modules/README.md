# Go Modules

Just like how _Python_ has `pip` and _Node_ has `npm`, **Go** has `go get`.

## For this example

We will use **go tour** which is library Google created for introducing people to modules in Go.

Here is a guide with steps you can follow.

At the end of this guide I'll write about the more concrete example I used.

## Init a project

```bash
go mod init <project_name>
```

The recommendation for the project name is to be the name of the repo where you will keep your project.

Not just how you want to name the folder where your project lives locally.

For example:

```bash
go mod init github.com/<user_name>/<project_name>
```

## Get a library

```bash
go get <library>
```

Example based on [A Tour of Go - Go offline](https://go.dev/tour/welcome/3):
```bash
go get golang.org/x/website/tour
```

### Now on the terminal

Run
```bash
tour
```

#### On Mac

If you face the issue where the commands is not being recognized:

```bash
❯❯❯ tour
zsh: correct tour to tput [nyae]?
```

Just say **no**

```bash
❯❯❯ tour
zsh: correct tour to tput [nyae]? n
```

And you should see:

1. The terminal starting a service in the `port 3999`.
2. Your browser opening a new tab.
   1. If not, try opening http://127.0.0.1:3999.

### Flags
You can add some flags:
- -v: for verbose, so you can see everything that is happening behind the scenes.
- -u: even when the library has already bin downloaded, it will tell `go get` to refetch it.

```bash
go get -v -u github.com/spf13/cobra@latest
```

### Having issues?

Run

```bash
export GOPROXY=direct
```

## Install a library

```bash
go install <library>
```

Example based on [A Tour of Go - Go offline](https://go.dev/tour/welcome/3):
```bash
go install -v golang.org/x/website/tour@latest
```

### If you have any trouble using these commands

Remember you can always run

```bash
go help install
```

to know more about the process.

### Having issues?

Run

```bash
export GOPROXY=direct
```

## `go get` vs `go install`

They might look the same, but each have their use cases and even when you can use either of these commands, it is important to know the differences.

Read the following resources:

1. https://tip.golang.org/doc/go1.16#modules
2. https://medium.com/@chaewonkong/difference-between-go-get-and-go-install-in-go-a076d7352186

## Verify the external code is not corrupt

```bash
go mod verify
```

### Example

In the next command details, there's an example of a use for `verify`

## Replace source code

You'll find yourself using this command whenever you have to edit one of the libraries you downloaded to use for some use case or need of your application.

```bash
go mod edit -replace online_lib_dir=local_lib_path
```

- [replace - Add a replacement of the given module path and version pair](https://fig.io/manual/go/mod/edit)

### Example

```bash
go mod edit -replace github.com/labstack/echo/v4=./echo/v4
```

Note: make sure the files you modified are ACTUALLY in the PATH you provided in the command.

You'll usually run `go mod verify` to make sure all modules are healthy.

## Remove the replace

```bash
go mod edit -dropreplace online_lib_dir
```

- [dropreplace - Drops a replacement of the given module path and version pair](https://fig.io/manual/go/mod/edit)

## Put together all the vendor code (third parties) used in our code

```bash
go mod vendor
```

This will create a folder in the project's directory.

It will contain all the code needed for our project to run as expected.

## Remove all unused external packages and fix some dependency issues

To check all the packages in your project and their dependencies to clean up and optimize your go. mod file

```bash
go mod tidy
```

If you want to clean your modules, run this command.

If by any chance you just installed a module, but the go.mod is throwing some complaints.

### For Example

You just installed a new package and you want to fix some issues you are pointed out to,
BUT you haven't used it yet.

Open the file where you will use this library and import it like this:

```go
import _ "github.com/spf13/cobra"
```

We use "_" to import packages that we are not currently using.

#### Before

```go
require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.8.1 // indirect // github.com/spf13/cobra should be direct [go mod tidy]
	github.com/spf13/pflag v1.0.5 // indirect
)
```

#### After running `go mod tidy`

```go
require github.com/spf13/cobra v1.8.1

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
```

Be sure to do that so you don't undo your installation.

## More about go modules

```bash
go help mod
```

## go.sum

If you're familiar with node and npm, you might think it is a type of lock file.

While it is not exactly a lock file, it works more as a dependency checker.

Remember: it does't necessarily represent the actual dependencies used in your project.

It's purpose is to verify dependencies during the building phase.

# Echo

Performed the [Quick Start Guide](https://echo.labstack.com/docs/quick-start).

## To Do

As you install echo, be sure to use tools such as `go mod tidy`.