load("@io_bazel_rules_go//go:def.bzl",
     "go_binary", "go_library", "cgo_library")

cgo_library(
  name = "cgo_default_library",
  srcs = [
    "bug_linux.go",
    "bug_windows.go",
  ],
)

go_library(
  name = "go_default_library",
  srcs = [
    "bug.go",
  ],
  deps = [ ":cgo_default_library" ],
)
