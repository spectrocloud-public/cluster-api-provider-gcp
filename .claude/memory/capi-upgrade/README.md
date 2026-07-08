# CAPI Provider Upgrade — how this works (plain-language)

Repo- & version-agnostic tooling for the periodic CAPI upgrade of this fork. All Claude content is
under `.claude/`: `skills/` = shared logic (synced across forks), `memory/capi-upgrade/` = this
fork's data + audit trail (per-repo, never synced).

## The one-liner
> Branch from the upstream release **tag**, replay the fork's spectro commits (from the integration
> branch) onto it, classify each **PICK / SKIP / REGENERATE / NEEDS-DECISION**, and stop short of
> anything unsafe — so you approve a plan before any code changes.

## Layout
```
.claude/
├── skills/                                  [SHARED, synced]
│   ├── capi-provider-upgrade/               reconcile skill + engine
│   │   ├── SKILL.md                  playbook
│   │   ├── capi-commit-reconcile.sh         engine (plan | execute)
│   │   ├── TEMPLATE-upgrade-record.md
│   │   └── test/fixture-harness.sh          self-contained algorithm test (run on any engine change)
│   ├── capi-validate.md                     validation gates (conversion fuzz, CRD diff, in-place upgrade)
│   ├── capi-palette-integrate.md            Palette monorepo wiring
│   └── capi-maas-release.md                 MAAS maintainer variant (spectro-owned; no reconcile)
└── memory/capi-upgrade/                     [PER-REPO, committed]
    ├── capi-upgrade.conf                    UPSTREAM_URL/PROVIDER/SPECTRO_SRC + classifier inputs
    ├── capi-upgrade.md                      the ledger
    ├── skip-commits.txt                     explicit human skip decisions
    ├── known-commits.tsv                    (optional) authoritative SHA→category map
    └── upgrades/                            generated records + summary.json
```

## How it runs
```
bash .claude/skills/capi-provider-upgrade/capi-commit-reconcile.sh --mode plan   [--target vX.Y.Z] [--skip-noise]
bash .claude/skills/capi-provider-upgrade/capi-commit-reconcile.sh --mode execute [--target vX.Y.Z]
```
- **plan** = read-only (throwaway worktree) → writes `upgrades/<ts>-PLAN-*.md` + `summary.json`.
- **execute** = creates the one branch `spectro-v<target>-master`, lands decisions, re-emits the record.
- Target defaults to the latest upstream tag. Versions/contract/n-3/axis are **discovered** from
  `go.mod` + `metadata.yaml` — nothing is hardcoded.

## Decisions
`PICK` (applies cleanly) · `SKIP` (already in target / skip-listed / --skip-noise) · `REGENERATE`
(generated artifact — re-made via `make`, not cherry-picked) · `NEEDS-DECISION` (diverged/unknown/
unsafe conflict → human). The branch is **INCOMPLETE** (not mergeable) while any NEEDS-DECISION or
regen/verify remains.

## Safety
plan mutates nothing; execute only ever creates/commits `spectro-v<target>-master` — never deletes/
force-updates branches, never `git clean`, never pushes, never touches `master`; aborts on
uncommitted tracked changes.

## Verify the engine
`bash .claude/skills/capi-provider-upgrade/test/fixture-harness.sh <engine>` — 10 assertions incl.
the no-silent-drop guard. Keep it green on any engine change.

## Hand-offs
Validation → `capi-validate`. Palette wiring → `capi-palette-integrate`. MAAS → `capi-maas-release`.
PR gates / review / Luma fan-out → the pipeline layer.
