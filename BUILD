# BUILD
load("@rules_go//go:def.bzl", "go_binary", "go_library")

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