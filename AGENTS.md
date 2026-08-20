# Learning Repo Source Ownership

The user owns all production Go source.

- Do not create, delete, or modify non-test `.go` files under `src/`.
- This includes blank scaffolds, declarations, stubs, and completed code.
- `*_test.go`, documentation, and repository tooling are agent-editable.
- Review production source read-only unless the user explicitly requests an exact source edit.
