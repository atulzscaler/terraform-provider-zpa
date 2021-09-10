Terraform Provider
==================
A basic [Terraform](http://terraform.io) provider for Aviatrix. Read this [tutorial](https://docs.aviatrix.com/HowTos/tf_aviatrix_howto.html) as an alternative to the README, only if the instructions are unclear.

Requirements
------------

-	Install [Terraform](https://www.terraform.io/downloads.html) 0.12.x/0.13.x/0.14.x/0.15.x (0.11.x or lower is incompatible)
-	Install [Go](https://golang.org/doc/install) 1.16+ (This will be used to build the provider plugin.)
-	Create a directory, go, follow this [doc](https://github.com/golang/go/wiki/SettingGOPATH) to edit ~/.bash_profile to setup the GOPATH environment variable)

Building The Provider (Terraform v0.12+)
---------------------

Clone repository to: `$GOPATH/src/github.com/SecurityGeekIO/terraform-provider-zpa`

```sh
$ mkdir -p $GOPATH/src/github.com/SecurityGeekIO
$ cd $GOPATH/src/github.com/SecurityGeekIO
$ git clone https://github.com/SecurityGeekIO/terraform-provider-zpa.git
```

To clone on windows
```sh
mkdir %GOPATH%\src\github.com\SecurityGeekIO
cd %GOPATH%\src\github.com\SecurityGeekIO
git clone https://github.com/SecurityGeekIO/terraform-provider-zpa.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/SecurityGeekIO/terraform-provider-zpa
$ make fmt
$ make build
```

To build on Windows
```sh
cd %GOPATH%\src\github.com\SecurityGeekIO\terraform-provider-zpa
go fmt
go install
```

Building The Provider (Terraform v0.13+)
-----------------------

### MacOS / Linux
Run the following command:
```sh
$ make build13
```

### Windows
Run the following commands for cmd:
```sh
cd %GOPATH%\src\github.com\SecurityGeekIO\terraform-provider-zpa
go fmt
go install
xcopy "%GOPATH%\bin\terraform-provider-zpa.exe" "%APPDATA%\terraform.d\plugins\zscaler.com\zpa\zpa\1.0.0\windows_amd64\" /Y
```
Run the following commands if using powershell:
```sh
cd "$env:GOPATH\src\github.com\SecurityGeekIO\terraform-provider-zpa"
go fmt
go install
xcopy "$env:GOPATH\bin\terraform-provider-zpa.exe" "$env:APPDATA\terraform.d\plugins\zscaler.com\zpa\zpa\1.0.0\windows_amd64\" /Y
```
Using SecurityGeekIO Provider (Terraform v0.12+)
-----------------------

Activate the provider by adding the following to `~/.terraformrc` on Linux/Unix.
```sh
providers {
  "zpa" = "$GOPATH/bin/terraform-provider-zpa"
}
```
For Windows, the file should be at '%APPDATA%\terraform.rc'. Do not change $GOPATH to %GOPATH%.

In Windows, for terraform 0.11.8 and lower use the above text.

In Windows, for terraform 0.11.9 and higher use the following at '%APPDATA%\terraform.rc'
```sh
providers {
  "zpa" = "$GOPATH/bin/terraform-provider-zpa.exe"
}
```

If the rc file is not present, it should be created

Using SecurityGeekIO Provider (Terraform v0.13+)
-----------------------

For Terraform v0.13+, to use a locally built version of a provider you must add the following snippet to every module
that you want to use the provider in.

```hcl
terraform {
  required_providers {
    zpa = {
      source  = "zscaler.com/zpa/zpa"
      version = "1.0.0"
    }
  }
}
```

Examples
--------

Visit [here](https://github.com/SecurityGeekIO/terraform-provider-zpa/tree/master/website/docs/) for the complete documentation for all resources on github