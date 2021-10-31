/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	dbv1 "github.com/yavosh/sqln-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	jobOwnerKey = ".metadata.controller"
	apiGVStr    = dbv1.GroupVersion.String()
)

// SqlnReconciler reconciles a Sqln object
type SqlnReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=db.narity.com,resources=sqlns,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=db.narity.com,resources=sqlns/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=db.narity.com,resources=sqlns/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *SqlnReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Reconcile ...", "name", req.NamespacedName)

	var sqlCR dbv1.Sqln
	if err := r.Get(ctx, req.NamespacedName, &sqlCR); err != nil {
		logger.Error(err, "unable to fetch")
		// we'll ignore not-found errors
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("sqlCR", "sqln", sqlCR)

	var childSets appsv1.StatefulSetList
	// client.MatchingFields{jobOwnerKey: req.Name}
	if err := r.List(ctx, &childSets, client.InNamespace(req.Namespace)); err != nil {
		logger.Error(err, "unable to list child Jobs")
		return ctrl.Result{}, err
	}

	for _, s := range childSets.Items {
		fmt.Printf("Set %+v\n", s)
	}

	// TODO(user): your logic here
	scheduledResult := ctrl.Result{RequeueAfter: time.Second * 10}

	return scheduledResult, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SqlnReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dbv1.Sqln{}).
		Complete(r)
}
