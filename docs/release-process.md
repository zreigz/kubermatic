# Release

This document explains our process and requirements for creating a new Kubermatic release. 

## Testing prior releasing

This section provides a non-exhaustive minimal list of functionality that must be validated prior
to creating a new release. All these tests must be executed with every supported minor Kubernetes version.

1. Conformance tests on all supported providers (AWS, Openstack, Digitalocean, VSphere, Azure, Hetzner)
1. Cloud provider functionality `service type LoadBalancer` on all providers that support it (AWS, Openstack, Azure)
1. Cloud provider functionality `PersistentVolume` on all providers that support it (AWS, Openstack, Vsphere, Azure)
1. HorizontalPodAutoscaler

## Releasing a new version

This section covers the process to create a new Kubermatic release.

### Major|Minor release
1. Kubernetes lifecycle
    - Ensure ConfigMap for new kubelet version is added
    - Ensure RBAC for that ConfigMap is added
    - Ensure end-of-life Kubernetes has been disabled
1. Branching out
    - A release branch(example: `release/v2.7`) has been created in:
      - https://github.com/kubermatic/kubermatic
      - https://github.com/kubermatic/kubermatic-installer
    - Default branch for https://github.com/kubermatic/kubermatic-installer has been set to the new branch
1. Add a new upgrade pre-submit in infra repo
1. Ensure all provider conformance tests run on the new branch
1. Ensure upgrade tests runs and can reconcile a cluster
1. Tagging
    - Tag the release
    - Ensure images are built and pushed
    - Ensure chart sync worked
1. Documenation:
    -Update changelog (https://github.com/kubermatic/gchl)
    - Use the `ACTION REQUIRED` sections of changelog to draft a migration guide (https://docs.kubermatic.io/kubermatic/v2.12/upgrading/)
    - Have Professional Services test the upgrade and update the migration guide if necessary

### Patch release
1. After all relevant patches have been implemented in the master & release branch, a release can be created in Github https://github.com/kubermatic/kubermatic/releases. 
    - Make sure to use the release branch to create the tag
    - Drone will publish the release(Docker images & charts) 
1. After the release has been created, the changelog must be updated. (Responsible: @kdomanski) 
