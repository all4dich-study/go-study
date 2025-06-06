# BUILD
load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_cross_binary")

# Define the main Go binary
# Define the mymodule library
go_library(
    name = "mymodule_lib", # A unique name for this library target
    srcs = glob(["mymodule/**/*.go"]), # Assumes Go files are in a 'mymodule' directory
    importpath = "go-study/mymodule", # Matches your go.mod module path + package path
    visibility = ["//visibility:public"], # Or restrict as needed
)

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [":mymodule_lib"], # Include the library as a dependency
)

go_binary(
    name = "print_fortune",
    srcs = ["print_fortune.go"],
    deps = ["//fortune"]
)

# Platform: from https://github.com/bazel-contrib/rules_go/blob/master/go/private/platforms.bzl
# go_cross_binary reference: https://github.com/bazel-contrib/rules_go/blob/master/go/private/rules/cross.bzl#L92
go_cross_binary(
    name = "main_linux_amd64",
    platform = "@rules_go//go/toolchain:linux_arm64",
    target = ":main",
)

go_cross_binary(
    name = "print_fortune_linux_amd64",
    platform = "@rules_go//go/toolchain:linux_amd64_cgo",
    target = ":print_fortune",
)

go_cross_binary(
    name = "main_aix_ppc64",
    platform = "@rules_go//go/toolchain:aix_ppc64",
    target = ":main",
)
