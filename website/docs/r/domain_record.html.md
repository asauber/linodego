---
layout: "linode"
page_title: "Linode: linode_domain_record"
sidebar_current: "docs-linode-resource-domain-record"
description: |-
  Manages a Linode Domain Record.
---

# linode\_domain

Provides a Linode Domain Record resource.  This can be used to create, modify, and delete Linodes Domain Records.
For more information, see [DNS Manager](https://www.linode.com/docs/platform/manager/dns-manager/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createDomainRecord).

## Example Usage

The following example shows how one might use this resource to configure a Domain Record attached to a Linode Domain.

```hcl
resource "linode_domain" "foobar" {
    type = "master"
    domain = "foobar.example"
    soa_email = "example@foobar.example"
}

resource "linode_domain_record" "foobar" {
    domain_id = "${linode_domain.foobar.id}"
    name = "www"
    record_type = "CNAME"
    target = "foobar.example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.

* `domain_id` - (Required) The ID of the Domain to access.  *Changing `domain_id` forces the creation of a new Linode Domain Record.*.

* `record_type` - (Required) The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `record_type` forces the creation of a new Linode Domain Record.*.

* `target` - (Required) The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
- - -

* `ttl_sec` - (Optional) 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.

* `priority` - (Optional) The priority of the target host. Lower values are preferred.

* `protocol` - (Optional) The protocol this Record's service communicates with. Only valid for SRV records.

* `service` - (Optional) The service this Record identified. Only valid for SRV records.

* `tag` - (Optional) The tag portion of a CAA record. It is invalid to set this on other record types.

* `port` - (Optional) The port this Record points to.

* `weight` - (Optional) The relative weight of this Record. Higher values are preferred.

## Attributes

This resource exports no additional attributes.

## Import

Linodes Domain Records can be imported using the Linode Domain `id` followed by the Domain Record `id` separated by a comma, e.g.

```sh
terraform import linode_domain_record.www-foobar 1234567,7654321
```

The Linode Guide, [Import Existing Infrastructure to Terraform](https://www.linode.com/docs/applications/configuration-management/import-existing-infrastructure-to-terraform/), offers resource importing examples for Domain Records and other Linode resource types.