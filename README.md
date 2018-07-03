# About

### Structure
Basically this application can be divided into three parts: __main application__, __services__ and __data stores__.

  - __main application__ The main application uses `gorilla mux` to handle http requests. It calls different `MakeEndpoint` methods defined in __services__ to create corresponding handler and handle incoming requests.

  - __services__ Service part is the middle layer between outside requests and the inner database services. It basically defines handlers that decode raw http requests:  
    - call __data store__ functions to do operations on the database
    - deal with errors
    - encode output for http response


  - __data stores__ The data stores uses SQL queries to communicate with database directly. It receives params passed from __services__ layers and abstract out data to form standard SQL queries. It also handles errors and wrap query results with formatted models for outside use.

# Get started

### Application setup
To Get started with the project, you need to have the following done first:

  1. Database setup
    - Install [PostgreSQL](https://www.postgresql.org/).

  2. Setup golang environment.
    - Download golang from [here](https://golang.org/dl/).

    <!-- - Setup account and create a `dbConfig.json` under root directory that specifies connections. Template can be found under `/deploy` -->

  2. Application setup
    - Clone the project and copy it under `$GOPATH/`. Usually `$GOPATH` is defined when you install golang. By default it is `/Users/<your-user-name>/go/src` for macOS. If you want to change your workspace, you might need to set `$GOPATH` environment variable. See [here](https://github.com/golang/go/wiki/SettingGOPATH).

    <!-- - Create a `httpConfig.json` under root directory that specifies server domain and port number. Template can be found under `/deploy`. You may set host to `localhost` and port number to be `8080`. -->

    - Ensure dependency: use `brew` to install `dep`:
      ```
      $ brew install dep
      $ brew upgrade dep
      ```

      Then under root directory, run

      ```
      $ dep init
      $ dep ensure
      $ dep ensure -update
      ```

    <!-- - Then you can start server by running

      ```
      $ go run main.go
      ``` -->
