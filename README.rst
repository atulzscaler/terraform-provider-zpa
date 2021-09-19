

Terraform Provider for ☁️Zscaler Private Access☁️
========================================================

**Attention:** This provider is not affiliated with, nor supported by Zscaler in any way.

Requirements
--------------
- Install - `Terraform <https://www.terraform.io/downloads.html>`_ 0.12.x/0.13.x/0.14.x/0.15.x (0.11.x or lower is incompatible)
- Install - `Go <https://golang.org/doc/install>`_ 1.16+ (This will be used to build the provider plugin.)
- Create a directory, go, follow this `doc <https://github.com/golang/go/wiki/SettingGOPATH>`_ to edit ~/.bash_profile to setup the GOPATH environment variable)

Building The Provider (Terraform v0.12+)
-------------------------------------------
- Clone repository to: `$GOPATH/src/github.com/SecurityGeekIO/terraform-provider-zpa`