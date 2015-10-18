# golang-thrift-minimal-example

This is a minimal working example of how to use golang and apache thrift.

It was extracted from the apache thrift tutorial; https://github.com/apache/thrift.

Go 1.5.1 and GO15VENDOREXPERIMENT=1 required, as vendor/ is used.
You will need to have the 'thrift' binary from 1.0.0-dev (current github source build) installed prior to running make && make run.

## version
Specifically the https://github.com/apache/thrift/tree/master/tutorial/go example is re-organized here
to make it easier to understand and extend.

The vendored golang thrift library code is commit fa0796d33208eadafb6f42964c8ef29d7751bfc2 on 1.0.0-dev,

The last commit there is Fri Oct 16 21:33:39 2015 +0200, from https://github.com/apache/thrift

See the vendor/git.apache.org directory for the vendored apache-thrift-golang library code. To allow the generated examples to
run no matter where you put this code, we generate into vendor/shared and vendor/tutorial. For actual use,
you'll probably want to generate into a unique directory that can be version controlled. See the -out flag to `thrift` to control where generated code goes.

LICENSE: Apache 2.0.
