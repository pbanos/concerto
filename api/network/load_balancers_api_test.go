package network

import (
	// "fmt"
	"testing"

	"github.com/ingrammicro/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewLoadBalancerServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewLoadBalancerService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetLoadBalancerList(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	GetLoadBalancerListMocked(t, loadBalancersIn)
	GetLoadBalancerListFailErrMocked(t, loadBalancersIn)
	GetLoadBalancerListFailStatusMocked(t, loadBalancersIn)
	GetLoadBalancerListFailJSONMocked(t, loadBalancersIn)
}

func TestGetLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range *loadBalancersIn {
		GetLoadBalancerMocked(t, &loadBalancerIn)
		GetLoadBalancerFailErrMocked(t, &loadBalancerIn)
		GetLoadBalancerFailStatusMocked(t, &loadBalancerIn)
		GetLoadBalancerFailJSONMocked(t, &loadBalancerIn)
	}
}

func TestCreateLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range *loadBalancersIn {
		CreateLoadBalancerMocked(t, &loadBalancerIn)
		CreateLoadBalancerFailErrMocked(t, &loadBalancerIn)
		CreateLoadBalancerFailStatusMocked(t, &loadBalancerIn)
		CreateLoadBalancerFailJSONMocked(t, &loadBalancerIn)
	}
}

func TestUpdateLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range *loadBalancersIn {
		UpdateLoadBalancerMocked(t, &loadBalancerIn)
		UpdateLoadBalancerFailErrMocked(t, &loadBalancerIn)
		UpdateLoadBalancerFailStatusMocked(t, &loadBalancerIn)
		UpdateLoadBalancerFailJSONMocked(t, &loadBalancerIn)
	}
}

func TestDeleteLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range *loadBalancersIn {
		DeleteLoadBalancerMocked(t, &loadBalancerIn)
		DeleteLoadBalancerFailErrMocked(t, &loadBalancerIn)
		DeleteLoadBalancerFailStatusMocked(t, &loadBalancerIn)
	}
}

func TestListLBNodes(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	lbNodesIn := testdata.GetLBNodeData()
	for _, loadBalancerIn := range *loadBalancersIn {
		GetLBNodeListMocked(t, lbNodesIn, loadBalancerIn.ID)
		GetLBNodeListFailErrMocked(t, lbNodesIn, loadBalancerIn.ID)
		GetLBNodeListFailStatusMocked(t, lbNodesIn, loadBalancerIn.ID)
		GetLBNodeListFailJSONMocked(t, lbNodesIn, loadBalancerIn.ID)
	}
}

func TestCreateLBNode(t *testing.T) {
	lbnsIn := testdata.GetLBNodeData()

	loadBalancersIn := testdata.GetLoadBalancerData()
	loadBalancerIn := (*loadBalancersIn)[0]

	for _, lbnIn := range *lbnsIn {
		CreateLBNodeMocked(t, &lbnIn, loadBalancerIn.ID)
		CreateLBNodeFailErrMocked(t, &lbnIn, loadBalancerIn.ID)
		CreateLBNodeFailStatusMocked(t, &lbnIn, loadBalancerIn.ID)
		CreateLBNodeFailJSONMocked(t, &lbnIn, loadBalancerIn.ID)
	}
}

func TestDeleteLBNodes(t *testing.T) {
	lbnsIn := testdata.GetLBNodeData()

	loadBalancersIn := testdata.GetLoadBalancerData()
	loadBalancerIn := (*loadBalancersIn)[0]

	for _, lbnIn := range *lbnsIn {
		DeleteLBNodeMocked(t, &lbnIn, loadBalancerIn.ID)
		DeleteLBNodeFailErrMocked(t, &lbnIn, loadBalancerIn.ID)
		DeleteLBNodeFailStatusMocked(t, &lbnIn, loadBalancerIn.ID)
	}
}
