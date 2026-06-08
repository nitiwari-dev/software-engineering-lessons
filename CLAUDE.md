# CLAUDE.md

Guidance for working in this repo. Read this first.

## What this repo is

A polyglot **mono-repo of software-engineering lessons** — each top-level directory is an independent learning module with its own toolchain. There is no shared root build; you cd into a module and use its native tools.

## Layout at a glance

| Path | Module | Toolchain | Tests |
|---|---|---|---|
| `agile/tdd/` | TDD with JUnit/Mockito/AssertJ examples (auth, checkout, generic) | Java + Maven | `mvn test` |
| `dsa-kotlin/` | Data structures & algorithms in Kotlin | Kotlin + Maven | `mvn test` |
| `code/fun-with-golang/` | Go-by-example walkthrough; `app/`, `config/`, `helper/`, `intro/{basic,medium,hard}` | Go 1.17 + Viper/Zerolog | none yet |
| `code/fun-with-kotlin/` | Functional-style Kotlin | Kotlin + Maven | `mvn test` |
| `code/fun-with-python/` | Jupyter setup, GenAI labs (`src/genai/`), LLM playgrounds (`src/llms/`), CLI (`src/cli/`) | Python + Docker (Jupyter) | none |
| `code/fun-with-typescript/` | TS basics & pre-requisites | TS | none |
| `code/fun-with-shell-script/` | `adb-bulk-export.sh`, `adb-bulk-delete.sh` for Android backups | sh | n/a |
| `system-design/topics/` | System-design topic implementations | Java + Maven | `mvn test` |
| `system-design/utils/` | Go utilities — `hash_func/`, `redirect_router/` | Go | none |
| `devops/{linux,ci-cd}/` | Cheatsheets (markdown only) | n/a | n/a |
| `hooks/` | `pre-commit` (runs tests), `pre-push` (branch-name check) | sh | n/a |
| `scripts/` | `install-hooks.sh`, `run-tests.sh`, `run-branch-name-check.sh` | sh | n/a |
| `.github/workflows/build-ci.yml` | Path-filtered CI: runs `mvn test` only for changed Maven modules | GitHub Actions | n/a |

## First-time setup

```sh
sh scripts/install-hooks.sh   # points git core.hooksPath at ./hooks
```

After this, every commit runs `scripts/run-tests.sh` (executes `mvn test` in `agile/tdd` and `dsa-kotlin`) and every push validates the branch name.

## Branch-name rule (enforced by pre-push)

Regex in `scripts/run-branch-name-check.sh`:

```
^((HEAD|feature|hotfix|conflict|bumpversion|revert|bug|fix|release|doc)(/|-)[A-Za-z0-9._-]+|^main$)
```

Examples: `feature-go`, `feature/go-channels`, `fix-build-ci`, `doc/readme`. Anything else fails `git push`.

## Running tests by module

```sh
# Java/Kotlin Maven modules — from the module dir
mvn test --file agile/tdd/pom.xml
mvn test --file dsa-kotlin/pom.xml
mvn test --file code/fun-with-kotlin/pom.xml
mvn test --file system-design/topics/pom.xml

# Go modules — from the module dir
cd code/fun-with-golang && go test ./...
cd system-design/utils && go test ./...

# Aggregate (only tdd + dsa, as CI/hook does)
sh scripts/run-tests.sh
```

CI in `build-ci.yml` is path-filtered — it only runs `mvn test` for the modules whose files changed in the PR. If you add a new testable module, wire it in there too.

## Conventions to follow

- **Treat each module as standalone.** Don't introduce a root build, shared deps, or cross-module imports — the value of this repo is that each module showcases its own ecosystem cleanly.
- **Markdown-only changes skip CI** (`paths-ignore: '**/*.md'`). Useful for docs PRs.
- **Hooks ignore `.md` changes for tests too** in practice, because the Maven modules don't touch markdown — but the pre-commit always runs `mvn test` regardless; expect a ~minute on docs-only commits.
- Go module uses **Go 1.17** (`code/fun-with-golang/go.mod`). Don't bump without a reason.
- Maven Java module targets **JDK 22** (CI setup). Local builds need JDK 22+ or Maven will fail.

## Quick orientation when asked to add something

- "Add a Go example" → `code/fun-with-golang/intro/{basic,medium,hard}/` — wire into `intro/intro.go` if it dispatches. Check `GOLANG_EXAMPLE.md` for the topic checklist.
- "Add a DSA problem" → `dsa-kotlin/src/main/kotlin/` + matching test in `src/test/kotlin/`.
- "Add a TDD lesson" → `agile/tdd/src/{main,test}/` under `auth/`, `checkout/`, or `generic/`.
- "Add a system-design topic" → Java goes in `system-design/topics/`, Go utility in `system-design/utils/`.
- "Add a GenAI lab" → notebook under `code/fun-with-python/src/genai/`; bring up Jupyter via `code/fun-with-python/setup/jupyter_notebook/init.sh`.

## What not to do

- Don't run `mvn test` from the repo root — there is no root POM.
- Don't bypass hooks with `--no-verify` to skip the test gate.
- Don't commit IDE files beyond what's in `.gitignore` / `.idea/codeStyles` already.
- Don't push from branches that don't match the allowed prefix — the hook will reject them.