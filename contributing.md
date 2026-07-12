# Contributing guide

This guide describes how to work on the DFS course project.

---

## Getting started
1. Create a local clone of the repository.
2. Install prerequisites (`go 1.24`, `docker`, `protoc`).
3. Run the test suite before and after making changes.

---

## Helpful commands for getting started (run from `dfs` directory)
| Step | Command | Explanation |
|------|---------| ----------------- |
| Format | `make fmt` | formats your code properly | 
| Dev setup | `make dev-setup` | installs protoc |
| Clean proto | `make clean` | cleans protobuf files in pkg/proto |
| Generate proto | `make proto` | regenerates protobuf files according to definitions|
| Unit tests | `make test` | runs all unit tests |
| e2e tests | `make e2e` | runs all e2e tests (requires docker engine running) |

---

## Commits
* Use the conventional‐commits style: `scope: summary` – e.g. `feat: new feature/enchancement`.  
* Reference an issue when applicable: `Fixes #42`.
* Sign each commit.

---

## Pull requests
1. Keep each change focused on one fix or feature.
2. Update or add tests when behavior changes.
3. Update documentation (README or docs/*) when behavior changes.
4. Explain why the change is needed in the submission notes.
5. Run all tests before submitting.

---

## Code style
* Go 1.24, standard formatting (`go fmt`, `goimports`).
* Log using `pkg/logging` helpers; avoid `fmt.Println` in production code.
* Make sure functions are small and testable. 
* Dependency injection pattern is your friend.
* Error handling: refer to [Protocol Error Handling Documentation](docs/protocol.md#error-handling-and-recovery)

---

## License / DCO
By contributing you agree that your code will be released under the MIT license and that you have the right to do so. 
