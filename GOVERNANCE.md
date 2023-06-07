# The Falco Project Governance

This document describes the fundamental principles **The Falco Project** adheres to. It defines the project governance, and regulates its extension and modification.

**Table of Contents**

- [Principles](#principles)
- [Project Evolution](#project-evolution)
- [Repositories](#repositories)
  - [Core repositories](#core-repositories)
  - [Repository ownership](#repository-ownership)
- [Community](#community)
  - [Adopters](#adopters)
  - [Community Members](#community-members)
  - [Contributors](#contributors)
  - [Reviewers](#reviewers)
  - [Maintainers](#maintainers)
  - [Core Maintainers](#core-maintainers)
  - [Emeritus Maintainers](#emeritus-maintainers)
- [Technical Advisory Groups](#technical-advisory-groups)
- [Decision making](#decision-making)
  - [Governance changes](#governance-changes)
    - [Editorial changes](#editorial-changes)
  - [Maintainership](#maintainership)
  - [Sensitive decisions](#sensitive-decisions)
  - [Ordinary decisions](#ordinary-decisions)
- [Voting](#voting)
  - [Consensus](#consensus)
  - [Majority vote](#majority-vote)
  - [Supermajority vote](#supermajority-vote)
- [License](#license)

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

## Principles

The Falco Project and its community adhere to the following principles:
 - **Open**: Falco is open source and open to contribution, accessible and welcoming for everyone.
 - **Respectful**: the community pledges to respect all people involved in the project.
 - **Diverse**: the project furthers the interest in the diversity of representation.
 - **Transparent**: discussions, collaboration, and decision-making are done in public.
 - **Vibrant**: evolution is better than stagnation.

The Falco Project is part of the broader [CNCF](https://www.cncf.io/) community, and adheres to its values.

## Project Evolution

The project source code and documentation live in various repositories under the [falcosecurity](https://github.com/falcosecurity) GitHub organization. Each repository represents a component or a subproject of The Falco Project. For transparency, project decisions are publicly made using GitHub issues and pull requests.

The Falco Project documents its evolution in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository, which is also a place designed to:
 - make decisions that regard the whole project, and
 - define rules and structures which span beyond the extent of a single repository (i.e., organization-wide).

## Repositories

The process of adding, removing, and managing repositories under the [falcosecurity](https://github.com/falcosecurity) GitHub organization is described in [REPOSITORIES.md](REPOSITORIES.md). The document assigns a [scope](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#scope) and a [status](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#status) to each repository.

### Core repositories

In particular, repositories essential for building, installing, running, documenting, or using Falco are considered [core repositories](https://github.com/falcosecurity/evolution#core) and are given the [*core*](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#core-scope) scope.

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
 - Respect the [Code Of Conduct](CODE_OF_CONDUCT.md).

### Contributors

Contributors are [Community Members](#community-members) who [contribute](https://opensource.guide/how-to-contribute/#what-it-means-to-contribute) directly to the project and add value to it. This can be through code, documentation, taking part in bug scrubs, opening issues, proposing a pull request, etc.

Defined by:
 - Having valid contributions (as per [GitHub definition](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-github-profile/managing-contribution-settings-on-your-profile/viewing-contributions-on-your-profile#what-counts-as-a-contribution)) under the [falcosecurity](https://github.com/falcosecurity) GitHub organization.
 - Any non-GitHub contribution but considered relevant for the project, as documented in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository.

Responsibilities:
 - Respect the [Contribution Guidelines](https://github.com/falcosecurity/.github/blob/main/CONTRIBUTING.md).
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
 - Comprehensive understanding of project governance.
 - Known to uphold the projects best interests.

Measurement of these criteria are subject to the determination of the existing maintainers as attested by the sponsoring maintainer (the person working with the interested contributor to show them the ropes).

Responsibilities:
 - Be active and proactive in communications, lead community calls, and help other community members.
 - Monitor [official communication channels](https://github.com/falcosecurity/community). Delayed responses are acceptable.
 - Triage GitHub issues and review pull requests (PRs).
 - Make sure that PRs are moving forward at the right pace, or closing them.
 - Participate when called upon in security releases. Although this should be a rare occurrence, if a serious vulnerability is found it may take up to several full days of work.
 - Mentoring, encouraging, and sponsoring new reviewers and maintainers.

To become a maintainer, [Contributors](#contributors) must express interest to the existing maintainers. The full process is documented in the [MAINTAINERS_GUIDELINES.md](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS_GUIDELINES.md) file.

The list of current [maintainers](./MAINTAINERS.md) is automatically updated by our infra (see the [falcosecurity/test-infra](https://github.com/falcosecurity/test-infra) repository).

### Core Maintainers 

Core Maintainers are [Maintainers](#maintainers) of at least one of the [core repositories](#core-repositories).
The Core Maintainers form a team that drives the direction, values, and governance of the overall project. They also serve as an escalation point for the overall project, and anything not easily managed by the maintainers of each repository.

Defined by:
 - The OWNERS file `approvers` entry of any [core repository](#core-repositories) (only the OWNERS file in the root directory must be considered).

Requirements (same as for [Maintainers](#maintainers) plus):
 - Consistent engagement in the community for at least 6 months.
 - Regarded as a trustworthy curator of contributions.

Responsibilities and privileges:
 - Overseeing the overall project health and growth.
 - Speaking on behalf of the project.
 - Maintaining the brand, mission, vision, values, and scope of the project.
 - Defining general guidelines for the project.
 - Administering the [falcosecurity](https://github.com/falcosecurity) GitHub organization.
 - Administering any assets or services owned or assigned to project.
 - Handle license and copyright issues.
 - Look out for issues or conflicts in any area of the project.
 - Serve as the last escalation point for an issue that can't be solved by other community roles.
 - Ability to create committees and delegate powers to them.

The full process to become a Core Maintainer is documented in the [MAINTAINERS_GUIDELINES.md](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS_GUIDELINES.md) file. 

### Emeritus Maintainers

Emeritus Maintainers are former [Maintainers](#maintainers) of a specific project area (they can still be active maintainers of another project area). The only path to this role is to be previously listed in the `approvers` entry of an OWNERS file of that project area.

Maintainers who are domain experts over certain areas of the codebase but can no longer dedicate the time needed to handle the responsibilities of reviewing and approving changes are encouraged to add themselves in the OWNERS file `emeritus_approvers` entry.

When an emeritus maintainer returns to being more active, they may be promoted back at the discretion of the current [Maintainers](#maintainers) of the relevant area of the project.

Defined by:
- The OWNERS file `emeritus_approvers` entry.

Responsibilities:
 - Lead by example.
 - May be consulted by [Maintainers](#maintainers).
 - May serve as sponsors, and commit to the long-term success of The Falco Project.

## Technical Advisory Groups

Technical Advisory Groups (TAGs) are groups of [Community Members](#community-members) with considerable technical experience on a specific matter of The Falco Project. TAGs have the purpose of helping foster project maturity, casting recommendations, and advancing the project concerning a particular topic.

Each TAG must have a charter specifying its name, interest topic, scope, mission, membership rules, and governance processes. TAGs must have at least one TAG chair at any given time. The TAG chairs are intended to be organizers and facilitators, responsible for the overall direction and governance of the TAG.

TAGs must respect The Falco Project's [principles](#principles), communicate in the open using [communication channels](https://github.com/falcosecurity/community), and periodically share a high-level summary of their work with the [community](#community).

Anyone can propose a new TAG. Proposals must first be submitted to the community. Once a sufficient discussion has occurred, the [Core Maintainers](#core-maintainers) consider the community feedback and will eventually approve using a majority vote.

A [repository](#repositories) is assigned to each new accepted TAG. TAG's charter must define the requirements and responsibilities for the repository's [owners](#repository-ownership), who effectively are [Maintainers](#maintainers) of a repository and therefore are subject to the same rules.

## Decision making

The Falco Project tries by default to find consensus. Using lazy consensus has to date never resulted in later disputes. However, sometimes voting is required to solve disputes, or for specific matters, as described in this section.

Our [communication channels](https://github.com/falcosecurity/community) should be used to find agreement before deciding to call a vote. Discussions can happen on any official channel. Cross-posting is recommended to give more visibility to essential topics.

### Governance changes

Material changes to this document must be discussed publicly on the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository, via GitHub issues or pull requests. Any change requires a [supermajority vote](#supermajority-vote) of [Core Maintainers](#core-maintainers).

The open nature of The Falco Project, its first [principle](#principles), will never be a matter of change.

#### Editorial changes

Editorial changes are changes that fix spelling or grammar, update links, or similar; they update the style, or keep the document up to date with obvious external changes. They do not change the intent or meaning of anything in this document. Such changes must be made via pull request, and are accepted by [lazy consensus](#consensus).

### Maintainership

A [Maintainer](#maintainers) may resign by notifying their willing using a GitHub issue or pull request. In such a case, they can be moved to [Emeritus Maintainers](#emeritus-maintainers) using [lazy consensus](#consensus).

[Maintainers](#maintainers) can be added (or moved to or from [Emeritus Maintainers](#emeritus-maintainers)) with a [majority vote](#majority-vote), if the criteria in this document and the [MAINTAINERS_GUIDELINES.md](https://github.com/falcosecurity/evolution/blob/main/MAINTAINERS_GUIDELINES.md) file are met. If inactivity is the criteria, at least the past six months must be considered.

In all other cases, to remove [Maintainers](#maintainers) from their role, a [supermajority vote](#supermajority-vote) is required.

If the decision regards a [Core Maintainer](#core-maintainers) appointment or removal, any [Core Maintainer](#core-maintainers) can request to escalate the decision to all [Core Maintainers](#core-maintainers).

In any case, the persons in question are not eligible to vote, and do not count towards the quorum.

### Sensitive decisions

Any sensitive matter that needs a decision to be taken privately, including but not limited to security disclosure or financial matters, may be discussed and voted on secretly if [Core Maintainers](#core-maintainers) are present and agree.

In such situations, [Core Maintainers](#core-maintainers) must be notified promptly using the CNCF mailing list for Falco maintainers: [cncf-falco-maintainers@lists.cncf.io](mailto:cncf-falco-maintainers@lists.cncf.io).

[Lazy consensus](#consensus) is allowed only in urgent situations that render a vote impractical.

### Ordinary decisions

Technical decisions, and decisions about any other matter, are made informally by the [Maintainers](#maintainers), and [lazy consensus](#consensus) is assumed.

A [majority vote](#majority-vote) is required only if:
- an eligible voter proposes a vote, or
- a guideline documented in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository requires a vote.

## Voting

Different voting methods are used depending on the circumstance, as laid out [above](#decision-making).

For all votes, voting must be open for **one week**. If reasonably justified, the voting period can be extended up to three weeks. The end date should be clearly stated in the call to vote. A vote may be called and closed early if enough votes have been cast, and further votes cannot change the outcome.

Unless otherwise specified in this document, the voting process must be public, and the only allowed voting mechanism is using **comments on issues or pull requests** in the [falcosecurity](https://github.com/falcosecurity) GitHub organization.

For public discussions, anyone interested is encouraged to participate and cast non-binding votes. 
 
Formal power to object or cast a binding vote is limited to **eligible voters**:

| Scope | Eligible voters |
| -------- | -------- | 
| Subdirectory | [Maintainers](#maintainers) of the directory, the parent directories, if any, and [Maintainers](#maintainers) of the whole repository |
| One or more repositories | [Maintainers](#maintainers) of those repositories |
| Any other | The [Core Maintainers](#core-maintainers) team |

Notes:
  - If the matter regards multiple repositories, maintainers from those repositories can join and make a voting together.
  - Some decisions are reserved for [Core Maintainers](#core-maintainers), such as [changes in governance](#governance-changes).

In case of vacation or prolonged absence, eligible voters may temporarily delegate their voting power to another [Maintainer](#maintainers) or [Emeritus Maintainer](#emeritus-maintainers) for up to three months. The appointed person, the initial and final dates of their assignment period must be publicly recorded.

When a decision cannot be taken within its default scope, or there are less than two eligible voters, the decision must be escalated to the parent scope.

No organization or company should be allowed more than 40% of eligible votes. If more than 40% of eligible votes are affiliated with an organization or company, they must decide who will cast votes. Affiliations must be publicly stated or acknowledged when a person is associated with or employed by an organization or company. If there are not enough organizations or companies to meet this rule, any [Maintainers](#maintainers) can request to escalate to the parent scope.

### Consensus

The default decision making mechanism for The Falco Project is [lazy consensus](https://openoffice.apache.org/docs/governance/lazyConsensus.html). This means that any decision is considered supported by all concerned persons, as long as nobody objects.

Silence on any consensus decision is an implicit agreement, and equivalent to an explicit agreement. Explicit agreement may be stated at will. Decisions may, but do not need to, be called out and put up for discussion on any [communication channels](https://github.com/falcosecurity/community) at any time and by anyone.

Consensus decisions can never override or go against the spirit of an earlier explicit vote.

If anybody raises objections, all the interested parties should work together towards a solution that all involved can accept. This solution is again subject to lazy consensus.

In case no consensus can be found, but a decision one way or the other must be made, any potential eligible voters in the scope of the decision may call a formal [majority vote](#majority-vote).

### Majority vote

Majority votes must be called explicitly. The subject must be prefixed with `vote:`. In the body, the call to vote must state the proposal being voted on. It should reference any discussion leading up to this point.

Votes may take the form of a single proposal, with the option to vote yes or no, or the form of multiple alternatives.

A vote on a single proposal is considered successful if more vote in favor than against.

If there are multiple alternatives, members may vote for one or more alternatives, or vote “no” to object to all alternatives. It is not possible to cast an “abstain” vote. A vote on multiple alternatives is considered decided in favor of one alternative if it has received the most votes in favor, and a vote from more than half of those voting. Should no alternative reach this quorum, another vote on a reduced number of options may be called separately.

### Supermajority vote

Supermajority votes must be called explicitly. The subject must be prefixed with `vote:`. In the body, the call to vote must state the proposal being voted on. It should reference any discussion leading up to this point.

Votes may take the form of a single proposal, with the option to vote yes or no, or the form of multiple alternatives.

A vote on a single proposal is considered successful if at least two thirds of those eligible vote in favor.

If there are multiple alternatives, members may vote for one or more alternatives, or vote “no” to object to all alternatives. A vote on multiple alternatives is considered decided in favor of one alternative if it has received the most votes in favor, and a vote from at least two thirds of those eligible to vote. Should no alternative reach this quorum, another vote on a reduced number of options may be called separately.

## License

Repository contents must be licensed under [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0) or adhere to the [Allowed Third-Party License Policy of the CNCF](https://github.com/cncf/foundation/blob/main/allowed-third-party-license-policy.md). 

Documentation is distributed under the [Creative Commons License version 4.0](https://creativecommons.org/licenses/by/4.0/legalcode).

In each repository, copyright notices can either be included in each contributed file or stored in designated files. Copyright notices must contain the `Copyright (C) XXXX The Falco Authors` statement (see the [CNCF Copyright Notices](https://github.com/cncf/foundation/blob/main/copyright-notices.md#copyright-notices) document), where `XXXX` is the most recent year the file was updated.

[Developer Certificate of Origin (DCO)](https://developercertificate.org/) commit sign-off is required for all new code contributions.
