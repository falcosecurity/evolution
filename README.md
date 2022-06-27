
<p><img align="right" src="./img/pidgeotto.gif"/></p>
<p></p>

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

As the need for a project grows, it can ultimately achieve the highest and most coveted status within The Falco Project: "*Official support*."

## Projects

As per [our governance model](https://github.com/falcosecurity/.github/blob/master/GOVERNANCE.md#adding-new-projects-to-the-falcosecurity-github-organization), to request a project to be added or promoted, please open an [issue](https://github.com/falcosecurity/evolution/issues/new/choose), and choose the appropriate template. Once sufficient discussion has taken place and the proposal has been accepted, the project will be listed in the table below with the given status.


| Project | Status | Note |
| --- | --- | --- |
| [falcosecurity/charts](https://github.com/falcosecurity/charts) | Official support | Each chart inherits its status from the related project, for example the Falco chart is official |
| [falcosecurity/falco](https://github.com/falcosecurity/falco) | Official support | The list of official artifacts can be found within the [official documentation](https://falco.org/docs/download/). These artifacts will be refined and amended as per the [Falco Artifacts Scope - Part 2](https://github.com/falcosecurity/falco/blob/master/proposals/20200506-artifacts-scope-part-2.md).  |
| [falcosecurity/falco-website](https://github.com/falcosecurity/falco-website) | Official support | | 
| [falcosecurity/libs](https://github.com/falcosecurity/libs) | Official support | Artifacts will be available once [Versioning and release process of the libs artifacts](https://github.com/falcosecurity/libs/blob/master/proposals/20210524-versioning-and-release-of-the-libs-artifacts.md) proposal is fully implemented. |
| [falcosecurity/plugins](https://github.com/falcosecurity/plugins) | Official support | |
| [falcosecurity/plugin-sdk-go](https://github.com/falcosecurity/plugin-sdk-go) | Official support | |
| [falcosecurity/client-go](https://github.com/falcosecurity/client-go) | Incubating | |
| [falcosecurity/client-py](https://github.com/falcosecurity/client-py) | Incubating | |
| [falcosecurity/client-rs](https://github.com/falcosecurity/client-rs) | Incubating | |
| [falcosecurity/driverkit](https://github.com/falcosecurity/driverkit) | Incubating | |
| [falcosecurity/event-generator](https://github.com/falcosecurity/event-generator) | Incubating | |
| [falcosecurity/falco-exporter](https://github.com/falcosecurity/falco-exporter) | Incubating | |
| [falcosecurity/falcoctl](https://github.com/falcosecurity/falcoctl) | Incubating | |
| [falcosecurity/falcosidekick](https://github.com/falcosecurity/falcosidekick) | Incubating | |
| [falcosecurity/falcosidekick-ui](https://github.com/falcosecurity/falcosidekick-ui) | Incubating | |
| [falcosecurity/katacoda-scenarios](https://github.com/falcosecurity/katacoda-scenarios) | Incubating | |
| [falcosecurity/kernel-crawler](https://github.com/falcosecurity/kernel-crawler) | Incubating | |
| [falcosecurity/kilt](https://github.com/falcosecurity/kilt) | Incubating | |
| [falcosecurity/pdig](https://github.com/falcosecurity/pdig) | Incubating | |
| [falcosecurity/plugin-sdk-cpp](https://github.com/falcosecurity/plugin-sdk-cpp) | Incubating | |
| [falcosecurity/test-infra](https://github.com/falcosecurity/test-infra) | Incubating | The prebuilt-driver artifacts are provided on a best-effort basis. |

### Special repositories

Some repositories have a special meaning and do not fit the above statuses. These are:

| Repository | Description |
| --- | --- |
| [falcosecurity/.github](https://github.com/falcosecurity/.github) |  This repository holds default community health files, such as the [`CONTRIBUTING.md`](https://github.com/falcosecurity/.github/blob/master/CONTRIBUTING.md), [`CODE_OF_CONDUCT.md`](https://github.com/falcosecurity/.github/blob/master/CODE_OF_CONDUCT.md), and [`GOVERNANCE.md`](https://github.com/falcosecurity/.github/blob/master/GOVERNANCE.md) files.  |
| [falcosecurity/community](https://github.com/falcosecurity/community) |  Community-related stuff of The Falco Project.  |
| [falcosecurity/evolution](https://github.com/falcosecurity/evolution) |  This repository holds the evolution process of The Falco Project.  |
| [falcosecurity/template-repository](https://github.com/falcosecurity/template-repository) |  This repository holds the template for creating a new project under the [falcosecurity](https://falcosecurity) GitHub's organization.  |

### Archivied repositories

In general, a repository can be archived at the discretion of The Falco Project community. Usually, maintainers can decide to archive a project that has not been maintained for a long time or does not fit the guidelines for the projects under the [falcosecurity](https://falcosecurity) GitHub's organization anymore. In other cases, a repository is archived to reserve its name for future use.

The list of archivied repository can be found [here](https://github.com/falcosecurity?q=&type=archived&language=&sort=name).

### Retired projects

Finally, projects that are no longer maintained or relevant to The Falco Project will be retired definitively. Periodically, the maintainers clean up the [falcosecurity](https://github.com/falcosecurity) and move these projects to the [Falco Projects Retirement Home](https://github.com/falcosecurity-retire) GitHub's organization.

## Contributing

See the [contributing guide](https://github.com/falcosecurity/.github/blob/master/CONTRIBUTING.md) and the [community code of conduct](https://github.com/falcosecurity/.github/blob/master/CODE_OF_CONDUCT.md).

## Join the Community

To get involved with the evolution of The Falco Project, please visit [the community repository](https://github.com/falcosecurity/community) to find more.
