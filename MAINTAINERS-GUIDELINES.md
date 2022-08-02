# Maintainers Guidelines

Maintainership-related decisions must be taken with respect to the rules established in our [governace](GOVERNANCE.md#maintainership). This document provides guidelines for the implementation of these decisions.

[Reviewers](GOVERNANCE.md#reviewers) have no maintainers power, but behave similarly and therefore this document include also guidelines for reviewers. Although they can be added and removed at the sole discretion of the maintainers, it is recommended to use similar criteria when doing so.

Both [Maintainers](GOVERNANCE.md#maintainers) and [Reviewers](GOVERNANCE.md#reviewers) are defined by [OWNERS](REPOSITORIES.md#owners) files. Most of the processes described below involve making pull requests (PRs) to correctly make changes to those files.

## Organization membership

[Maintainers](GOVERNANCE.md#maintainers) and [Reviewers](GOVERNANCE.md#reviewers) are also [organization members](https://github.com/orgs/falcosecurity/people).

If they are not yet [organization members](https://github.com/orgs/falcosecurity/people), they should open a PR to add them to the `members` entry of [org.yaml](https://github.com/falcosecurity/test-infra/blob/master/config/org.yaml).

[Maintainers](GOVERNANCE.md#maintainers) are usually not removed from [organization members](https://github.com/orgs/falcosecurity/people) because they become [Emeritus Maintainers](GOVERNANCE.md#emeritus-maintainers), unless they request so.

Former [Reviewers](GOVERNANCE.md#reviewers) are removed from the [organization members](https://github.com/orgs/falcosecurity/people) if they are no longer listed in any [OWNERS](REPOSITORIES.md#owners) file.


## Onboarding a Reviewer

If [Community Members](GOVERNANCE.md#community-members) believe they match the criteria to become [Reviewers](GOVERNANCE.md#reviewers) of a repository (or a subdirectory) they can propose themselves by opening a PR to add themselves to the `reviewers` entry of the [OWNERS](REPOSITORIES.md#owners) of the repository (or the subdirectory). The person in question must publicly express their interest and should discuss it with the other persons listed in the OWNERS file and the community before proposing themself.

[Maintainers](GOVERNANCE.md#maintainers) will review the PR and decide.

If the decision is to grant the revier status, then the person in question must become a member of the [falcosecurity](https://github.com/falcosecurity) Github organization (see the [Organization membership](#organization-membership) section).

## Onboarding a Maintainer

If [Community Members](GOVERNANCE.md#community-members) believe they match the criteria to become [Maintainers](GOVERNANCE.md#maintainers) of a repository (or a subdirectory) they can propose themselves by opening a PR to add themselves to the `approvers` entry of the [OWNERS](REPOSITORIES.md#owners) of the repository (or the subdirectory). The person in question must publicly express their interest and should discuss it with the other persons listed in the OWNERS file and the community before proposing themself.

[Maintainers](GOVERNANCE.md#maintainers) will review the PR and decide. Before taking the decision, existing maintainers may ask the person in question to shadow them or apply for a reviewer position for a period.

If the decision is to grant the maintainer status, then the person in question must:
- Become a member of the [falcosecurity](https://github.com/falcosecurity) Github organization, see the [Organization embership](#organization-membership) section.
- Join with `*-maintainers` [GitHub team](https://docs.github.com/en/organizations/organizing-members-into-teams/about-teams) relative to the repository they became maintainer of (i.e. the [`falco-maintainers`](https://github.com/orgs/falcosecurity/teams/falco-maintainers) team for [falcosecurity/falco](https://github.com/falcosecurity/falco)); One can do so by opening a PR to change the [org.yaml](https://github.com/falcosecurity/test-infra/blob/master/config/org.yaml) file.
- Update the [`persons.json`](https://github.com/falcosecurity/test-infra/blob/master/images/update-maintainers/persons.json) list by opening a PR to add their information.

## Offboarding a Reviewer

Reviewers of a repository (or a directory) can lose their status by voluntarily stepping down for personal reasons, or due to [inactivity](#review-maintainer-activity).

In such a case, a PR is required to remove the person in question from the `reviewers` entry of the respective [OWNERS](REPOSITORIES.md#owners) file.

Furthermore, former [Reviewers](GOVERNANCE.md#reviewers) are removed from the [organization members](https://github.com/orgs/falcosecurity/people) if they are no longer listed in any [OWNERS](REPOSITORIES.md#owners) file.


## Offboarding a Maintainer

Maintainers of a repository (or a directory) can lose their status by voluntarily stepping down for personal reasons, or due to inactivity.

In such a case, a PR is required to move the person in question from the `approvers` entry to the `emeritus_approvers` entry of the respective [OWNERS](REPOSITORIES.md#owners) file. 

The person in question must be mentioned in the body of the PR. This acts as a final contact attempt so that they can provide their feedback.

Another PR is required to remove them from GitHub team defined by the [org.yaml](https://github.com/falcosecurity/test-infra/blob/master/config/org.yaml) file.

