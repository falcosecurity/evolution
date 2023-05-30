# Repositories

This document describes the lifecycle of repositories under the [falcosecurity](https://github.com/falcosecurity) along with their criteria, structure, scope, and status.

**Table of Contents**

  - [Criteria](#criteria)
  - [Owners](#owners)
  - [License](#license)
  - [Status](#status)
  - [Lifecycle](#lifecycle)
    * [Addition](#addition)
    * [Change of Status](#change-of-status)
    * [Archiviation](#archiviation)
    * [Unarchiviation](#unarchiviation)
    * [Removal](#removal)
    * [Core Maintainers duties and privileges](#core-maintainers-duties-and-privileges)

**Resources**
<!-- NAVIGATION_LINKS -->
 - [Governance](https://github.com/falcosecurity/evolution/blob/main/GOVERNANCE.md)
 - [Code Of Conduct](https://github.com/falcosecurity/evolution/blob/main/CODE_OF_CONDUCT.md)
 - [Maintainers Guidelines](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS_GUIDELINES.md)
 - [Maintainers List](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS.md)
 - [Repositories Guidelines](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md)
 - [Repositories List](https://github.com/falcosecurity/evolution/blob/main/README.md#repositories)
 - [Adopters List](https://github.com/falcosecurity/falco/blob/master/ADOPTERS.md)
 - [Contributing](https://github.com/falcosecurity/.github/blob/main/CONTRIBUTING.md)
 - [Security policy](https://github.com/falcosecurity/.github/blob/main/SECURITY.md)
 - [Join the Community](https://github.com/falcosecurity/community)
<!-- /NAVIGATION_LINKS -->

## Criteria

Repositories host specific parts of The Falco Project such as core codebase components, documentation, tools, libraries, or subprojects. The [falcosecurity](https://github.com/falcosecurity) GitHub organization owns repositories exclusively related to The Falco Project or its ecosystem.

## Owners

Repositories must contain [OWNERS](https://www.kubernetes.dev/docs/guide/owners/) files following the Kubernetes specification. OWNERS files are used to designate responsibility over different parts of the repository codebase and serve as the implementation mechanism for the two-phase code review process used by each project. 

Each repository must have an `OWNERS` file in the root directory. Each sub-directory that contains a unit of independent code or content may also contain an OWNERS file. This file applies to everything within its directory, including the OWNERS file itself, sibling files, and child directories. There must be only one OWNERS file per directory.

As for the [Kubernetes specification](https://www.kubernetes.dev/docs/guide/owners/), OWNERS files are in YAML format and support a given set of keys, of which The Falco Project observes the following:

- `approvers`: Users that can `/approve` or `/lgtm` PRs. See [GOVERNANCE.md](GOVERNANCE.md#maintainers) for more details.
- `reviewers`: Users that are good candidates to `/lgtm` PRs. See [GOVERNANCE.md](GOVERNANCE.md#reviewers) for more details.
- `emeritus_approvers`: Users that used to be in `approvers` but can no longer dedicate the time needed to handle the responsibilities of reviewing and approving changes. See [GOVERNANCE.md](GOVERNANCE.md#emeritus-maintainers) for more details.

In OWNERS files, users are referenced by their GitHub usernames or aliases. In a given OWNERS file, a given user can't be listed in more than one of `approvers`, `reviewers`, and `emeritus_approvers`.

## License

Refer to [GOVERNANCE.md](GOVERNANCE.md#license).

## Scope

*Scopes* are defined by [Core Maintainers](GOVERNANCE.md#core-maintainers) following the standard [decision-making process](GOVERNANCE.md#decision-making). Every repository within the [falcosecurity](https://github.com/falcosecurity) GitHub organization must have an assigned scope, which characterizes its **role and responsibilities**. The scope of a repository is determined by maintainers and might be subject to change over time. However, only [Core Maintainers](https://github.com/falcosecurity/evolution/blob/main/GOVERNANCE.md#core-maintainers) can decide whether a specific repository should be given the [*core*](#core-scope) scope. See the [Lifecycle](#lifecycle) section for more details. 

Please note that the assigned scope pertains to the repository and might not necessarily reflect all the hosted components or the artifacts it provides. For example, a mono repo that contains different components may still be scoped as *core* if its function or some hosted components align with the *core* definition. Notable examples include the [charts](https://github.com/falcosecurity/charts) and the [plugins](https://github.com/falcosecurity/plugins) repositories.

The *scope* of each repository is tracked by the [repository.yaml](https://github.com/falcosecurity/evolution/blob/main/repositories.yaml) file, and you can find an overview of that in the [README.md](README.md) of this repository.

The descriptions for the currently defined scopes are provided in the below sub-sections.

### Core Scope

Core repositories form the heart of Falco and are critically important. Our [governance](https://github.com/falcosecurity/evolution/blob/main/GOVERNANCE.md#core-repositories) precisely defines them as:

> repositories essential for building, installing, running, documenting, or using Falco

They provide foundational code, primary libraries, crucial APIs, deployment tools, documentation, and more. They are the most important repositories in the project and are the most likely to be used (directly or indirectly) by adopters. 

### Ecosystem Scope

Repositories under the ecosystem scope are extensions of the core project. They provide optional components that may be useful to Falco and its adopters. While they may not be necessary for the basic functioning of Falco, ecosystem repositories often offer value-added features, integrations, utilities, and services that help adopters make the most out of Falco.

### Infra Scope

Repositories under the infra scope are dedicated to supporting the infrastructure of The Falco Project. They are not intended to be consumed by adopters but serve as the backbone for the functioning, management, and maintenance of the project and [falcosecurity](https://github.com/falcosecurity) GitHub organization. The most notable repository in this scope is [test-infra](https://github.com/falcosecurity/test-infra).

### Special Scope

Any other repository not matching a specific scope is included in this category. This includes repositories with a unique function or a particular purpose for The Falco Project as a whole or the [falcosecurity](https://github.com/falcosecurity) GitHub organization, including (but not limited to) the [evolution](https://github.com/falcosecurity/evolution) repository, the [community](https://github.com/falcosecurity/community) repository, the [.github](https://github.com/falcosecurity/.github) repository, forks, mirrors, and other particular ones like the [template repositories](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-template-repository).

It's worth noting that *special* repositories usually do not have a designated status. This is because they are not intended to be used by adopters and may not be subject to the same lifecycle as other repositories.

## Status

*Statuses* are defined by [Core Maintainers](GOVERNANCE.md#core-maintainers) following the standard [decision-making process](GOVERNANCE.md#decision-making). Every repository within [falcosecurity](https://github.com/falcosecurity) GitHub organization must have an assigned status, which characterizes its **maturity level**. The status of a repository is determined by maintainers and might be subject to change over time. See the [Lifecycle](#lifecycle) section for more details.

Please keep in mind that the status designation refers strictly to the repository itself and might not necessarily reflect the hosted components or the artifacts it provides. This is especially true for mono repos. In such cases, the status of a component or an artifact should be denoted in the corresponding subfolder or documentation.

The *status* of each repository is tracked by the [repository.yaml](https://github.com/falcosecurity/evolution/blob/main/repositories.yaml) file, and you can find an overview of that in the [README.md](README.md) of this repository.

The definitions for each status are outlined in the below sub-sections.

### Stable

Repositories with the status *stable* have reached a high degree of maturity and reliability and are actively curated by maintainers. The components provided by those repositories are intended to be used in production by adopters and are officially supported by The Falco Project. Significant changes or updates are typically less frequent.

### Incubating

Repositories with the status *incubating* are in an intermediate stage of maturity. These repositories may still be subject to significant changes as feedback is gathered and improvements are made. The components provided by those repositories might be used in production by adopters, but it's not recommended. Generally, those projects are not safe for mission-critical purposes. The level of support for those projects may vary case by case.

### Sandbox 

Repositories with the status *sandbox* are in the earliest stage of development. These repositories are not recommended for production use but are intended for users interested in experimenting with cutting-edge features, contributing to early-stage development, or providing feedback. 

### Archived

Repositories with the status *archived* are no longer maintained or updated and are only kept for historical purposes. They are not recommended for any use, and contributions are not accepted.

## Lifecycle

This section describes the lifecycle of the repositories hosted under the [falcosecurity](https://github.com/falcosecurity) GitHub organization.

Anyone can submit proposals regarding repositories' lifecycle by opening a GitHub issue in the
[falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository (see below sections). The [Core Maintainers](GOVERNANCE.md#core-maintainers) will take into account the community feedback and decide on the proposal.

### Addition

New projects can be contributed to the falcosecurity organization by opening a GitHub issue in the
[falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository.

If the decision is to add the proposed project, then one of the falcosecurity GitHub organization admins will assist the issue opener in transferring the repository to the falcosecurity organization and configuring it in [falcosecurity/test-infra](https://github.com/falcosecurity/test-infra). Upon addition, the repository must be reviewed to make sure it respects the [criteria](#criteria), [owners](#owners), and [license](#license) points of this document. In particular, when a repository is added, the proposed owners are reviewed as described by the [MAINTAINERS_GUIDELINES.md](MAINTAINERS_GUIDELINES.md) and eventually accepted.

When a repository is first contributed to the [falcosecurity](https://github.com/falcosecurity) GitHub organization, it is assigned a [scope](#scope) and a [status](#status).

Once the repository is added to the falcosecurity GitHub organization, it and its content will be owned and licensed by The Falco Project and will be subject to its [governance](GOVERNANCE.md).

### Change of Status or Scope

Actively maintaining a repository might cause the evolution of its maturity, scope, and involvement, in The Falco Project. In those cases, the [Maintainers](GOVERNANCE.md#maintainers) of a given repository can propose changing its status by opening a GitHub issue in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository. The [Core Maintainers](GOVERNANCE.md#core-maintainers) will take into account the community feedback and decide on the proposal.

For instance, this is the path by which projects can be promoted from *incubating* to *stable* status or demoted from *incubating* to *sandbox*.

#### Archiviation

Repositories showing little to no activity during the time span of a year can be proposed for archiviation by opening a GitHub issue in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository. [Archived repositories](https://docs.github.com/en/repositories/archiving-a-github-repository/archiving-repositories) will remain inside the falcosecurity GitHub organization but will be read-only and will not be maintained. As such, OWNERS files contained in archived repositories are not valid.

In some cases, a repository is archived to reserve its name for future use.

#### Unarchiviation

Archived repositories can be proposed for unarchiviation by opening a GitHub issue in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository. If the decision is to unarchive the repository, then it must be reviewed to make sure it respects the [criteria](#criteria), [owners](#owners), and [license](#license) points of this document. In general, the same rules as for new repositories apply. The new proposed owners are reviewed as described by the [MAINTAINERS_GUIDELINES.md](MAINTAINERS_GUIDELINES.md) and eventually accepted.

### Removal

Repositories that show little relevance, are not maintained, or no longer have a purpose inside The Falco Project, can be proposed for removal by opening a GitHub issue in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository. Removed repositories will stop being maintained and will no longer be part of the falcosecurity GitHub organization. In such a case, one of the falcosecurity GitHub organization admins will assist the issue opener in transferring the repository to the [Falco Projects Retirement Home](https://github.com/falcosecurity-retire) GitHub organization and removing its configuration from [falcosecurity/test-infra](https://github.com/falcosecurity/test-infra).

### Core Maintainers duties and privileges

Since [Core Maintainers](GOVERNANCE.md#core-maintainers) as a team are responsible for the maintenance of the falcosecurity GitHub organization, they consequently have the following reserved powers:
 - decide on assigning or removing the *core* status of a repository;
 - become maintainer of *special* repositories;
 - become maintainer (or take control over) non-functioning or abandoned repositories (i.e., repositories with less than two active maintainers or disputed ones);
 - be the last escalation point for repositories disputes.

In all the above cases, [Core Maintainers](GOVERNANCE.md#core-maintainers) who volunteer to intervene must act with the spirit of serving the entire falcosecurity organization.