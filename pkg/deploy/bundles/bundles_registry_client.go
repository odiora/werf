package bundles

import (
	"github.com/werf/3p-helm/pkg/chart"
	bundles_registry "github.com/werf/werf/v2/pkg/deploy/bundles/registry"
)

type BundlesRegistryClient interface {
	PullChartToCache(ref *bundles_registry.Reference) error
	LoadChart(ref *bundles_registry.Reference) (*chart.Chart, error)
	SaveChart(ch *chart.Chart, ref *bundles_registry.Reference) error
	PushChart(ref *bundles_registry.Reference) error
}
