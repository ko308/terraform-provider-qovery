---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "qovery_container Resource - terraform-provider-qovery"
subcategory: ""
description: |-
  Provides a Qovery container resource. This can be used to create and manage Qovery container registry.
---

# qovery_container (Resource)

Provides a Qovery container resource. This can be used to create and manage Qovery container registry.

## Example Usage

```terraform
resource "qovery_container" "my_container" {
  # Required
  environment_id = qovery_environment.my_environment.id
  registry_id    = qovery_container_registry.my_container_registry.id
  name           = "MyContainer"
  image_name     = "qovery-api"
  tag            = "1.0.0"

  # Optional
  entrypoint            = "/dev/api"
  state                 = "RUNNING"
  auto_preview          = "true"
  cpu                   = 500
  memory                = 512
  min_running_instances = 1
  max_running_instances = 1
  environment_variables = [
    {
      key   = "ENV_VAR_KEY"
      value = "ENV_VAR_VALUE"
    }
  ]
  secrets = [
    {
      key   = "SECRET_KEY"
      value = "SECRET_VALUE"
    }
  ]

  depends_on = [
    qovery_environment.my_environment,
    qovery_container_registry.my_container_registry
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) Id of the environment.
- `image_name` (String) Name of the container image.
- `name` (String) Name of the container.
- `registry_id` (String) Id of the registry.
- `tag` (String) Tag of the container image.

### Optional

- `arguments` (Set of String) List of arguments of this container.
- `auto_preview` (Boolean) Specify if the environment preview option is activated or not for this container.
- `cpu` (Number) CPU of the container in millicores (m) [1000m = 1 CPU].
	- Must be: `>= 250`.
	- Default: `500`.
- `entrypoint` (String) Entrypoint of the container.
- `environment_variables` (Attributes Set) List of environment variables linked to this container. (see [below for nested schema](#nestedatt--environment_variables))
- `max_running_instances` (Number) Maximum number of instances running for the container.
	- Must be: `>= -1`.
	- Default: `1`.
- `memory` (Number) RAM of the container in MB [1024MB = 1GB].
	- Must be: `>= 1`.
	- Default: `512`.
- `min_running_instances` (Number) Minimum number of instances running for the container.
	- Must be: `>= 1`.
	- Default: `1`.
- `ports` (Attributes Set) List of storages linked to this container. (see [below for nested schema](#nestedatt--ports))
- `secrets` (Attributes Set) List of secrets linked to this container. (see [below for nested schema](#nestedatt--secrets))
- `state` (String) State of the container.
	- Can be: `RUNNING`, `STOPPED`.
	- Default: `RUNNING`.
- `storage` (Attributes Set) List of storages linked to this container. (see [below for nested schema](#nestedatt--storage))

### Read-Only

- `built_in_environment_variables` (Attributes Set) List of built-in environment variables linked to this container. (see [below for nested schema](#nestedatt--built_in_environment_variables))
- `external_host` (String) The container external FQDN host [NOTE: only if your container is using a publicly accessible port].
- `id` (String) Id of the container.
- `internal_host` (String) The container internal host.

<a id="nestedatt--environment_variables"></a>
### Nested Schema for `environment_variables`

Required:

- `key` (String) Key of the environment variable.
- `value` (String) Value of the environment variable.

Read-Only:

- `id` (String) Id of the environment variable.


<a id="nestedatt--ports"></a>
### Nested Schema for `ports`

Required:

- `internal_port` (Number) Internal port of the container.
	- Must be: `>= 1` and `<= 65535`.
- `publicly_accessible` (Boolean) Specify if the port is exposed to the world or not for this container.

Optional:

- `external_port` (Number) External port of the container.
	- Required if: `ports.publicly_accessible=true`.
	- Must be: `>= 1` and `<= 65535`.
- `name` (String) Name of the port.
- `protocol` (String) Protocol used for the port of the container.
	- Can be: `HTTP`.
	- Default: `HTTP`.

Read-Only:

- `id` (String) Id of the port.


<a id="nestedatt--secrets"></a>
### Nested Schema for `secrets`

Required:

- `key` (String) Key of the secret.
- `value` (String, Sensitive) Value of the secret [NOTE: will always be empty].

Read-Only:

- `id` (String) Id of the secret.


<a id="nestedatt--storage"></a>
### Nested Schema for `storage`

Required:

- `mount_point` (String) Mount point of the storage for the container.
- `size` (Number) Size of the storage for the container in GB [1024MB = 1GB].
	- Must be: `>= 1`.
- `type` (String) Type of the storage for the container.
	- Can be: `FAST_SSD`.

Read-Only:

- `id` (String) Id of the storage.


<a id="nestedatt--built_in_environment_variables"></a>
### Nested Schema for `built_in_environment_variables`

Read-Only:

- `id` (String) Id of the environment variable.
- `key` (String) Key of the environment variable.
- `value` (String) Value of the environment variable.

## Import

Import is supported using the following syntax:

```shell
terraform import qovery_container.my_container "<container_id>"
```
