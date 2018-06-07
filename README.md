# go-vanityurls-apex

This is a hard fork of [GoogleCloudPlatform/govanityurls](https://github.com/GoogleCloudPlatform/govanityurls) specifically tweaked to be deployed to AWS API Gateway / Lambda, and managed via [Apex Up](https://github.com/apex/up).

## Pre-reqs
  * GNU Make
  * [Apex Up](https://github.com/apex/up)
  * An AWS account keys for IAM user/role with [policy as described by apex docs](https://up.docs.apex.sh/#aws_credentials.iam_policy_for_up_cli)

## Installation

```
go get -u go.bourbon.stream/go-vanityurls-apex
```

## Getting Started

Follow directions at [GoogleCloudPlatform/govanityurls](https://github.com/GoogleCloudPlatform/govanityurls) for your `vanity.yaml` configuration file. Next, run `make init` which will use `up` to guide you through getting your `up` project setup.

## Why the hard-fork?

`govanityurls` contains only a main package so it can not be imported.

## TODO

  * Use S3 to store config yaml, have server pull it down and cache somewhere.
