package(default_visibility = ["//visibility:public"])

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_proto_library")
load("@com_github_grpc_grpc//bazel:cc_grpc_library.bzl", "cc_grpc_library")

proto_library(
    name = "api_proto",
    srcs = ["api/api.proto"],
)

cc_proto_library(
    name = "api_cc_proto",
    deps = [":api_proto"],
)

cc_grpc_library(
    name = "api_cc_grpc",
    srcs = [":api_proto"],
    grpc_only = True,
    deps = [":api_cc_proto"],
)

cc_binary(
    name = "processor",
    srcs = ["processor.cc"],
    defines = ["BAZEL_BUILD"],
    deps = [
        ":api_cc_grpc",
        # http_archive made this label available for binding
        "@com_github_grpc_grpc//:grpc++",
    ],
)