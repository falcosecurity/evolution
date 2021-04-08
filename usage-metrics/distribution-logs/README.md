# Distribution Logs

With Distribution we mean the mechanism used by Falco to distribute
release artifacts.

This folder contains a set of utilities useful to those with access
to the S3 bucket containing the logs of the CloudFront distribution.

## Why?

The purpose of this is to transform the compressed logs of the CloudFront
distribution into a csv file that can be used in Spreadsheets or scripts
to generate statistics.

The output format of this tool is:

```
driverversion,target,kernelrelease,kernelversion,drivertype,sourceip,200,2xx,3xx,404,4xx,5xx,unknown
2aa88dcf6243982697811df4c1b484bcbe9488a2,amazonlinux2,4.14.209-160.335.amzn2.x86,64,module,10.0.0.48,3,0,0,0,0,0,0
2aa88dcf6243982697811df4c1b484bcbe9488a2,ubuntu-generic,4.4.0-154-generic,181,module,10.10.10.115,1,0,0,0,0,0,0
2aa88dcf6243982697811df4c1b484bcbe9488a2,cos,4.19.150%252B,1,module,10.10.10.93,0,0,0,374,0,0,0
2aa88dcf6243982697811df4c1b484bcbe9488a2,amazonlinux2,4.14.209-160.335.amzn2.x86,64,module,10.0.0.14,3,0,0,0,0,0,0
2aa88dcf6243982697811df4c1b484bcbe9488a2,amazonlinux2,4.14.198-152.320.amzn2.x86,64,module,10.0.0.216,1,0,0,0,0,0,0
2aa88dcf6243982697811df4c1b484bcbe9488a2,amazonlinux2,4.14.209-160.335.amzn2.x86,64,module,10.0.0.177,380,0,0,0,0,0,0
```

## Limitations

This only generate statistics for driver (probe and kernel module) now.
Support for packages (deb, rpm, tgz) is not present.

## Tools Included:

### `fetch.sh`

#### Requirements:

- The [aws cli](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
- Access to the `logging-falco-distribution` bucket

A script that downloads the logs from the distribution s3 bucket and puts them under the `_output/bucket` folder in the current directory.

### `decompress.sh`

A script that takes the logs downloaded with `fetch.sh` and converts them in a single file called `_output/all.log` that can be used by `distribution-logs` for processing.

### `distribution-logs` 

This tool is written in Go, you will need to compile it prior to usage.


```
go build -o distribution-logs .
```

Using the tool is straightforward, remember to do `fetch.sh` and `decompress.sh`
first.


```
./distribution-logs > /tmp/statistics.csv
```
