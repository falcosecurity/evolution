
# Falco Project Evolution

This repo aims to document the evolution process of The Falco Project.

It is also a place for various community-maintained resources. In particular, this repository provides a space for the community to work, document, and build out third-party integrations in a safe, and productive way.

## Adoption model

The criteria will remain loose, and tighten as needed at the discretion of the Falco open source community.

### Sandbox

The "*Sandbox*" level serves as a place for the community to test-drive ideas/projects/code. 

Resources with this status can be found within this repo:

- [`/deploy`](deploy/) - deployment resources
- [`/examples`](examples/) - various examples
- [`/integrations`](integrations/) - third-party integrations
- *More will come soon!*

> When adding a new directory, propose your motivations to the [maintainers](OWNERS).

### Incubating

The "*Incubating*" level refers to those projects (usually promoted from "*Sandbox*") that need their repository.

This status is assigned as needed, and can best be measured by the need to cut a release and use the GitHub release features.

### Official Support

As the need for a project grows, it can ultimately achieve the highest and most coveted status within The Falco Project: "*Offical support*."

Currently, only those artifacts included in the [falcosecurity/falco](https://github.com/falcosecurity/falco) repository got the "Official support" status. These artifacts will be refined and amended as per the [Falco Artifacts Scope - Part 2](https://github.com/falcosecurity/falco/blob/master/proposals/20200506-artifacts-scope-part-2.md).

The full list of those artifacts can be found within the [official documentation](https://falco.org/docs/download/.)

## Contributing

See the [contributing guide](https://github.com/falcosecurity/falco/blob/master/CONTRIBUTING.md) and the [community code of conduct](https://github.com/falcosecurity/falco/blob/master/CODE_OF_CONDUCT.md).

## Join the Community

To get involved with the evolution of The Falco Project, please visit [the community repository](https://github.com/falcosecurity/community) to find more.

# Adding a new project walk-through 

So you have an idea for some new Falco work and are ready to get started. 

Falco values speed, and innovation so we try to never stand in the way or block progress in any way. As projects gain traction, we exchange flexability for structure. For instance a project that is 1 day old is relatively ungoverned, wheras the official production ready Falco packages are heavily governed. 

Below find instructions on how to contribute your work to Falco.

### Getting Started (sandbox) 

**Quick, Simple, Ungoverned, Unofficial**

If you have a small change, or just would like to add a small amount of sample code to share with the community you can add it to the evolution repository [here](https://github.com/falcosecurity/evolution). We call changes like this `sandbox` contributions, and are the most flexible and ungoverned contributions. 

### Getting Started (incubation)

**Concept of OWNERS, Release Process, New falcosecurity Repository, Unofficial**

If you would like to start a new project to donate to Falco you can get started by creating a [new github repository](https://github.com/new) under your own account or organization. Work on your project until you are ready to bring it to the Falco community. There is no standard here for how mature this needs to be, it can be a proof of concept or a well used tool. Open up an issue in the [issue tracker](https://github.com/falcosecurity/evolution/issues) for the community to adopt the project. A great example of this is the [helm chart](https://github.com/falcosecurity/evolution/issues/12) adoption issue.

The most **critical** component of an `incubation` project is that it **must** have an `OWNERS` file and a concept of a release for the community. Without these defined the community will not be able to adopt a new `incubation` repository. 

### Getting Started (official support) 

**The most coveted of all status, official support**

After an `incubation` level project has at least 1 well understood production user, a steady release process, and timely response from the `OWNERS` the Flaco maintainers can decide to promote this to `official support` such that it will be added to the official [downloads](https://falco.org/docs/download/) page on falco.org.
