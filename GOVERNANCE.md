# The Falco Project Governance

This document describes the fundamental principles **The Falco Project** adheres to. It defines the project governance, and regulates its extension and modification.

## Project Evolution

The project source code and documentation live in various repositories under the [falcosecurity](https://github.com/falcosecurity) GitHub organization. Each repository represents a component or a subproject of The Falco Project. For transparency, project decisions are publicly made using GitHub issues and pull requests.

The Falco Project documents its evolution in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository, which is also a place designed to:
 - make decisions that regard the whole project, and
 - define rules and structures which span beyond the extent of a single repository (i.e., organization-wide).


## Repositories

The process of adding, removing, and managing repositories under the [falcosecurity](https://github.com/falcosecurity) GitHub organization is described in [REPOSITORIES.md](REPOSITORIES.md). The document assigns a scope and a status to each repository.

### Core repositories

In particular, repositories essential for building, installing, running, documenting, or using Falco are considered **core repositories** and are given the *Official* status.

### Repository ownership

[OWNERS](https://www.kubernetes.dev/docs/guide/owners/) files are used to designate responsibility over different parts of The Falco Project codebase, and serve as the implementation mechanism for the code review process. Each repository must have an `OWNERS` file in the root directory, and can optionally have other `OWNERS` files in subdirectories. The `OWNERS` file applies to everything within the directory, including the OWNERS file itself, sibling files, and child directories.

## Community

The Falco Project is driven by its community. The [falcosecurity/community](https://github.com/falcosecurity/community) repository documents the official communication channels, community calls, and other initiatives. 

The below section outlines the different roles of community members within the project, along with the responsibilities and privileges that come with them.

### Adopters 

Adopters are any organizations publicly stating that they successfully leveraged The Falco Project, or repackaged it as a component of a service offering. See the [CNCF Adopters definition](https://github.com/cncf/toc/blob/main/FAQ.md#what-is-the-definition-of-an-adopter) for more details.

Defined by:
 - The [ADOPTERS.md](https://github.com/falcosecurity/falco/blob/master/ADOPTERS.md) document.

Responsibilities:
 - Publicly announcing they are adopters of The Falco Project.

### Community Members

Community Members are all users who interact with the project. This could be through Slack, GitHub discussions, joining public project meetings, mailing lists, etc.

Responsibilities:
 - Respect the [Code Of Conduct](https://github.com/falcosecurity/.github/blob/master/CODE_OF_CONDUCT.md).

### Contributors

Contributors are [Community Members](#community-members) who [contribute](https://opensource.guide/how-to-contribute/#what-it-means-to-contribute) directly to the project and add value to it. This can be through code, documentation, taking part in bug scrubs, opening issues, proposing a pull request, etc.

Defined by:
 - Having valid contributions (as per [GitHub definition](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-github-profile/managing-contribution-settings-on-your-profile/viewing-contributions-on-your-profile#what-counts-as-a-contribution)) under the [falcosecurity](https://github.com/falcosecurity) GitHub organization.
 - Any non-GitHub contribution but considered relevant for the project, as documented in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository.

Responsibilities:
 - Respect the [Contribution Guidelines](https://github.com/falcosecurity/.github/blob/master/CONTRIBUTING.md).
 - [Sign off](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt---signoff) Git Commits to certify they adhere to the [Developer Certificate of Origin (DCO)](https://developercertificate.org/).

### Reviewers

Reviewers are [Contributors](#contributors) who have technical experience in an area of the project, and are willing to help in reviewing pull requests. They are added or removed at the sole discretion of repository maintainers. They are also [members](https://github.com/orgs/falcosecurity/people) of the [falcosecurity](https://github.com/falcosecurity) GitHub organization.

Defined by:
 - The OWNERS file `reviewers` entry.

Responsibilities:
 - Review pull requests.
 - Follow the [Maintainers](#maintainers) guidelines.

### Maintainers

Maintainers are [Contributors](#contributors) who have shown significant and sustained contribution. They are highly experienced reviewers and contributors to a specific area of the project. They are also [members](https://github.com/orgs/falcosecurity/people) of the [falcosecurity](https://github.com/falcosecurity) GitHub organization.

Defined by: 
 - The OWNERS file `approvers` entry.

Requirements:
 - Active contribution and participation in one or more areas of the project.
 - Domain expertise and a good understanding of the code-base of those areas.

Responsibilities:
 - Be active and proactive in communications, lead community calls, and help other community members.
 - Monitor [official communication channels](https://github.com/falcosecurity/community). Delayed responses are  acceptable.
 - Triage GitHub issues and review pull requests (PRs).
 - Make sure that PRs are moving forward at the right pace, or closing them.
 - Participate when called upon in security releases. Although this should be a rare occurrence, if a serious vulnerability is found it may take up to several full days of work.

To become a maintainer, [Contributors](#contributors) must express interest to the existing maintainers. The full process is documented in the [MAINTAINERS-GUIDELINES.md](https://github.com/falcosecurity/evolution/blob/master/MAINTAINERS-GUIDELINES.md) file.

The list of current [maintainers](./maintainers.yaml) is automatically updated by our infra (see the [falcosecurity/test-infra](https://github.com/falcosecurity/test-infra) repository).

### Core Maintainers 

Core Maintainers are [Maintainers](#maintainers) of at least one of the [core repositories](#core-repositories).
The Core Maintainers form a team that drives the direction, values, and governance of the overall project. They also serve as an escalation point for the overall project, and anything not easily managed by the maintainers of each repository.

Defined by:
 - The OWNERS file `approvers` entry of any [core repository](#core-repositories) (only the OWNERS file in the root directory must be considered).

Requirements:
 - Same as for [Maintainers](#maintainers).

Responsibilities and privileges:
 - Overseeing the overall project health and growth.
 - Speaking on behalf of the project.
 - Maintaining the brand, mission, vision, values, and scope of the project.
 - Defining general guidelines for the project.
 - Administering the [falcosecurity](https://github.com/falcosecurity) GitHub organization.
 - Administering any assets or services owned or assigned to project.
 - Look out for issues or conflict in any area of the project.
 - Serve as the last escalation point for an issue that can't be solved by other community roles.
 - Ability to create committees and delegate powers to them.

The full process to become a Core Maintainer is documented in the [MAINTAINERS-GUIDELINES.md](https://github.com/falcosecurity/evolution/blob/master/MAINTAINERS-GUIDELINES.md) file. 

### Emeritus Maintainers

Emeritus Maintainers are former [Maintainers](#maintainers) of a specific project area (they can still be an active maintainer of another project area). The only path to this role is to be previously listed in the `approvers` entry of an OWNERS file of that project area.

Maintainers who are domain experts over certain areas of the codebase but can no longer dedicate the time needed to handle the responsibilities of reviewing and approving changes are encouraged to add themselves in the OWNERS file `emeritus_approvers` entry.

When an emeritus maintainer returns to being more active, they may be promoted back at the discretion of the current [Maintainers](#maintainers) of the relevant area of the project.

Defined by:
- The OWNERS file `emeritus_approvers` entry.

Responsibilities:
 - Lead by example.
 - May be consulted by [Maintainers](#maintainers).
 - May serve as sponsors, and commit to the long-term success of The Falco Project.

## Conflict resolution and voting

In general, we prefer that governance issues, technical issues, and maintainer membership are amicably worked out
between the persons involved. If a dispute cannot be decided independently, the maintainers can be
called in to decide an issue.

If the [maintainers](maintainers.yaml) themselves cannot decide an issue, the issue will be resolved by **voting**.

Unless specified otherwise, a vote passes when greater than 50% + 1 of the votes are in favour.

Voting can happen through GitHub issues, GitHub pull requests, or on the CNCF Falco mailing list.

### Single project conflicts

When a conflict regards a single project (or single area of a project), each one of its maintainers can use a single vote.

The voting rights on a conflict regarding a single area of a project are given to:

- the maintainers of the specific project area (those listed in the OWNERS file of the specific area/subfolder)
- the maintainers of the whole project (those listed in the root ONWERS file)

In case the conflict is on the maintainership status it is required a majority of 66% of votes to resolve the dispute.

### Organization conflicts

When a conflict regards organization-wide issues, maintainers will use **organizational voting**.

So that no single company/organization can dominate the choices:

- Individuals not associated with or employed by a company or organization are allowed 1 organization vote
- Each company or organization receives 1 organization vote
  - Any maintainer from a company/organization may cast the vote for that organization, but it is common sense to first discuss the vote with other maintainers from the same company/organization

In case the conflict is about changing this governance document, it is required a majority of 66% of votes
to resolve the disputes.
