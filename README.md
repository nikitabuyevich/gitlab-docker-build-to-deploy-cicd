![Logo of Gitlab](./gitlab.png)

# Docker Build to Deploy CI/CD &middot; [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](./LICENSE)

A generic `Gitlab CI/CD` pipeline which has the ability to build a `Docker` image on one machine and pull and run the image on another.

In order to get the full benefits of this CI/CD pipeline, you will need to have `two separate servers`. One which you treat as the `build` server and one which you want to `deploy` to.

If you wish to use one machine for everything, you can do so as well by registering your `Gitlab Runner` to have the `build`, `docker`, and `deploy` tags on it.

## Benefits for Using Separate Servers

- The `build server` stores multiple images that other `Docker` builds can reuse
- The `deploy server` only has one finalized image / container which it runs
  - Drastically reduces space
  - Separation of concerns
- Great for `Microservices` - Build once, deploy to many

## Getting Started

- Install `Docker` and `Gitlab Runner` onto the `build` and `deploy` servers
- Initiate the `Gitlab Runner` registration on both servers
- For the `build` server, add `build` and `docker` as the tags
- For the `deploy` server, add `deploy` as the tag
- Set `shell` as the `executor` on both servers

## Sample Project

Feel free to modify the [docker-compose.yml](sample-project/docker-compose.yml) or the [.gitlab-ci.yml](sample-project/.gitlab-ci.yml) files to satisfy your needs. You can largely ignore the [Dockerfile](sample-project/Dockerfile) as it pertains mainly to `go`.

This is just one way to structure the project but the CI/CD pipeline will be roughly the same for other languages as well.

### What the Sample Project is Doing

- Runs a simple `go` web server which returns "Hello World" on port 3000
- The `Dockerfile` builds the application and its tests in two separate binaries
- The `docker-compose.yml` file has two separate services - One for the application and one for the tests
- The `.gitlab-ci.yml` file runs the tests through `docker-compose`
- The `Docker` image name is called `go-docker` which is later referenced in the `.gitlab-ci.yml` file - You'll need to update the image name for your project
- The image name that gets pushed out to Gitlab's Docker Registry uses the following format: `<GitLab project registry URL>:<branch name>`
  - An example of the final image name is `registry.gitlab.com/nikitabuyevich/go-docker:master`
  - Feel free to change the tag to satisfy your needs
- Anything that references `$CI_*` is a specific `Gitlab CI/CD` environment variable - See [Predefined environment variables](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html)

### Note

If you are planning to run [CentOS](https://www.centos.org/) as your operating system on your servers, you can utilize the [Bootstrap / Harden CentOS](https://github.com/nikitabuyevich/bootstrap-centos) bash script which will secure your server and install and register `Docker` and `Gitlab Runner`.

### Prerequisites

- [Docker](https://docker.com/) - Build and run images / containers
- [Gitlab Runner](https://docs.gitlab.com/runner/) - Integrate a Gitlab CI/CD pipeline with a server

## Built With

- [Gitlab CI/CD](https://docs.gitlab.com/ee/ci/) - A tool built into GitLab for software development through continuous integration, delivery, and deployment.

## Authors

- **Nikita Buyevich** - [nikitabuyevich.com](https://nikitabuyevich.com/)

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
