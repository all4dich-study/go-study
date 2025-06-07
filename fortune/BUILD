load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "fortune",
    srcs = ["fortune.go"],
    importpath = "go-study/fortune",
    visibility = ["//visibility:public"],
)

go_test(
    name = "fortune_test",
    srcs = ["fortune_test.go"],
    embed = [":fortune"],
)
