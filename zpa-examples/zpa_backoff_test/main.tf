terraform {
    required_providers {
        zpa = {
            version = "1.0.0"
            source = "zscaler.com/zpa/zpa"
        }
    }
}

provider "zpa" {}


//  resource "zpa_segment_group" "example1" {
//    count = "10"
//    name = "Terraform-${count.index + 1}"
//    description = "All Other Services"
//    enabled = true
//    policy_migrated = true
//  }

 resource "zpa_application_server" "example1" {
   name = "example1"
   description = "example1"
   enabled = true
   address = "1.1.1.1"
 }
 

 resource "zpa_application_server" "example2" {
   name = "example2"
   description = "example2"
   enabled = true
   address = "2.2.2.2"
 }
 
 resource "zpa_application_server" "example3" {
   name = "example3"
   description = "example3"
   enabled = true
   address = "3.3.3.3"
 }

  resource "zpa_application_server" "example4" {
   name = "example4"
   description = "example4"
   enabled = true
   address = "4.4.4.4"
 }

  resource "zpa_application_server" "example5" {
   name = "example5"
   description = "example5"
   enabled = true
   address = "5.5.5.5"
 }

   resource "zpa_application_server" "example6" {
   name = "example6"
   description = "example6"
   enabled = true
   address = "6.6.6.6"
 }

    resource "zpa_application_server" "example7" {
   name = "example7"
   description = "example7"
   enabled = true
   address = "7.7.7.7"
 }

     resource "zpa_application_server" "example8" {
   name = "example8"
   description = "example8"
   enabled = true
   address = "8.8.8.8"
 }

      resource "zpa_application_server" "example9" {
   name = "example9"
   description = "example9"
   enabled = true
   address = "9.9.9.9"
 }

       resource "zpa_application_server" "example10" {
   name = "example10"
   description = "example10"
   enabled = true
   address = "10.10.10.10"
 }

        resource "zpa_application_server" "example11" {
   name = "example11"
   description = "example11"
   enabled = true
   address = "11.11.11.11"
 }

         resource "zpa_application_server" "example12" {
   name = "example12"
   description = "example12"
   enabled = true
   address = "12.12.12.12"
 }