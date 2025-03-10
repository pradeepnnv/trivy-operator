// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/aquasecurity/trivy-operator/pkg/generated/clientset/versioned/typed/aquasecurity/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAquasecurityV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAquasecurityV1alpha1) ClusterComplianceDetailReports(namespace string) v1alpha1.ClusterComplianceDetailReportInterface {
	return &FakeClusterComplianceDetailReports{c, namespace}
}

func (c *FakeAquasecurityV1alpha1) ClusterComplianceReports(namespace string) v1alpha1.ClusterComplianceReportInterface {
	return &FakeClusterComplianceReports{c, namespace}
}

func (c *FakeAquasecurityV1alpha1) ClusterConfigAuditReports() v1alpha1.ClusterConfigAuditReportInterface {
	return &FakeClusterConfigAuditReports{c}
}

func (c *FakeAquasecurityV1alpha1) ClusterVulnerabilityReports() v1alpha1.ClusterVulnerabilityReportInterface {
	return &FakeClusterVulnerabilityReports{c}
}

func (c *FakeAquasecurityV1alpha1) ConfigAuditReports(namespace string) v1alpha1.ConfigAuditReportInterface {
	return &FakeConfigAuditReports{c, namespace}
}

func (c *FakeAquasecurityV1alpha1) VulnerabilityReports(namespace string) v1alpha1.VulnerabilityReportInterface {
	return &FakeVulnerabilityReports{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAquasecurityV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
