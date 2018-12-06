package features

import (
	"context"

	"github.com/rancher/rio/controllers/backend/features/monitoring"
	"github.com/rancher/rio/controllers/backend/features/nfs"
	"github.com/rancher/rio/types"
	spacev1beta1 "github.com/rancher/rio/types/apis/space.cattle.io/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

type featureController struct {
	features spacev1beta1.FeatureController
}

func Register(ctx context.Context, rContext *types.Context) {
	f := featureController{
		features: rContext.Global.Feature.Interface().Controller(),
	}
	rContext.Global.Feature.Interface().AddHandler(ctx, "feature", f.sync)
}

func (f featureController) sync(key string, feature *spacev1beta1.Feature) (runtime.Object, error) {
	if key == "" || feature == nil {
		return feature, nil
	}
	switch key {
	case "nfs":
		return feature, nfs.Reconcile(feature)
	case "monitoring":
		return feature, monitoring.Reconcile(feature)
	}
	return feature, nil
}
