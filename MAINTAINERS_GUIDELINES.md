# Maintainers Guidelines

Maintainership-related decisions must be taken with respect to the rules established in our [governance](GOVERNANCE.md#maintainership). This document provides guidelines for the implementation of these decisions.

[Reviewers](GOVERNANCE.md#reviewers) have no maintainers power, but behave similarly and therefore this document includes also guidelines for reviewers.

Both [Maintainers](GOVERNANCE.md#maintainers) and [Reviewers](GOVERNANCE.md#reviewers) are defined by [OWNERS](REPOSITORIES.md#owners) files. Most of the processes described below involve making pull requests (PRs) to correctly make changes to those files.

**Table of Contents**
  - [Organization membership](#organization-membership)
  - [Onboarding a Reviewer](#onboarding-a-reviewer)
  - [Onboarding a Maintainer](#onboarding-a-maintainer)
  - [Offboarding a Reviewer](#offboarding-a-reviewer)
  - [Offboarding a Maintainer](#offboarding-a-maintainer)
  - [Review maintainers activity](#review-maintainers-activity)
    * [How inactivity is measured](#how-inactivity-is-measured)
  - [Mentoring programs](#mentoring-programs)

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

## Organization membership

[Maintainers](GOVERNANCE.md#maintainers) and [Reviewers](GOVERNANCE.md#reviewers) are also [organization members](https://github.com/orgs/falcosecurity/people).

If they are not yet [organization members](https://github.com/orgs/falcosecurity/people), they should open a PR to add them to the `members` entry of [org.yaml](https://github.com/falcosecurity/test-infra/blob/master/config/org.yaml).

[Maintainers](GOVERNANCE.md#maintainers) are usually not removed from [organization members](https://github.com/orgs/falcosecurity/people) because they become [Emeritus Maintainers](GOVERNANCE.md#emeritus-maintainers), unless they request so.

Former [Reviewers](GOVERNANCE.md#reviewers) are removed from the [organization members](https://github.com/orgs/falcosecurity/people) if they are no longer listed in any [OWNERS](REPOSITORIES.md#owners) file.


## Onboarding a Reviewer

If [Community Members](GOVERNANCE.md#community-members) believe they match the criteria to become [Reviewers](GOVERNANCE.md#reviewers) of a repository (or a subdirectory) they can propose themselves by opening a PR to add themselves to the `reviewers` entry of the [OWNERS](REPOSITORIES.md#owners) of the repository (or the subdirectory). The person in question must publicly express their interest and should discuss it with the other persons listed in the OWNERS file and the community before proposing themself.

New reviewers can also be proposed and sponsored by existing [Maintainers](GOVERNANCE.md#maintainers) and [Reviewers](GOVERNANCE.md#reviewers).

[Maintainers](GOVERNANCE.md#maintainers) will review the PR and decide.

If the decision is to grant the reviewer status, then the person in question must become a member of the [falcosecurity](https://github.com/falcosecurity) Github organization (see the [Organization membership](#organization-membership) section).

## Onboarding a Maintainer

If [Community Members](GOVERNANCE.md#community-members) believe they match the criteria to become [Maintainers](GOVERNANCE.md#maintainers) of a repository (or a subdirectory) they can propose themselves by opening a PR to add themselves to the `approvers` entry of the [OWNERS](REPOSITORIES.md#owners) of the repository (or the subdirectory). The person in question must publicly express their interest and should discuss it with the other persons listed in the OWNERS file and the community before proposing themself.

New maintainers can also be proposed and sponsored by existing [Maintainers](GOVERNANCE.md#maintainers).

[Maintainers](GOVERNANCE.md#maintainers) will review the PR and decide. Before taking the decision, existing maintainers may ask the person in question to shadow them or apply for a reviewer position for a period.

If the decision is to grant the maintainer status, then the person in question must:
- If they aren't already, become a member of the [falcosecurity](https://github.com/falcosecurity) Github organization (see the [Organization embership](#organization-membership) section).
- Join with `*-maintainers` [GitHub team](https://docs.github.com/en/organizations/organizing-members-into-teams/about-teams) relative to the repository they became maintainer of (i.e. the [`falco-maintainers`](https://github.com/orgs/falcosecurity/teams/falco-maintainers) team for [falcosecurity/falco](https://github.com/falcosecurity/falco)); One can do so by opening a PR to change the [org.yaml](https://github.com/falcosecurity/test-infra/blob/master/config/org.yaml) file.
- Update the [`people/affiliations.json`](people/affiliations.json) file by opening a PR to add their information.
- Only for first-time core maintainers:
  - go to https://maintainers.cncf.io/ and open a PR to be listed as a Falco maintainer;
  - ask to be added to the `cncf-falco-maintainers@lists.cncf.io` mailing list.

## Offboarding a Reviewer

Reviewers of a repository (or a directory) can lose their status by voluntarily stepping down for personal reasons, an extended period of inactivity, a period of failing to meet the requirements for the role, a violation of the [Code Of Conduct](CODE_OF_CONDUCT.md) and/or at the maintainers' discretion.

In such a case, a PR is required to remove the person in question from the `reviewers` entry of the respective [OWNERS](REPOSITORIES.md#owners) file. [Maintainers](GOVERNANCE.md#maintainers) will review the PR and decide.

Furthermore, former [Reviewers](GOVERNANCE.md#reviewers) are removed from the [organization members](https://github.com/orgs/falcosecurity/people) if they are no longer listed in any [OWNERS](REPOSITORIES.md#owners) file.


## Offboarding a Maintainer

Maintainers of a repository (or a directory) can lose their status by voluntarily stepping down for personal reasons, or due to [inactivity](#review-maintainers-activity).

In such a case:
- A PR is required to move the person in question from the `approvers` entry to the `emeritus_approvers` entry of the respective [OWNERS](REPOSITORIES.md#owners) file. The person in question must be mentioned in the body of the PR. This acts as a final contact attempt so that they can provide their feedback.
- Another PR is required to remove them from GitHub team defined by the [org.yaml](https://github.com/falcosecurity/test-infra/blob/master/config/org.yaml) file.
- Only for core maintainers who are losing their status:
  - go to https://maintainers.cncf.io/ and open a PR to remove them under Falco;
  - ask to remove them from the `cncf-falco-maintainers@lists.cncf.io` mailing list.

## Review maintainers activity

The [Maintainers](GOVERNANCE.md#maintainers)' activity is periodically reviewed. Any [Maintainer](GOVERNANCE.md#maintainers) that does not show significant [activity](#how-inactivity-is-measured) on the repository (or the subdirectory) they maintain can be removed from the `approvers` entry of the respective [OWNERS](REPOSITORIES.md#owners) of the repository (or the subdirectory), as described in the [Offboarding a Maintainer](#offboarding-a-maintainer) section.

[Maintanership decisions](GOVERNANCE.md#maintainership) must be made on a per-OWNERS-file basis. So, a maintainer can be inactive in a project area but still involved elsewhere.

Inactive maintainers are proposed for review by any other [Maintainer](GOVERNANCE.md#maintainers) or, whenever possible, by the automation. The review is performed by opening a PR where other maintainers of the repository (or the subdirectory) can discuss and decide. If the persons under consideration voluntarily step down, the PR can be merged by [lazy consensus](GOVERNANCE.md#consensus); otherwise, a [majority vote](GOVERNANCE.md#majority-vote) is needed.

### How inactivity is measured

[Maintainers](GOVERNANCE.md#maintainers) contributions can be measured by using the CNCF [DevStats](https://devstats.cncf.io/) project (see also [API reference](https://github.com/cncf/devstatscode/blob/master/API.md)).

An inactive person is defined as someone with less than 10 recorded contributions within the past six months.

Exceptions are allowed for vacation, sick leave, maternity and paternity leave, or planned absences. Moreover, since this method does not consider special situations and does not track all [Maintainers](GOVERNANCE.md#maintainers) duties, other exceptions can be made at the discretion of existing maintainers. In particular, the criteria can be loose and tightened as needed for [*Special*](REPOSITORIES.md#status) repositories and those with very little activity.

## Mentoring programs

The community promotes initiatives to seek new maintainers to ensure that the project grows healthy and increases the maintainers' diversity.

Existing maintainers regularly open mentoring programs for aspirant maintainers. Mentorship is the most practical way to share knowledge. The goal is to help aspirants understand the maintainer's activities and duties.

Mentoring programs may be tailored to the needs of a particular repository or area of the project. However, they must at least include:
- a mentoring period where one or more maintainers with enough experience guide the participants, who will learn the dynamics of being a maintainer by performing concrete activities;
- an evaluation process that must consider the technical merit, the participation in the community, as well as the other requirements to become a maintainer.

Whenever a new program starts, it must be announced to the community via the official [communication channels](https://github.com/falcosecurity/community).

### Core Maintainers duties

[Core Maintainers](GOVERNANCE.md#core-maintainers) as a team are responsible for maintaining the [falcosecurity](https://github.com/falcosecurity) GitHub organization; thus, they can intervene in any situation concerning their responsibility. If needed, they can ask to [become maintainers of a repository](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#core-maintainers-duties-and-privileges).

[Core Maintainers](GOVERNANCE.md#core-maintainers) who volunteer to intervene must act with the spirit of serving the entire falcosecurity organization.
