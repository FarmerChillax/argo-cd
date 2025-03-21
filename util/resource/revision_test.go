package resource

import (
	"testing"

	. "github.com/argoproj/gitops-engine/pkg/utils/testing"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/argoproj/argo-cd/v3/test"
)

func TestGetRevision(t *testing.T) {
	type args struct {
		obj *unstructured.Unstructured
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Nil", args{}, 0},
		{"Empty", args{obj: NewPod()}, 0},
		{"Invalid", args{obj: revisionExample("deployment.kubernetes.io/revision", "garbage")}, 0},
		{"Garbage", args{obj: revisionExample("garbage.kubernetes.io/revision", "1")}, 0},
		{"Deployments", args{obj: revisionExample("deployment.kubernetes.io/revision", "1")}, 1},
		{"Rollouts", args{obj: revisionExample("rollout.argoproj.io/revision", "1")}, 1},
		{"ControllerRevision", args{obj: test.NewControllerRevision()}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetRevision(tt.args.obj), "GetRevision()")
		})
	}
}

func revisionExample(name, value string) *unstructured.Unstructured {
	pod := NewPod()
	pod.SetAnnotations(map[string]string{name: value})
	return pod
}
