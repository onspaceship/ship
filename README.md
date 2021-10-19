<p align="center">
  <img src="https://static.onspaceship.com/FullColor.svg" width="150">
</p>

<h3 align="center">
  Ship
</h3>

<p align="center">
  The Spaceship CLI tool
</p>

---

Ship is a small CLI tool for working with the Spaceship platform.

## Commands

Ship provides several commands, with some still in development. You can see the ones currently available by running `ship` with no command or `ship help` to get more detailed help output.

Currently implemented commands will be checked off below:

- ✅ **`login`** - Logs ship into your Spaceship account.
- ✅ **`logout`** - Logs ship out of the currently logged-in Spaceship account.
- 🔳 **`delivery`** - A set of sub-commands for deliveries.
  - 🔳 **`delivery list`** - Gets a list of recent deliveries.
  - 🔳 **`delivery setup`** - Sets up delivery of an app to your current Kubernetes cluster.
- 🔳 **`build`** - A set of sub-commands for builds.
  - 🔳 **`build list`** - Gets a list of recent builds.
  - 🔳 **`build create`** - Creates a new manual image build from the current repo.
- 🔳 **`k8s`** - A set of sub-commands for working with Kubernetes.
  - 🔳 **`k8s install`** - Installs the Spaceship Agent into your current Kubernetes cluster.
  - 🔳 **`k8s new`** - Set up a Deployment in Kubernetes for the current repo.
- 🔳 **`ship it`** - Starts a delivery for the current repo. :shipit:

## Development

Ship requires [go 1.17 or higher](https://golang.org/) to build.

## Contributing

1. [Fork the repo](https://github.com/onspaceship/ship/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request here on GitHub
