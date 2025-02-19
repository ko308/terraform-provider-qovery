---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qovery_container_registry Resource - terraform-provider-qovery"
subcategory: ""
description: |-
  Provides a Qovery container registry resource. This can be used to create and manage Qovery container registry.
---

# qovery_container_registry (Resource)

Provides a Qovery container registry resource. This can be used to create and manage Qovery container registry.

## Example Usage

```terraform
resource "qovery_container_registry" "my_container_registry" {
  # Required
  organization_id = qovery_organization.my_organization.id
  name            = "my_aws_creds"
  kind            = "DOCKER_HUB"
  url             = "https://docker.io"
  config = {
    username = "<my_username>"
    password = "<my_password>"
  }

  # Optional
  description = "My Docker Hub Registry"

  depends_on = [
    qovery_organization.my_organization
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `kind` (String) Kind of the container registry.
	- Can be: `DOCKER_HUB`, `DOCR`, `ECR`, `PUBLIC_ECR`, `SCALEWAY_CR`.
- `name` (String) Name of the container registry.
- `organization_id` (String) Id of the organization.
- `url` (String) URL of the container registry.

### Optional

- `config` (Attributes) Configuration needed to authenticate the container registry. (see [below for nested schema](#nestedatt--config))
- `description` (String) Description of the container registry.

### Read-Only

- `id` (String) Id of the container registry.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `access_key_id` (String) Required if kind is `ECR` or `PUBLIC_ECR`.
- `password` (String) Required if kind is `DOCKER_HUB`.
- `region` (String) Required if kind is `ECR` or `SCALEWAY_CR`.
- `scaleway_access_key` (String) Required if kind is `SCALEWAY_CR`.
- `scaleway_secret_key` (String) Required if kind is `SCALEWAY_CR`.
- `secret_access_key` (String) Required if kind is `ECR` or `PUBLIC_ECR`.
- `username` (String) Required if kind is `DOCKER_HUB`.

## Import

Import is supported using the following syntax:

```shell
terraform import qovery_container_registry.my_container_registry "<organization_id>,<container_registry_id>"
```
