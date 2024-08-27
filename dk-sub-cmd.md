To create a new sub-command for the Docker CLI, you can follow these steps:

1. **Create a new Go package**: Start by creating a new Go package that will contain the implementation of your sub-command. This package should be located within the Docker codebase, typically in the `cli/command/` directory.

2. **Implement the sub-command**: Within your new package, create a new file (e.g., `mysubcommand.go`) that will contain the implementation of your sub-command. This file should include the following:
   - A new struct that implements the `cli.Command` interface, which defines the behavior of your sub-command.
   - A `NewMySubCommand()` function that returns an instance of your sub-command struct.
   - The implementation of the `Run()` method, which is the entry point for your sub-command.

Here's an example of what your `mysubcommand.go` file might look like:

```go
package mysubcommand

import (
    "flag"
    "fmt"

    "github.com/docker/docker/cli"
    "github.com/docker/docker/cli/command"
)

type mySubCommand struct {
    // Add any necessary fields for your sub-command
}

// NewMySubCommand creates a new instance of the mySubCommand struct
func NewMySubCommand(dockerCli *command.DockerCli) *mySubCommand {
    return &mySubCommand{}
}

// Run implements the cli.Command interface and is the entry point for your sub-command
func (c *mySubCommand) Run(args []string) error {
    // Implement the logic of your sub-command here
    fmt.Println("My sub-command is running!")
    return nil
}

// Configure sets up the flags for your sub-command
func (c *mySubCommand) Configure(flags *flag.FlagSet) {
    // Add any necessary flags for your sub-command
}
```

3. **Register the sub-command**: In the `cli/command/commands.go` file, add a new entry for your sub-command in the `AddCommands()` function:

```go
import (
    // ...
    "github.com/docker/docker/cli/command/mysubcommand"
)

func AddCommands(dockerCli *command.DockerCli) {
    // ...
    cmd.AddCommand(mysubcommand.NewMySubCommand(dockerCli))
}
```

Keep in mind that this is a simplified example, and the actual implementation of a new sub-command may involve more complex logic, handling of flags, and integration with the existing Docker CLI codebase. You may also need to update the documentation and add tests for your new sub-command.

Yes, the Docker CLI (Command-Line Interface) has a pluggable sub-command architecture, which allows developers to extend the functionality of the Docker CLI by adding custom sub-commands.

The Docker CLI is designed to be extensible, and it uses a plugin system to allow third-party developers to add new sub-commands. These sub-commands can be used to interact with Docker in various ways, such as managing containers, images, networks, and volumes, or integrating with other tools and services.

The process of creating a custom sub-command for the Docker CLI involves the following steps:

1. **Implementing the sub-command**: The sub-command is implemented as a standalone executable that follows a specific interface and naming convention.
2. **Packaging the sub-command**: The sub-command is packaged as a Docker plugin, which is a directory containing the executable and a metadata file that describes the plugin.
3. **Installing the sub-command**: The plugin is installed on the Docker host, typically by placing the plugin directory in a specific location (e.g., `/usr/local/lib/docker/cli-plugins`).
4. **Using the sub-command**: Once the sub-command is installed, it can be accessed from the Docker CLI using the `docker <sub-command>` syntax.

The Docker CLI's pluggable sub-command architecture allows developers to extend the functionality of the Docker CLI to meet their specific needs, without modifying the core Docker codebase. This makes it easier to integrate Docker with other tools and services, and to create custom workflows and automation scripts.

**References:**

* [base-cli-plugin](https://github.com/docker/base-cli-plugin)
* [example plugin helloworld](https://github.com/docker/cli/tree/master/cli-plugins/examples)
* https://github.com/docker/buildx/blob/master/cmd/buildx/main.go
* [go-plugins-helpers](https://github.com/docker/go-plugins-helpers)
* https://www.docker.com/blog/build-your-first-docker-extension/
* https://docs.docker.com/desktop/extensions-sdk/
* https://docs.docker.com/extensions/extensions-sdk/architecture/
* https://github.com/docker/cli/blob/master/docs/extend/plugin_api.md
* https://docs.docker.com/extensions/extensions-sdk/build/backend-extension-tutorial/
* https://forums.docker.com/
* https://www.docker.com/blog/developing-go-apps-docker/



**Sketch:**

There are engine-plugins, and CLI-plugins. Engine plugins are for extending the Docker Engine, while CLI plugins are for extending the Docker CLI. CLI plugins are standalone executables that are invoked by the Docker CLI to perform specific tasks. They can be used to add new sub-commands, customize the behavior of existing commands, or integrate with external tools and services.

A docker extension extends Docker Desktop. An extension can consist of a Frontend, a Backend, and executibles. The cli plugin can be packaged as an executible of the extension. The Frontend can then call the plugin through the extension SDK. https://docs.docker.com/extensions/extensions-sdk/dev/api/backend/#invoke-an-extension-binary-on-the-host