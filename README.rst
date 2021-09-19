

Terraform Provider for ☁️Zscaler Private Access☁️
========================================================

**Attention:** This provider is not affiliated with, nor supported by Zscaler in any way.

Requirements
--------------
- Install - `Terraform <https://www.terraform.io/downloads.html>`_ 0.12.x/0.13.x/0.14.x/0.15.x (0.11.x or lower is incompatible)
- Install - `Go <https://golang.org/doc/install>`_ 1.16+ (This will be used to build the provider plugin.)
- Create a directory, go, follow this `doc <https://github.com/golang/go/wiki/SettingGOPATH>`_ to edit ~/.bash_profile to setup the GOPATH environment variable)

Building The Provider (Terraform v0.12+)
==========================================
- Clone repository to: `$GOPATH/src/github.com/SecurityGeekIO/terraform-provider-zpa`

.. code-block:: console

 $ mkdir -p $GOPATH/src/github.com/terraform-providers
 $ cd $GOPATH/src/github.com/terraform-providers
 $ git clone https://github.com/SecurityGeekIO/terraform-provider-zpa.git

To clone on windows
.. code-block:: console

 $ mkdir %GOPATH%\src\github.com\terraform-providers
 $ cd %GOPATH%\src\github.com\terraform-providers
 $ git clone https://github.com/SecurityGeekIO/terraform-provider-zpa.git

 Enter the provider directory and build the provider

.. code-block:: console

 $ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-zpa
 $ make fmt
 $ make build

 To build on Windows
 --------------------

.. code-block:: console

 $ cd %GOPATH%\src\github.com\terraform-providers\terraform-provider-zpa
 $ go fmt
 $ go install

 Building The Provider (Terraform v0.13+)
==========================================

MacOS / Linux
--------------
Run the following command:

.. code-block:: console

$ make build13