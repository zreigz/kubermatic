package main

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/types"

	rbacusercluster "github.com/kubermatic/kubermatic/api/pkg/controller/rbac-user-cluster"
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/resources"

	appsv1 "k8s.io/api/apps/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *testRunner) testUserclusterControllerRBAC(log *zap.SugaredLogger, cluster *kubermaticv1.Cluster, userClusterClient, seedClusterClient ctrlruntimeclient.Client) error {
	ctx := context.Background()
	clusterNamespace := fmt.Sprintf("cluster-%s", cluster.Name)

	log = log.With("cluster", clusterNamespace)
	log.Info("Testing user cluster RBAC controller")

	// check if usercluster-controller was deployed on seed cluster
	deployment := &appsv1.Deployment{}
	if err := seedClusterClient.Get(ctx, types.NamespacedName{Namespace: clusterNamespace, Name: resources.UserClusterControllerDeploymentName}, deployment); err != nil {
		return fmt.Errorf("failed to get Deployment %s: %v", resources.UserClusterControllerDeploymentName, err)
	}

	if deployment.Status.AvailableReplicas == 0 {
		return fmt.Errorf("Deployment %s is not ready", resources.UserClusterControllerDeploymentName)
	}

	// check user cluster resources: ClusterRoles and ClusterRoleBindings
	for _, resourceName := range rbacResourceNames() {
		rlog := log.With("cluster-role", resourceName)

		rlog.Info("Validating ClusterRole...")
		clusterRole := &rbacv1.ClusterRole{}
		if err := userClusterClient.Get(ctx, types.NamespacedName{Name: resourceName}, clusterRole); err != nil {
			return fmt.Errorf("failed to get ClusterRole %s: %v", resourceName, err)
		}

		defaultClusterRole, err := rbacusercluster.GenerateRBACClusterRole(resourceName)
		if err != nil {
			return fmt.Errorf("failed to generate default ClusterRole %s: %v", resourceName, err)
		}

		if !equality.Semantic.DeepEqual(clusterRole.Rules, defaultClusterRole.Rules) {
			return fmt.Errorf("incorrect ClusterRole.Rules were returned, got: %v, want: %v", clusterRole.Rules, defaultClusterRole.Rules)
		}

		rlog = log.With("cluster-role-binding", resourceName)

		rlog.Info("Validating ClusterRoleBinding...")
		clusterRoleBinding := &rbacv1.ClusterRoleBinding{}
		if err := userClusterClient.Get(ctx, types.NamespacedName{Name: resourceName}, clusterRoleBinding); err != nil {
			return fmt.Errorf("failed to get ClusterRoleBinding %s: %v", resourceName, err)
		}

		defaultClusterRoleBinding, err := rbacusercluster.GenerateRBACClusterRoleBinding(resourceName)
		if err != nil {
			return fmt.Errorf("failed to generate default ClusterRoleBinding %s: %v", resourceName, err)
		}

		if !equality.Semantic.DeepEqual(clusterRoleBinding.RoleRef, defaultClusterRoleBinding.RoleRef) {
			return fmt.Errorf("incorrect ClusterRoleBinding.RoleRef was returned, got: %v, want: %v", clusterRoleBinding.RoleRef, defaultClusterRoleBinding.RoleRef)
		}
		if !equality.Semantic.DeepEqual(clusterRoleBinding.Subjects, defaultClusterRoleBinding.Subjects) {
			return fmt.Errorf("incorrect ClusterRoleBinding.Subjects were returned, got: %v, want: %v", clusterRoleBinding.Subjects, defaultClusterRoleBinding.Subjects)
		}
	}

	return nil
}

func rbacResourceNames() []string {
	return []string{rbacusercluster.ResourceOwnerName, rbacusercluster.ResourceEditorName, rbacusercluster.ResourceViewerName}
}
