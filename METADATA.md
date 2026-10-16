# Project metadata and ownership

Repository `OWNERS` files define roles. `.project` holds the accepted YAML metadata;
this repository holds the human documentation. The single
[`project-sync` Go tool](https://github.com/falcosecurity/test-infra/tree/master/tools/project-sync)
connects them through reviewed PRs authored by `poiana`.

## Where to make changes

| Change | Source |
| --- | --- |
| Approvers, reviewers, emeritus roles | Relevant repository's root or nested `OWNERS`; aliases in root `OWNERS_ALIASES` |
| Repository scope or status | [`.project/config/repository.yaml`](https://github.com/falcosecurity/.project/blob/main/config/repository.yaml) |
| Repository description, GitHub settings, organization admins, purpose-specific teams | [`.project/config/org.yaml`](https://github.com/falcosecurity/.project/blob/main/config/org.yaml) |
| Display name or company | [`.project/config/affiliations.yaml`](https://github.com/falcosecurity/.project/blob/main/config/affiliations.yaml) |

Generated sections of `MAINTAINERS.md` and `README.md`, and the root `.project/maintainers.yaml`, must not be edited manually.
Repository teams include all active approvers and reviewers. The CNCF roster lists
approvers only, with one unmanaged team per repository. Only `core-maintainers` is
managed by CNCF. Its members are root approvers of `Core` repositories; nested
ownership and ownership of `Special` repositories do not confer core status.

## What happens after a merge

1. A merged `OWNERS` or `OWNERS_ALIASES` change triggers reconciliation of all active repositories.
2. `poiana` proposes a `.project` PR to update organization membership, repository teams, and the maintainer roster. New affiliation entries receive editable placeholders.
3. Once that metadata is accepted, `poiana` proposes updated documentation here and updates to root `OWNERS` in every `Special` repository when needed.
4. Merging a Special `OWNERS` PR triggers another metadata reconciliation, so its repository team follows the accepted owners.

Changes to the repository catalog also update derived metadata and the OCI trigger
catalog through PRs. Changes to affiliations or repository descriptions update the
relevant generated documentation. A daily reconciliation recovers missed events.
Repeated runs keep one PR per target and purpose and leave unchanged proposals alone.
Reviewers must approve each proposal; the tool never merges its own PRs.

## Organization membership

A valid GitHub user listed as approver, reviewer, or emeritus in any active
repository's `OWNERS` is eligible for organization membership. Removing the last
such role removes eligibility. Organization admins are the only exception: their
list is maintained manually and they remain in the organization.

An invalid active account stops reconciliation. An invalid emeritus account is
omitted from membership with a report; its historical `OWNERS` entry is preserved.
Archived repositories contribute no roles and are removed from all active metadata.
Purpose-specific teams remain manually maintained. If they retain someone who is
no longer eligible, reconciliation reports a conflict for maintainers to resolve.

Peribolos applies the accepted `.project/config/org.yaml` to GitHub, including
membership removals, using its existing change limits. Team `maintainers` are eligible
organization admins; other approvers and reviewers have the team `member` role.
These GitHub team roles are separate from approval permissions in `OWNERS`.

## Retired automation

This migration approves archiving `falcosecurity/peribolos-syncer`. Its jobs were
[removed from test-infra](https://github.com/falcosecurity/test-infra/pull/2024),
and `project-sync` replaces its role. The separate daily maintainer generator and
this repository's `utils` program are retired in the coordinated migration.
Peribolos itself remains the component that applies organization configuration.

See the [rollout instructions](https://github.com/falcosecurity/test-infra/blob/master/docs/project-sync-rollout.md)
for merge order and the explicit GitHub archival step. Removing a catalog entry
does not archive a repository.
