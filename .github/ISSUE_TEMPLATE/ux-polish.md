---
name: "UX polish"
about: "Small copy/flow tweaks that improve clarity and consistency"
title: "UX: "
labels: ["ux", "polish"]
assignees: []
---

## Summary

Briefly describe the UX polish opportunity (copy tweak, message consistency, help text, defaults, etc.).

## Current behavior

What is shown today? Include exact text/screenshots:

```
(paste exact output)
```

## Proposed behavior

What should it be instead? Provide exact replacement text and rationale.

```
(proposed text/output)
```

## Context

- Command / area:
- Why this helps:
- Severity: Low / Medium

## Acceptance criteria

- [ ] Updated message/help text matches proposal
- [ ] Tests/docs/help examples (if any) updated
- [ ] No change in behavior beyond copy/UX text

## Example (from this repo)

Current:
```
You are not logged in. Please run 'cosine login' first.
```

Proposed:
```
You are not logged in. Please run 'cos login' first.
```

Rationale: The installed binary name is `cos`, so referencing `cos login` avoids confusion for users who do not have a `cosine` alias installed.