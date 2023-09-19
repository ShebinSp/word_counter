* You can learn here about the components used in go.yml

1. Workflow Setup
    * Workflow is an automated process made up of several jobs that are carried out when triggered by events. Workflows are defined in YAML
      files and stored in a .github/workflows directory in the repository.

    * `name: Go`: This sets a name for the workflow as "Go".

2. Triggering the workflow
    * Events
      Events are defined triggers that start a workflow. Such as creating a branch, opening a pull requst or push request etc.

    * `on: `:Specifies when the workflow should be triggered.
    * `push:`: This workflow is triggered when  code is pushes to the repository.
        * `branches: ["main"]`: It specifies that the workflow should only trigger when changes are pushed to the "main" branch.
    * `pull_request`:This workflow is also triggered when a pull request in made.
        * `branches: ["main"]`: Similar to the push trigger, it specifies that pull request to the "main" branch should trigger the workflow.

3. Jobs
    * Jobs are tasks that are executed in a workflow when an event is triggered. A workflow can have multiple jobs runnig in parallel.

    *  `build:`: This defines a job named "build".

4. Job Configuration
    * Runner
       A runner is a series of tasks that are executed in a workflow when triggered by an event. Each runner is responsible for executing a 
       single job.

    * `runs-on: ubuntu-latest`: This job runs on a virtual machine with the latest version of Ubuntu.

5. Steps
    * Actions
       Actions are used to perform complex tasks that you may import into workflows, such as sending a notification email. You can 
       build your own actions or reuse open-source actions available on the GitHub marketplace.
       
    * `uses: actions/checkout@v3`: This step checks out(downloads) your repository's code, so that the workflow can work with it
    * `name: Set up Go`: This step up the Go environment for the job.
        * `uses: actions/setup-go@v4`: It uses a pre-defined GitHub Actions setup for Go.
        * `with:`: Specifies configuration options for setting up Go.
            *   `go-version: '1.20'`: This sets the Go version to 1.20.
    * `name: Build`: This step compiles(builds) the Go code.
        * `run: go build -v ./...`: It runs the Go build command to build all packages (./...) in the project. The '-v' flag makes it verbose,
           so it shows more detailed output.
    * `name: Test`: This step runs tests for the Go code.
        * `run: go test -v ./...`: It runs the Go test command to execute tests for all packages('./...') in the project.