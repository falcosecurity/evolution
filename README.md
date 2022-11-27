
<p><img align="right" src="./img/pidgeotto.gif"/></p>
<p></p>

# Falco Project Evolution

This repository aims to document the evolution process of The Falco Project.

It provides a space for the community to work together, discuss ideas, and document processes. It is also a place to make decisions that regard the whole [falcosecurity](https://github.com/falcosecurity) organization and define rules and structures that span beyond the extent of a single repository.

**Table of Contents**

 - [Governance](#governance)
 - [Code Of Conduct](#code-of-conduct)
 - [Maintainers](#maintainers)
   * [Maintainers Guidelines](./MAINTAINERS_GUIDELINES.md)
   * [Maintainers List](./MAINTAINERS.md)
 - [Repositories](#repositories)
   * [Repositories Guidelines](./REPOSITORIES.md)
   * [Official](#official)
   * [Incubating](#incubating)
   * [Sandbox](#sandbox)
   * [Special](#special)
   * [Archived](#archived)
   * [Retired](#retired)
 - [Contributing](#contributing)
 - [Security policy](#security-policy)
 - [Join the Community](#join-the-community)

## Governance

The Falco Project governance model is documented in the [GOVERNANCE.md](./GOVERNANCE.md) file.

## Code Of Conduct

We follow the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/main/code-of-conduct.md).

Please contact [cncf-falco-maintainers@lists.cncf.io](mailto:cncf-falco-maintainers@lists.cncf.io) 
or the Linux Foundation mediator, Mishi Choudhary [mishi@linux.com](mailto:mishi@linux.com) to report an issue.

## Maintainers

The process to become a maintainer is documented in the [MAINTAINERS_GUIDELINES.md](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS_GUIDELINES.md) file.

You can find the list of current maintainers in the [MAINTAINERS.md](./MAINTAINERS.md) file.

## Repositories

The Falco Project follows a simple **adoption model** for repositories. Each repository gets a *[status](./REPOSITORIES.md#status)* that indicates the level of adoption (ie. the maturity level) or, for particular repositories, its scope. The criteria for adoption or changing the status will remain loose and tightened as needed at the discretion of the community.

You can find more details in the [REPOSITORIES.md](./REPOSITORIES.md) file.

In the sections below, we list the repositories grouped by status.

### Official

[Core repositories](./GOVERNANCE.md#core-repositories) can ultimately achieve the highest and most coveted status within The Falco Project: "*Official*."

You can request the promotion of a repository by submitting an [issue](https://github.com/falcosecurity/evolution/issues/new?assignees=&labels=kind%2Fofficialsupport&template=officialsupport_request.md).

**List of repositories in *Official* status (core repositories)**
<!-- REPOSITORY-OFFICIAL-TABLE -->
|                                         NAME                                          |                               DESCRIPTION                                |
|---------------------------------------------------------------------------------------|--------------------------------------------------------------------------|
| [falcosecurity/charts](https://github.com/falcosecurity/charts)                       | Helm charts for running Falco with Kubernetes                            |
| [falcosecurity/deploy-kubernetes](https://github.com/falcosecurity/deploy-kubernetes) | Kubernetes deployment resources for Falco                                |
| [falcosecurity/falco](https://github.com/falcosecurity/falco)                         | Cloud Native Runtime Security                                            |
| [falcosecurity/falco-website](https://github.com/falcosecurity/falco-website)         | Hugo content to generate website content. Hosted by the CNCF             |
| [falcosecurity/libs](https://github.com/falcosecurity/libs)                           | libsinsp, libscap, the kernel module driver, and the eBPF driver sources |
| [falcosecurity/plugin-sdk-go](https://github.com/falcosecurity/plugin-sdk-go)         |                                                                          |
| [falcosecurity/plugins](https://github.com/falcosecurity/plugins)                     | Falco plugins registry                                                   |
<!-- /REPOSITORY-OFFICIAL-TABLE -->

### Incubating

The "*Incubating*" level refers to those repositories that contain non-core components or any subprojects that don't yet have an adequate level of maturity. You can request the incubation of a repository by submitting an [issue](https://github.com/falcosecurity/evolution/issues/new?assignees=&labels=kind%2Fincubating&template=incubating_request.md).

**List of repositories in *Incubating* status**
<!-- REPOSITORY-INCUBATING-TABLE -->
|                                           NAME                                            |                                DESCRIPTION                                |
|-------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| [falcosecurity/client-go](https://github.com/falcosecurity/client-go)                     | Go client and SDK for Falco                                               |
| [falcosecurity/driverkit](https://github.com/falcosecurity/driverkit)                     | Kit for building Falco drivers: kernel modules or eBPF probes             |
| [falcosecurity/event-generator](https://github.com/falcosecurity/event-generator)         | Generate a variety of suspect actions that are detected by Falco rulesets |
| [falcosecurity/falco-exporter](https://github.com/falcosecurity/falco-exporter)           | Prometheus Metrics Exporter for Falco output events                       |
| [falcosecurity/falco-aws-terraform](https://github.com/falcosecurity/falco-aws-terraform) | Terraform Module for Falco AWS Resources                                  |
| [falcosecurity/falcoctl](https://github.com/falcosecurity/falcoctl)                       | Administrative tooling for Falco                                          |
| [falcosecurity/falcosidekick](https://github.com/falcosecurity/falcosidekick)             | Connect Falco to your ecosystem                                           |
| [falcosecurity/falcosidekick-ui](https://github.com/falcosecurity/falcosidekick-ui)       | A simple WebUI with latest events from Falco                              |
| [falcosecurity/kernel-crawler](https://github.com/falcosecurity/kernel-crawler)           | A tool to crawl Linux kernel versions                                     |
| [falcosecurity/kilt](https://github.com/falcosecurity/kilt)                               | Kilt is a project that defines how to inject foreign apps into containers |
| [falcosecurity/libs-sdk-go](https://github.com/falcosecurity/libs-sdk-go)                 | Go SDK for Falco libs                                                     |
| [falcosecurity/plugin-sdk-cpp](https://github.com/falcosecurity/plugin-sdk-cpp)           | Falco plugins SDK for C++                                                 |
| [falcosecurity/test-infra](https://github.com/falcosecurity/test-infra)                   | Falco workflow & testing infrastructure                                   |
| [falcosecurity/syscalls-bumper](https://github.com/falcosecurity/syscalls-bumper)         | A tool to automatically update supported syscalls in libs                 |
<!-- /REPOSITORY-INCUBATING-TABLE -->

### Sandbox

You can find "*Sandbox*" level ideas/projects/code under the [falcosecurity/contrib](https://github.com/falcosecurity/contrib) repository.

### Special

Some repositories have a special meaning and do not fit the above statuses. They serve a particular purpose or function in the [falcosecurity](https://github.com/falcosecurity) organization and are curated by [core maintainers](./GOVERNANCE.md#core-maintainers).

**List of *Special* repositories**
<!-- REPOSITORY-SPECIAL-TABLE -->
|                                 NAME                                  |              DESCRIPTION               |
|-----------------------------------------------------------------------|----------------------------------------|
| [falcosecurity/.github](https://github.com/falcosecurity/.github)     | Default community health files         |
| [falcosecurity/community](https://github.com/falcosecurity/community) | The Falco Project Community            |
| [falcosecurity/evolution](https://github.com/falcosecurity/evolution) | Evolution process of The Falco Project |
<!-- /REPOSITORY-SPECIAL-TABLE -->

### Archived

In general, a repository can be archived at the discretion of The Falco Project community. Usually, maintainers can decide to archive a project that has not been maintained for a long time or does not fit the guidelines for the projects under the [falcosecurity](https://github.com/falcosecurity) GitHub's organization anymore. In other cases, a repository is archived to reserve its name for future use.

The list of archived repositories can be found [here](https://github.com/falcosecurity?q=&type=archived&language=&sort=name).

### Retired

Finally, repositories that are no longer maintained or relevant to The Falco Project will be retired definitively. Periodically, the maintainers clean up the [falcosecurity](https://github.com/falcosecurity) and move these projects to the [Falco Projects Retirement Home](https://github.com/falcosecurity-retire) GitHub's organization.

## Contributing

See the [contributing guide](https://github.com/falcosecurity/.github/blob/main/CONTRIBUTING.md) and the [code of conduct](./CODE_OF_CONDUCT.md).

## Security policy

To report a security vulnerability, please follow our [security policy](https://github.com/falcosecurity/.github/blob/main/SECURITY.md).

## Join the Community

To get involved with The Falco Project, please visit [the community repository](https://github.com/falcosecurity/community) to find more.
