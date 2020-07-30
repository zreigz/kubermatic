/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package kubernetes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	providerconfig "github.com/kubermatic/machine-controller/pkg/providerconfig/types"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	kubermaticapiv1 "github.com/kubermatic/kubermatic/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/pkg/provider"
	"github.com/kubermatic/kubermatic/pkg/resources"
	"github.com/kubermatic/kubermatic/pkg/util/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// ExternalClusterProvider struct that holds required components in order to provide connection to the cluster
type ExternalClusterProvider struct {
	// createSeedImpersonatedClient is used as a ground for impersonation
	createMasterImpersonatedClient impersonationClient
	clientPrivileged               ctrlruntimeclient.Client
	restMapperCache                *restmapper.Cache
}

// NewExternalClusterProvider returns an external cluster provider
func NewExternalClusterProvider(createMasterImpersonatedClient impersonationClient, client ctrlruntimeclient.Client) (*ExternalClusterProvider, error) {
	return &ExternalClusterProvider{
		createMasterImpersonatedClient: createMasterImpersonatedClient,
		clientPrivileged:               client,
		restMapperCache:                restmapper.New(),
	}, nil
}

// New creates a brand new external cluster in the system with the given name
func (p *ExternalClusterProvider) New(userInfo *provider.UserInfo, project *kubermaticapiv1.Project, cluster *kubermaticapiv1.ExternalCluster) (*kubermaticapiv1.ExternalCluster, error) {
	masterImpersonatedClient, err := createImpersonationClientWrapperFromUserInfo(userInfo, p.createMasterImpersonatedClient)
	if err != nil {
		return nil, err
	}
	addProjectReference(project, cluster)
	if err := masterImpersonatedClient.Create(context.Background(), cluster); err != nil {
		return nil, err
	}
	return cluster, nil
}

// NewUnsecured creates a brand new external cluster in the system with the given name
//
// Note that this function:
// is unsafe in a sense that it uses privileged account to create the resource
func (p *ExternalClusterProvider) NewUnsecured(project *kubermaticapiv1.Project, cluster *kubermaticapiv1.ExternalCluster) (*kubermaticapiv1.ExternalCluster, error) {
	addProjectReference(project, cluster)
	if err := p.clientPrivileged.Create(context.Background(), cluster); err != nil {
		return nil, err
	}
	return cluster, nil
}

func addProjectReference(project *kubermaticapiv1.Project, cluster *kubermaticapiv1.ExternalCluster) {
	if cluster.Labels == nil {
		cluster.Labels = make(map[string]string)
	}
	cluster.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion: kubermaticapiv1.SchemeGroupVersion.String(),
			Kind:       kubermaticapiv1.ProjectKindName,
			UID:        project.GetUID(),
			Name:       project.Name,
		},
	}
	cluster.Labels[kubermaticapiv1.ProjectIDLabelKey] = project.Name
}

func (p *ExternalClusterProvider) GenerateClient(cfg *clientcmdapi.Config) (*ctrlruntimeclient.Client, error) {
	iconfig := clientcmd.NewNonInteractiveClientConfig(
		*cfg,
		resources.KubeconfigDefaultContextKey,
		&clientcmd.ConfigOverrides{},
		nil,
	)

	clientConfig, err := iconfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	// Avoid blocking of the controller by increasing the QPS for user cluster interaction
	clientConfig.QPS = 20
	clientConfig.Burst = 50

	client, err := p.restMapperCache.Client(clientConfig)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (p *ExternalClusterProvider) CreateOrUpdateKubeconfigSecretForCluster(ctx context.Context, cluster *kubermaticapiv1.ExternalCluster, kubeconfig *clientcmdapi.Config) error {

	rawData, err := json.Marshal(kubeconfig)
	if err != nil {
		return err
	}
	kubeconfigRef, err := p.ensureKubeconfigSecret(ctx, cluster, map[string][]byte{
		resources.ExternalClusterKubeconfig: rawData,
	})
	if err != nil {
		return err
	}
	cluster.Spec.KubeconfigReference = kubeconfigRef
	return nil
}

func (p *ExternalClusterProvider) ensureKubeconfigSecret(ctx context.Context, cluster *kubermaticapiv1.ExternalCluster, secretData map[string][]byte) (*providerconfig.GlobalSecretKeySelector, error) {
	name := cluster.GetKubeconfigSecretName()

	namespacedName := types.NamespacedName{Namespace: resources.KubermaticNamespace, Name: name}
	existingSecret := &corev1.Secret{}
	if err := p.clientPrivileged.Get(ctx, namespacedName, existingSecret); err != nil && !kerrors.IsNotFound(err) {
		return nil, fmt.Errorf("failed to probe for secret %q: %v", name, err)
	}

	if existingSecret.Name == "" {
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: resources.KubermaticNamespace,
			},
			Type: corev1.SecretTypeOpaque,
			Data: secretData,
		}

		if err := p.clientPrivileged.Create(ctx, secret); err != nil {
			return nil, fmt.Errorf("failed to create kubeconfig secret: %v", err)
		}
	} else {
		if existingSecret.Data == nil {
			existingSecret.Data = map[string][]byte{}
		}

		requiresUpdate := false

		for k, v := range secretData {
			if !bytes.Equal(v, existingSecret.Data[k]) {
				requiresUpdate = true
				break
			}
		}

		if requiresUpdate {
			existingSecret.Data = secretData
			if err := p.clientPrivileged.Update(ctx, existingSecret); err != nil {
				return nil, fmt.Errorf("failed to update kubeconfig secret: %v", err)
			}
		}
	}

	return &providerconfig.GlobalSecretKeySelector{
		ObjectReference: corev1.ObjectReference{
			Name:      name,
			Namespace: resources.KubermaticNamespace,
		},
	}, nil
}