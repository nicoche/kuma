package service

import (
	"context"
	"time"

	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_cache "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	envoy_xds "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/kumahq/kuma/pkg/core"
	core_manager "github.com/kumahq/kuma/pkg/core/resources/manager"
	mads_generator "github.com/kumahq/kuma/pkg/mads/v1/generator"
	mads_reconcile "github.com/kumahq/kuma/pkg/mads/v1/reconcile"
	util_watchdog "github.com/kumahq/kuma/pkg/util/watchdog"
	util_xds "github.com/kumahq/kuma/pkg/util/xds"
	util_xds_v3 "github.com/kumahq/kuma/pkg/util/xds/v3"
	"github.com/kumahq/kuma/pkg/xds/cache/mesh"
)

func NewSnapshotGenerator(rm core_manager.ReadOnlyResourceManager, meshCache *mesh.Cache) util_xds_v3.SnapshotGenerator {
	return mads_reconcile.NewSnapshotGenerator(rm, mads_generator.MonitoringAssignmentsGenerator{}, meshCache)
}

func NewVersioner() util_xds_v3.SnapshotVersioner {
	return util_xds_v3.SnapshotAutoVersioner{UUID: core.NewUUID}
}

func NewReconciler(hasher envoy_cache.NodeHash, cache util_xds_v3.SnapshotCache,
	generator util_xds_v3.SnapshotGenerator, versioner util_xds_v3.SnapshotVersioner,
) mads_reconcile.Reconciler {
	return mads_reconcile.NewReconciler(hasher, cache, generator, versioner)
}

type restReconcilerCallbacks struct {
	reconciler mads_reconcile.Reconciler
}

func (r *restReconcilerCallbacks) OnFetchRequest(ctx context.Context, request util_xds.DiscoveryRequest) error {
	nodei := request.Node()

	node, ok := nodei.(*envoy_core.Node)
	if !ok {
		return errors.Errorf("expecting a v3 Node, got: %v", nodei)
	}

	// TODO we need to reconcile on demand because if we us MeshMetric we will only reconcile in watchdog for single client
	// that send request first because of watchdog callback implementation: pkg/util/xds/v3/watchdog_callbacks.go:48
	// Moreover grpc sotw server is never used in real world scenario so we probably need to thing of different syncing cache mechanism
	// if performance of ondemand reconcile will turn out to be poor.
	// Also we probably can remove sync tracker since it will run and recompute MADS response that won't be used
	// Issue: https://github.com/kumahq/kuma/issues/8764
	return r.reconciler.Reconcile(ctx, node)
}

func (r *restReconcilerCallbacks) OnFetchResponse(request util_xds.DiscoveryRequest, response util_xds.DiscoveryResponse) {
}

func NewReconcilerRestCallbacks(reconciler mads_reconcile.Reconciler) util_xds.RestCallbacks {
	return &restReconcilerCallbacks{reconciler: reconciler}
}

func NewSyncTracker(reconciler mads_reconcile.Reconciler, refresh time.Duration, log logr.Logger) envoy_xds.Callbacks {
	return util_xds_v3.NewWatchdogCallbacks(func(ctx context.Context, node *envoy_core.Node, streamID int64) (util_watchdog.Watchdog, error) {
		log := log.WithValues("streamID", streamID, "node", node)
		return &util_watchdog.SimpleWatchdog{
			NewTicker: func() *time.Ticker {
				return time.NewTicker(refresh)
			},
			OnTick: func(context.Context) error {
				log.V(1).Info("on tick")
				return reconciler.Reconcile(ctx, node)
			},
			OnError: func(err error) {
				log.Error(err, "OnTick() failed")
			},
		}, nil
	})
}

func NewXdsContext(log logr.Logger) (envoy_cache.NodeHash, util_xds_v3.SnapshotCache) {
	hasher := hasher{}
	logger := util_xds.NewLogger(log)
	return hasher, util_xds_v3.NewSnapshotCache(false, hasher, logger)
}

type hasher struct{}

func (_ hasher) ID(node *envoy_core.Node) string {
	// now that we start differentiating between clients are we ok with this config growing for old mechanism (under `mesh.metrics`)
	// or should there be a switch here?
	return node.Id
}
