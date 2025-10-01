# UX: Login hint references `cosine` instead of `cos`

labels: ux, polish
assignees: 

## Summary

When a user runs `cos` without being logged in, the CLI prints an instruction referencing `cosine login` instead of `cos login`. This can be confusing when only the `cos` binary is installed.

## Current behavior

Exact output:
```
You are not logged in. Please run 'cosine login' first.
```

## Proposed behavior

Replace the message so it references the installed binary name:

```
You are not logged in. Please run 'cos login' first.
```

If we officially support the `cosine` alias, we could say:
```
You are not logged in. Please run 'cos login' (or 'cosine login') first.
```

## Context

- Command / area: Default command path (running `cos` with no args, effectively `cos start`)
- Why this helps: Avoids user confusion when only `cos` is installed
- Severity: Low

## Acceptance criteria

- [ ] Updated message/help text matches proposal
- [ ] Any references in docs/help/examples also updated
- [ ] No change in behavior beyond the copy update