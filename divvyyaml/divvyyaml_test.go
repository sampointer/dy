package divvyyaml

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestCorrectParse(t *testing.T) {
	k8sGood, err := ioutil.ReadFile("../examples/k8s_deployment_output.yaml")
	if err != nil {
		t.Fatalf("Could not load test data: %v", err)
	}

	cfnGood, err := ioutil.ReadFile("../examples/cloudformation_ec2_output.yaml")
	if err != nil {
		t.Fatalf("Could not load test data: %v", err)
	}

	t.Run("k8s", func(t *testing.T) {
		compareDocuments(t, "../examples/k8s_deployment", k8sGood)
	})
	t.Run("cfn", func(t *testing.T) {
		compareDocuments(t, "../examples/cloudformation_ec2", cfnGood)
	})
}

func compareDocuments(t *testing.T, path string, good []byte) {
	var dy DivvyYaml
	err := dy.Parse(path)
	require.Nil(t, err)
	require.Equal(t, dy.Doc, string(good), "documents should match")
}
