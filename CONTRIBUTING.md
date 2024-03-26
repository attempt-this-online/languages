# Contributing

For general instructions, see [attempt-this-online/docs/contributing.md](https://github.com/attempt-this-online/attempt-this-online/blob/main/docs/contributing.md).

## Continuous integration

* [pre-commit.ci](https://pre-commit.ci) will check some things in your pull request.
  If the check fails, don't worry - pre-commit.ci can **fix it for you**.
  It will automatically commit the change to your pull request branch.

  * The `.images.gitlab-ci.yml` file is automatically generated.
    If you didn't run the script to automatically regenerate it,
    pre-commit.ci will do that for you

* Your new/changed image will also be built by [GitHub Actions](https://github.com/attempt-this-online/languages/actions).
  Any other images which depend on it will also be rebuilt.
  The builds **need to succeed** before your pull request can be merged.

* Every Saturday morning (UTC), [GitLab](https://gitlab.pxeger.com/attempt-this-online/languages/-/pipelines)
  will do a clean rebuild of every image in the repository.
  This is what the `.images.gitlab-ci.yml` file is for.
  The built images will be pushed to `registry.gitlab.pxeger.com`.
