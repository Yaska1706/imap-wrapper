# Branch Protection Rules for `main`

This document outlines the required branch protection settings that must be configured in GitHub repository settings to enforce merge requirements before code reaches the `main` branch.

## Required Configuration

### 1. Require status checks to pass before merging

Configure the following required status checks:

- `Lint` — Ensure all code passes golangci-lint linter checks
- `Unit tests` — Ensure all unit tests pass
- `E2E tests` — Ensure all end-to-end tests pass
- `Merge gate check` — Final gate that verifies all three checks above have passed

**How to configure:**
1. Go to repository **Settings** → **Branches**
2. Find "Branch protection rules" → **Edit** the rule for `main`
3. Under "Require status checks to pass before merging", enable it
4. Add the status checks listed above

### 2. Require pull request reviews before merging

- **Required number of approvals**: 1 (minimum)
- **Dismiss stale pull request approvals when new commits are pushed**: Enabled
- **Require review from code owners**: Enabled (if `CODEOWNERS` file exists)

**How to configure:**
1. Under "Require a pull request before merging", enable it
2. Set "Number of approvals required" to `1`
3. Check "Dismiss stale pull request approvals when new commits are pushed"
4. Optionally check "Require review from code owners"

### 3. Require signed commits (GPG signing)

- **Require commits to be signed and verified**: Enabled

**How to configure:**
1. Under "Require signed commits", enable it
2. This requires all commits in the PR to be GPG-signed before merge is allowed

**Note:** Contributors must configure GPG signing locally:
```bash
git config --global user.signingkey <YOUR_GPG_KEY_ID>
git config --global commit.gpgsign true
```

### 4. Restrict who can push to main

- **Allow force pushes**: Disable
- **Allow deletions**: Disable

**How to configure:**
1. Under "Allow force pushes", select "Do not allow force pushes"
2. Under "Allow deletions", ensure it's disabled

## Recommended Additional Settings

### Require conversation resolution before merging

- Enforce that all review comments and suggestions must be resolved before merge is allowed

**How to configure:**
1. Under "Require conversation resolution before merging", enable it

### Require branches to be up to date before merging

- Ensures branch is up-to-date with `main` before merge

**How to configure:**
1. Under "Require branches to be up to date before merging", enable it

### Include administrators

- Apply the same rules to administrators (no bypassing)

**How to configure:**
1. Under "Include administrators", enable it

## Pull Request Workflow

Once branch protection is enabled, the workflow for merging to `main` is:

1. Create a feature branch from `main`
2. Push changes and open a pull request (PR)
3. GitHub Actions CI automatically runs:
   - **Lint check** — validates code style and quality
   - **Unit tests** — validates functionality
   - **E2E tests** — validates integration
   - **Merge gate** — ensures all above pass
4. At least 1 code review approval is required (unless disabled)
5. All commits must be GPG-signed (if enabled)
6. Once all checks pass and approval is obtained, "Squash and merge" or "Rebase and merge"

## Exemptions

To exempt specific users or roles from branch protection:

1. Go to **Settings** → **Branches** → Edit branch protection rule
2. Scroll to "Allow specified actors to bypass required pull requests"
3. Add users, teams, or roles (e.g., admins) who can bypass these rules

**Note:** This should be used sparingly, only for hotfixes or emergency deployments.

## For Contributors

Before submitting a PR:

1. Ensure your branch is up-to-date with `main`
2. Sign your commits with GPG:
   ```bash
   git commit -S -m "Your commit message"
   ```
3. Push to your branch and open a PR against `main`
4. Wait for CI to complete (usually 5-10 minutes)
5. Request a code review from a team member
6. Address any review comments
7. Once approved and all checks pass, your PR is ready to merge

## References

- [GitHub: Managing a branch protection rule](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/managing-a-branch-protection-rule)
- [GitHub: About required commit signing](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-required-commit-signing)
- [Git: Signing commits with GPG](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits)
