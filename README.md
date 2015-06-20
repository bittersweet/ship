# Ship

## What?

A commandline tool to ship your projects. I've got a ton of projects that all
have different ways to deploy, rsync, capistrano (2, 3...) etc, and this tool
makes it so I don't have to remember it all for every project.

Just define steps in a project local `.ship` file and run `ship`!

## Installation

Clone this repository and run `make`, Go is required.

## Notes

* Doesn't seem to work in realtime with Capistrano 3, they have some funky
issues with stdout/stderr though [1](https://github.com/capistrano/sshkit/issues/100)
[2](https://github.com/capistrano/sshkit/issues/86)
