package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gov1alpha1 "github.com/mbellgb/godoc-k8s/api/v1alpha1"
)

// PackageReconciler reconciles a Package object
type PackageReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=go.mbell.dev,resources=packages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=go.mbell.dev,resources=packages/status,verbs=get;update;patch

func (r *PackageReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("package", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *PackageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&gov1alpha1.Package{}).
		Complete(r)
}
