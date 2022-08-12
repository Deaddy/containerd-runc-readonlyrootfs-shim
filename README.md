# readonlyrootfs-runc-shim

A containerd, runc-based, shim for enforcing readonly rootfs at container
level, because you can not trust users.

Code base is

https://github.com/pelotech/ext-secrets-runc-shim

### Installation

1. make to obtain binary
2. put binary in $PATH of k8s nodes

3. Edit `/etc/containerd/config.toml` and replace the contents of the following section as so:

```toml
[plugins.cri.containerd.runtimes.runc]
  # This is an existing value that needs to be changed
  runtime_type = "io.containerd.runc-readonlyrootfs-shim"
```

And that's it! All pods on this node should now run via the shim. No Webhooks, no Custom Resources, no CLI commands.

## Caveats

This is a rather quick hack to fix the feature of containerd to allow
kubernetes to override readonlyrootfs and not providing a default value

## Building

To build the shim:

```sh
make
# OR
make build
```
