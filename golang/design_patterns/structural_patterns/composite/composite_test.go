package composite_test

import (
	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/structural_patterns/composite/composite"
)

func Example_composite() {
	kubernetes := composite.NewFolder("/etc/kubernetes")

	kubernetes.Add(composite.NewFile("admin.conf"))
	kubernetes.Add(composite.NewFile("controller-manager.conf"))
	kubernetes.Add(composite.NewFile("kubelet.conf"))

	pki := composite.NewFolder("/pki")
	pki.Add(composite.NewFile("apiserver-etcd-client.crt"))
	pki.Add(composite.NewFile("apiserver-etcd-client.key"))
	pki.Add(composite.NewFile("ca.crt"))
	pki.Add(composite.NewFile("ca.key"))
	kubernetes.Add(pki)

	kubernetes.Search("kubelet")

	// Output:
	// Serching recursively for keyword kubelet in folder /etc/kubernetes
	// Searching for keyword kubelet in file admin.conf
	// Searching for keyword kubelet in file controller-manager.conf
	// Searching for keyword kubelet in file kubelet.conf
	// Serching recursively for keyword kubelet in folder /pki
	// Searching for keyword kubelet in file apiserver-etcd-client.crt
	// Searching for keyword kubelet in file apiserver-etcd-client.key
	// Searching for keyword kubelet in file ca.crt
	// Searching for keyword kubelet in file ca.key
}
