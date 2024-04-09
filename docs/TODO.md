# Overview and planning

Implementation:
- Any `gown` command reads the entire directory of source
- From the ast create a `Project` representation
- Perform operations on `Project`, which in turn modify the underlaying ast
- Write `Project` changes back to files

Features:
- Everything revolves around modules and data.
- Create module and add data objects to it.
- Optionally generate Stores for data.
- Optionally generate handlers for such data stores.
- Stores act as basic database operations (create, read, update, delete, read_all)
- Queries act as advanced reading operations? (derive from cli args)
- Commands act as advanced writing operations? (derive from cli args)
- Try to generate SQL queries also
- Handle migrations in case of sql
- Add support for mutliple data sources, but have a default
- This could be accomplished if DataStore was an interface, so
  an SQL implementation could be injected. Because it's not written
  by hand, I would assume it's okay to make such an abstraction.

Examples:
- initialize project
  `gown init <project name>`
  `gown init first_project`
- create module in `app/`
  `gown add module [--route <base route>] <module name>`
  `gown add module --route /user-management UserManagement`
- create data source
  `gown add datasource <type> <source/uri>`
  `gown add datasource rdbms "sqlite3://memory"`
- create data in module or top level
  `gown add data [--module <module name>] <data struct name> <fields and constraints>`
  `gown add data --module UserManagement User id:int:identity name:string:unique`
- create store
  `gown add store [--with-handlers] <data store name>`
  `gown add store UserStore`
- create handler (with optional presets)
  `gown add handler [--protected <roles>] [--basic, --webhook, --websocket, --module <module name>] <route>`
  `gown add handler --module UserManagement /help`
  - will generate route with module base and <route>
    added to it `/user-management/help`


Project strucutre:
- app/
 - app.go (contains main `Application` struct)
 - users/
  - user.go (contains user model)
  - user_store.go (contains `UserStore` service, refereced by `Application`)
  - profile_store.go (contains `ProfileStore` service, refereced by `Application`)
  - prefereces_store.go (contains `PreferencesStore` service, refereced by `Application`)
- web/
 - static
  - favicon.ico
 - api.go
- cmd/
 - web/
  - main.go (entry point to web application)
- setup/
 - setup.go (dependency injection to create app and web structs, load config and env vars, etc)
- Makefile
- .gitignore
- README.md

```go
// app/app.go
type Application struct {
    UserStore UserStore
}

// app/users/user_store.go
type UserStore struct {
  db *sqlx.DB
}

// same for other services/stores in users module...
```

## Structure Idea:

Create Components to represent  the structure of the project (Project, Module, Store, etc).
The Components should be just data and references to ast nodes.

The Loader reads the project files and creates a project representation via Components.
Implement Queries to simplify search of go objects, like functions, structs etc.
Implement Commands to apply changes to the project structure based on the loaded Components.

The actual command functions in the CLI package would
call the Loader to load the project components and then
invoke the Commands to perform operations.
The commands would use queries to find objects and make changes.
Implementing some common functions to add object to the AST and make
changes would be great. I don't know if they should be called something
like Changes or Mutations or Operations.

For example:
- `gown add module User` calls `cli.addModule`
- `cli.addModule` loads project via `Loader.Load`
- `cli.addModule` calls `commands.AddModule(project)`
- `commands.AddModule` creates `AddFileOperation("user.go")` to create new file
- `commands.AddModule` creates `AddStructOperation("user.go", structName, fields)`
  to add a new struct to the newly created file
- Then finally `commands.AddModule` applies the operations in order to achieve the desired result.
- In case one of the operations fails, the command should be able to revert the changes to not break the app.

Pay attention to how the cli commands are simply loading the project and performing a Command.
The Command has the project state in the components, so it would just perform manipulation on those Components
via small Operations/Mutations/Changes that are reversible. Commands can use Queries to get information or
references to some parts of the code, like getting a function in some module or a method on a service component.
Queries, unlike Operations, don't mutate any state but just do lookup on the AST. A Command does not need to
perform mutations, but can just write to console; for example a ListModules command could just use a Query
to get a list of module names to print to console.

- Does it make sense to have an undo on each Operation which would restore the project to it's previous state?
Each operation would be a small atomic modification to the project structure, like create file, create struct,
add field to struct, add function, add method, etc. Seems like it would be simple to undo such a small operation.


NOTE: Idea to add go-lua to support running own scripts that use Operations and Queries?

