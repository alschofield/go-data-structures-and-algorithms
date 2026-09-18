set shell := ["powershell.exe", "-NoProfile", "-Command"]

go := env_var_or_default("GO", "go")

default: help

help:
  @Write-Output "Run safe default tests: just test-all"
  @Write-Output "Run one leaf's default tests: just test data-structures/linear/stacks/stack"
  @Write-Output "Run an implemented leaf's contracts: just contract data-structures/linear/stacks/stack"
  @Write-Output "Benchmark an implemented leaf: just benchmark data-structures/linear/stacks/stack"
  @Write-Output "Run race-enabled default tests: just race"
  @Write-Output "Compare saved benchmark output: just benchmark-compare before.txt after.txt"

test name="":
  if ([string]::IsNullOrWhiteSpace('{{name}}')) { Write-Error "Set NAME to a taxonomy topic."; exit 1 }; {{go}} test ./src/{{name}}; if ($LASTEXITCODE -ne 0) { {{go}} test ./...; exit $LASTEXITCODE }

test-all:
  {{go}} test ./...

contract name="":
  if ([string]::IsNullOrWhiteSpace('{{name}}')) { Write-Error "Set NAME to a taxonomy topic."; exit 1 }; {{go}} test -tags=contract ./src/{{name}}

benchmark name="":
  if ([string]::IsNullOrWhiteSpace('{{name}}')) { Write-Error "Set NAME to a taxonomy topic."; exit 1 }; {{go}} test -tags=contract -run '^$' -bench . -benchmem ./src/{{name}}

race:
  {{go}} test -race ./...

benchmark-compare old="" new="":
  if ([string]::IsNullOrWhiteSpace('{{old}}') -or [string]::IsNullOrWhiteSpace('{{new}}')) { Write-Error "Set OLD and NEW benchmark output files."; exit 1 }; benchstat {{old}} {{new}}
