git_repository(
    name = "io_bazel_rules_go",
    remote = "git://github.com/bazelbuild/rules_go.git",
    commit = "537a37a3c1c6ca06eedc6cac3f24e92f8519f092",
)
load("@io_bazel_rules_go//go:def.bzl",
     "go_repositories",
     "new_go_repository")

go_repositories()

new_go_repository(
  name = "org_golang_x_text",
  importpath = "golang.org/x/text",
  commit = "a9a820217f98f7c8a207ec1e45a874e1fe12c478",
)
