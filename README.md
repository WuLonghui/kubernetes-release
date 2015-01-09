# kubernetes-release

Deploy [Kubernetes](http://kubernetes.io) easily with this
[BOSH](http://docs.cloudfoundry.org/bosh/) release.

## Version
* Kubernetes: v0.7.2

## Deploy Kubernetes
### BOSH
* Install [BOSH](https://github.com/cloudfoundry/bosh)
* Deployment Kubernetes:
    * Openstack
```
$ bosh upload stemcell https://s3.amazonaws.com/bosh-jenkins-artifacts/bosh-stemcell/openstack/bosh-stemcell-2605-openstack-kvm-ubuntu-trusty-go_agent.tgz
$ cd kubernetes-release
$ bosh create release
$ bosh upload release
$ vi templates/k8s-infrastructure-openstack.yml #Change the openstack-stub, like network settings
$ ./generate_deployment_manifest openstack $(bosh status --uuid) > manifest.yml
$ bosh deployment manifest.yml
$ bosh -n deploy
```

### BOSH Lite
* Install [BOSH Lite](https://github.com/cloudfoundry/bosh-lite) and
  boot the Vagrant VM.
* Deploy Kubernetes:

```
$ bosh upload stemcell https://s3.amazonaws.com/bosh-jenkins-artifacts/bosh-stemcell/warden/bosh-stemcell-348-warden-boshlite-ubuntu-trusty-go_agent.tgz
$ cd kubernetes-release
$ bosh create release 
$ bosh upload release 
$ ./generate_deployment_manifest warden $(bosh status --uuid) > manifest.yml
$ bosh deployment manifest.yml
$ bosh -n deploy
```

## Running the Guestbook example

The release includes an errand to deploy the
[GuestBook example](https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples/guestbook).
```
$ bosh run errand guestbook-example
```

## Clean up the Kubernetes
If you want to delete all the services, replication controllers and pods, you can run
```
$ bosh run errand kubernetes-clean-up
```

##Note
Forked from [CloudCredo/kubernetes-release](https://github.com/CloudCredo/kubernetes-release) 
