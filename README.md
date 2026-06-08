[![Test | Build | Deploy](https://github.com/nitiwari-dev/engineering-lessons/actions/workflows/build-ci.yml/badge.svg)](https://github.com/nitiwari-dev/engineering-lessons/actions/workflows/build-ci.yml)

## Engineering Lessons

A polyglot mono-repo to help software engineers grow in length and breadth — TDD, DSA, language playgrounds, system design, devops, and CI/CD, each as a standalone module with its own toolchain.

> For an agent-friendly walkthrough (layout, build commands per module, conventions), see [CLAUDE.md](./CLAUDE.md).

### Quick walkthrough

| Module | What's inside |
|---|---|
| [`agile/tdd/`](./agile/tdd) | Live TDD demos in Java with JUnit/Mockito/AssertJ — auth, checkout, generic |
| [`dsa-kotlin/`](./dsa-kotlin) | Data structures & algorithms in Kotlin (Maven + JUnit5) |
| [`code/fun-with-golang/`](./code/fun-with-golang) | Go-by-example walkthrough (see `GOLANG_EXAMPLE.md` for the checklist) |
| [`code/fun-with-kotlin/`](./code/fun-with-kotlin) | Functional-style Kotlin |
| [`code/fun-with-python/`](./code/fun-with-python) | Dockerized Jupyter, GenAI labs (FFT/PEFT, ROUGE/BLEU), LLM playgrounds, CLI |
| [`code/fun-with-typescript/`](./code/fun-with-typescript) | TS basics & pre-requisites |
| [`code/fun-with-shell-script/`](./code/fun-with-shell-script) | `adb` bulk export/delete helpers |
| [`system-design/`](./system-design) | Fundamentals — hash functions, URL redirect router (Go + Java) |
| [`devops/`](./devops) | Linux & CI/CD cheatsheets |
| [`hooks/`](./hooks), [`scripts/`](./scripts) | Git hooks + helper scripts wired by `install-hooks.sh` |

### Topic checklist

- [x] agile
  - [x] tdd
- [x] code
  - [x] fun-with-golang
  - [x] fun-with-kotlin
  - [x] fun-with-python (Jupyter, GenAI labs, LLM setups)
  - [x] fun-with-typescript
  - [x] fun-with-shell-script
- [x] system-design (hash functions, redirect short url)
- [ ] devops
  - [x] linux (cheatsheet)
  - [x] ci-cd (cheatsheet)
- [ ] dsa-kotlin
  - [ ] easy
  - [ ] medium
  - [ ] hard
- [x] hooks — git hooks utility to execute checks before commit/push
- [x] scripts
  - [x] install-hooks.sh
  - [x] run-branch-name-check.sh
  - [x] run-tests.sh

### Install pre-commit + pre-push hooks

Runs `mvn test` for `agile/tdd` and `dsa-kotlin` before every commit, and validates branch name before every push.

```sh
sh scripts/install-hooks.sh
```

### Branch naming rule

Branches must start with one of:

`HEAD | feature | hotfix | conflict | bumpversion | revert | bug | fix | release | doc`

Examples: `feature-go`, `feature/go-channels`, `fix-build-ci`, `doc/readme`.

### Running tests

```sh
# Aggregate (tdd + dsa, same as pre-commit hook)
sh scripts/run-tests.sh

# Per-module
mvn test --file agile/tdd/pom.xml
mvn test --file dsa-kotlin/pom.xml
mvn test --file code/fun-with-kotlin/pom.xml
mvn test --file system-design/topics/pom.xml

cd code/fun-with-golang && go test ./...
```

CI (`.github/workflows/build-ci.yml`) is path-filtered — it only tests modules whose files changed in the PR. Markdown-only PRs skip CI.