# Falco artifacts on S3

This repo contains an infrastructure implementation to host Falco distribution artifacts
on Amazon S3.

Proposals related to this document:

- https://github.com/falcosecurity/falco/blob/master/proposals/20201025-drivers-storage-s3.md

TODO: we need to back the Terraform state to s3 too and write instructions on how to apply
the same way we have for the EKS clusters in [PR#198](https://github.com/falcosecurity/test-infra/pull/198/files)
