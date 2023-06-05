
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
   * [Core](#core)
   * [Ecosystem](#ecosystem)
   * [Infra](#infra)
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

The Falco Project applies a straightforward **adoption model** for its repositories. Each repository is given a *[scope](./REPOSITORIES.md#scope)*, which outlines its purpose, and a *[status](./REPOSITORIES.md#status)* that indicates its maturity level.

For more detailed information, please refer to the [REPOSITORIES.md](./REPOSITORIES.md) file.

In the sections that follow, we present the repositories, grouped by their *scope*.

### Core

Core repositories, as defined by Falco's [governance](https://github.com/falcosecurity/evolution/blob/main/GOVERNANCE.md#core-repositories), are critically important as they are essential for building, installing, running, documenting, and using Falco.

For more information, click on the badge below.

[![Falco Core Repository](./repos/badges/falco-core-blue.svg)](./REPOSITORIES.md#core-scope)

<!-- REPOSITORY-CORE-TABLE -->
|                                         NAME                                          |                                                                                STATUS                                                                                |                                                                                                   DESCRIPTION                                                                                                    |
|---------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [falcosecurity/charts](https://github.com/falcosecurity/charts)                       | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | Helm charts repository for Falco and its ecosystem.                                                                                                                                                              |
| [falcosecurity/deploy-kubernetes](https://github.com/falcosecurity/deploy-kubernetes) | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | Kubernetes deployment resources for Falco and its ecosystem.                                                                                                                                                     |
| [falcosecurity/falco](https://github.com/falcosecurity/falco)                         | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | Falco is a cloud native runtime security tool for Linux operating systems. It is designed to detect and alert on abnormal behavior and potential security threats in real-time.                                  |
| [falcosecurity/falco-website](https://github.com/falcosecurity/falco-website)         | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | Falco website and documentation repository.                                                                                                                                                                      |
| [falcosecurity/falcoctl](https://github.com/falcosecurity/falcoctl)                   | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | The official CLI tool for working with Falco and its ecosystem components.                                                                                                                                       |
| [falcosecurity/libs](https://github.com/falcosecurity/libs)                           | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | Foundational libraries that constitute the core of Falco's functionality, offering essential features including kernel drivers and eBPF probes.                                                                  |
| [falcosecurity/plugin-sdk-go](https://github.com/falcosecurity/plugin-sdk-go)         | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | Plugins SDK for Go that facilitates writing plugins for Falco or applications built on top of Falco's libs.                                                                                                      |
| [falcosecurity/plugins](https://github.com/falcosecurity/plugins)                     | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | Plugins serve as extensions for Falco and applications built on top of Falco's libraries. This repository contains the official registry for all Falco plugins and host plugins maintained by The Falco Project. |
| [falcosecurity/rules](https://github.com/falcosecurity/rules)                         | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable) | Official rulesets for Falco provide pre-defined detection rules for various security threats and abnormal behaviors.                                                                                             |
<!-- /REPOSITORY-CORE-TABLE -->

### Ecosystem

Ecosystem repositories extend the core project by providing optional components, including value-added features, integrations, utilities, and services that, while not essential for basic Falco functioning, enrich its utility for adopters.

For more information, click on the badge below.

[![Falco Ecosystem Repository](./repos/badges/falco-ecosystem-blue.svg)](./REPOSITORIES.md#ecosystem-scope)

<!-- REPOSITORY-ECOSYSTEM-TABLE -->
|                                            NAME                                             |                                                                                   STATUS                                                                                    |                                                            DESCRIPTION                                                            |
|---------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| [falcosecurity/client-go](https://github.com/falcosecurity/client-go)                       | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | Go client and SDK for Falco.                                                                                                      |
| [falcosecurity/contrib](https://github.com/falcosecurity/contrib)                           | [![Sandbox](https://img.shields.io/badge/status-sandbox-red?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#sandbox)             | Sandbox repository to test-drive ideas/projects/code.                                                                             |
| [falcosecurity/driverkit](https://github.com/falcosecurity/driverkit)                       | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | Kit for building Falco drivers (kernel modules or eBPF probes).                                                                   |
| [falcosecurity/event-generator](https://github.com/falcosecurity/event-generator)           | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | Testing tool to generate a variety of suspect actions that are detected by Falco rules.                                           |
| [falcosecurity/falco-exporter](https://github.com/falcosecurity/falco-exporter)             | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable)        | Prometheus Metrics Exporter for Falco output events.                                                                              |
| [falcosecurity/falco-aws-terraform](https://github.com/falcosecurity/falco-aws-terraform)   | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | Terraform Module for Falco AWS Resources.                                                                                         |
| [falcosecurity/falcosidekick](https://github.com/falcosecurity/falcosidekick)               | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable)        | Falcosidekick seamlessly integrates Falco with your ecosystem, enabling event forwarding to multiple outputs in a fan-out manner. |
| [falcosecurity/falcosidekick-ui](https://github.com/falcosecurity/falcosidekick-ui)         | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | A simple WebUI with latest events from Falco.                                                                                     |
| [falcosecurity/flycheck-falco-rules](https://github.com/falcosecurity/flycheck-falco-rules) | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | A custom checker for Falco rules files that can be loaded using the Flycheck syntax checker for GNU Emacs.                        |
| [falcosecurity/kilt](https://github.com/falcosecurity/kilt)                                 | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | Kilt is a project that defines how to inject foreign apps into containers.                                                        |
| [falcosecurity/libs-sdk-go](https://github.com/falcosecurity/libs-sdk-go)                   | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | Go SDK for Falco libs.                                                                                                            |
| [falcosecurity/plugin-sdk-cpp](https://github.com/falcosecurity/plugin-sdk-cpp)             | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | Falco plugins SDK for C++.                                                                                                        |
<!-- /REPOSITORY-ECOSYSTEM-TABLE -->

### Infra

Infra repositories, such as the prominent [test-infra](https://github.com/falcosecurity/test-infra), underpin The Falco Project's infrastructure, serving the project's functioning, management, and maintenance.

For more information, click on the badge below.

[![Falco Infra Repository](./repos/badges/falco-infra-blue.svg)](./REPOSITORIES.md#infra-scope)
<!-- REPOSITORY-INFRA-TABLE -->
|                                       NAME                                        |                                                                                   STATUS                                                                                    |                             DESCRIPTION                             |
|-----------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------|
| [falcosecurity/kernel-crawler](https://github.com/falcosecurity/kernel-crawler)   | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | A tool to crawl Linux kernel versions.                              |
| [falcosecurity/pigeon](https://github.com/falcosecurity/pigeon)                   | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | Secrets and config manager for Falco's infrastructure.              |
| [falcosecurity/test-infra](https://github.com/falcosecurity/test-infra)           | [![Stable](https://img.shields.io/badge/status-stable-brightgreen?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#stable)        | Test infrastructure and automation workflows for The Falco Project. |
| [falcosecurity/testing](https://github.com/falcosecurity/testing)                 | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | All-purpose test suite for Falco and its ecosystem.                 |
| [falcosecurity/syscalls-bumper](https://github.com/falcosecurity/syscalls-bumper) | [![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating) | A tool to automatically update supported syscalls in libs.          |
<!-- /REPOSITORY-INFRA-TABLE -->

### Special

Finally, some repositories have a special meaning and do not fit the above scopes. They serve a particular purpose or function in the [falcosecurity](https://github.com/falcosecurity) organization and are curated by [core maintainers](./GOVERNANCE.md#core-maintainers). 

See [REPOSITORIES.md#special-scope](./REPOSITORIES.md#special-scope) for more information.

<!-- REPOSITORY-SPECIAL-TABLE -->
|                                 NAME                                  | STATUS |                                                    DESCRIPTION                                                    |
|-----------------------------------------------------------------------|--------|-------------------------------------------------------------------------------------------------------------------|
| [falcosecurity/.github](https://github.com/falcosecurity/.github)     | *n/a*  | Default files for all repos in the Falcosecurity GitHub org.                                                      |
| [falcosecurity/community](https://github.com/falcosecurity/community) | *n/a*  | Falco community content and resources.                                                                            |
| [falcosecurity/evolution](https://github.com/falcosecurity/evolution) | *n/a*  | A space for the community to work together, discuss ideas, define processes, and document the evolution of Falco. |
<!-- /REPOSITORY-SPECIAL-TABLE -->

### Archived

In general, a repository can be archived at the discretion of The Falco Project community. Usually, maintainers can decide to archive a project that has not been maintained for a long time or does not fit the guidelines for the projects under the [falcosecurity](https://github.com/falcosecurity) GitHub's organization anymore. In other cases, a repository is archived to reserve its name for future use.

The list of archived repositories can be found [here](https://github.com/falcosecurity?q=&type=archived&language=&sort=name).

### Retired

Repositories that are no longer maintained or relevant to The Falco Project will be retired definitively. Periodically, the maintainers clean up the [falcosecurity](https://github.com/falcosecurity) and move these projects to the [Falco Projects Retirement Home](https://github.com/falcosecurity-retire) GitHub's organization.

## Contributing

See the [contributing guide](https://github.com/falcosecurity/.github/blob/main/CONTRIBUTING.md) and the [code of conduct](./CODE_OF_CONDUCT.md).

## Security policy

To report a security vulnerability, please follow our [security policy](https://github.com/falcosecurity/.github/blob/main/SECURITY.md).

## Join the Community

To get involved with The Falco Project, please visit [the community repository](https://github.com/falcosecurity/community) to find more.
