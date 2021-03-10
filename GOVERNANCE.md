# Governance

- [Governance](#governance)
  - [Maintainers](#maintainers)
  - [Process for becoming a maintainer](#process-for-becoming-a-maintainer)
  - [Maintainer responsibilities](#maintainer-responsibilities)
  - [When does a maintainer lose maintainer status](#when-does-a-maintainer-lose-maintainer-status)
    - [How to step down](#how-to-step-down)
    - [Project inactivity](#project-inactivity)
  - [Conflict resolution and voting](#conflict-resolution-and-voting)
  - [Adding new projects to the falcosecurity GitHub organization](#adding-new-projects-to-the-falcosecurity-github-organization)

## Maintainers

The list of the current [maintainers](./maintainers.yaml) is automatically updated by the [Falco Infra](https://github.com/falcosecurity/test-infra).

## Process for becoming a maintainer

- Express interest to the existing maintainers that you or your organization is interested in becoming a
  maintainer. Becoming a maintainer generally means that you are going to be spending substantial
  time (>25%) for the foreseeable future.
- Depending on which project you want to become a maintainer for, you should have domain expertise and be extremely
  proficient in its main language. Ultimately your goal is to become a maintainer that will represent your
  organization.
  - For example, to become a `falcosecurity/falco` maintainer you will need C++ proficiency.
  - On the other hand, to become a `falcosecurity/falcosidekick` maintainer, you will need Go proficiency.
- We will expect you to start contributing increasingly complicated PRs, under the guidance
  of the existing maintainers.
- We may ask you to do some PRs from our backlog.
- As you gain experience with the code base and our standards, we will ask you to do code reviews
  for incoming PRs (i.e., all maintainers are expected to shoulder a proportional share of
  community reviews).
- Be active and proactive in communications, lead community calls and help other community members.
- After a period of approximately 2-3 months of working together and making sure we see eye to eye,
  the existing maintainers will confer and decide whether to grant maintainer status or not.
  We make no guarantees on the length of time this will take, but 2-3 months is the approximate
  goal.

## Maintainer responsibilities

- Monitor Slack (delayed response is perfectly acceptable).
- Triage GitHub issues and perform pull request reviews for other maintainers and the community.
- During GitHub issue triage, apply all applicable [labels](https://github.com/falcosecurity/falco/labels)
  to each new issue. Labels are extremely useful for future issue follow up. Which labels to apply
  is somewhat subjective so just use your best judgment.
- Make sure that ongoing PRs are moving forward at the right pace or closing them.
- Participate when called upon in the security releases. Note that although this should be a rare
  occurrence, if a serious vulnerability is found, the process may take up to several full days of
  work to implement. This reality should be taken into account when discussing time commitment
  obligations with employers.
- In general continue to be willing to spend at least 25% of one's time working on Falco (~1.25
  business days per week).

## When does a maintainer lose maintainer status

Maintainers can be removed from the projects if they require so, or due to project inactivity.

### How to step down

It's totally normal that a maintainer's life changes and they now suddenly have different life priorities.
When this happens, a maintainer is expected to make sure to help other maintainers to keep up their started work.

If this decision is final, please remember to remove yourself from the OWNERS files to avoid being flooded with requests
when you don't really want to.

This kind of decision is hard, remember to reach to another maintainer if you need support with your reasoning, the community
is always there to help!

### Project inactivity

Any existing maintainer that does not show significant activity on the project they maintain can be removed.
Periodically, maintainers review the list of maintainers and their activity during the past six months.

In case the maintainer involvement in the past six months doesn't meet the requirements in this file they will be contacted
to ask wether they want to continue being a maintainer. If they decide to step down they open a pull request to be removed from the OWNERS files.

In case the maintainer on the other hand wants to continue with the role but can't perform maintainer duties, other maintainers will open a votation to discuss the removal by following the [next section](#conflict-resolution-and-voting) process.

## Conflict resolution and voting

In general, we prefer that governance issues, technical issues, and maintainer membership are amicably worked out
between the persons involved. If a dispute cannot be decided independently, the maintainers can be
called in to decide an issue.

If the maintainers themselves cannot decide an issue, the issue will be resolved by voting.

When a conflict regards a single project, each one of its maintainers can use a single vote.

When a conflict regards organization-wide issues, each maintainer can use a single vote.

Maintainers of [offically supported projects](https://github.com/falcosecurity/evolution#official-support) can use a double vote in organization-wide issues or when in deadlocks.

In case the conflict is about changing this governance or on maintainership status it is required a majority of 66% of votes to resolve the disputes.

In all the other cases, the voting process is a simple majority in which each maintainer receives one vote.

## Adding new projects to the falcosecurity GitHub organization

New projects will be added to the falcosecurity organization via GitHub issue discussion in the [falcosecurity/evolution](https://github.com/falcosecurity/evolution) repository.

Once sufficient discussion has taken place (~3-5 business
days but depending on the volume of conversation), the maintainers will
decide whether the new project should be added.

See the [section above on voting](#conflict-resolution-and-voting) if the maintainers cannot easily decide.
