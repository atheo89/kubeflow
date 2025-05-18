package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elyra secret creation", func() {
	ctx := context.Background()
	const (
		Name      = "test-notebook"
		Namespace = "elyra-test-ns"
	)

	BeforeEach(func() {
		By("Creating test namespace if not exists")
		err := cli.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: Namespace}})
		if err != nil && !apierrs.IsAlreadyExists(err) {
			Expect(err).ToNot(HaveOccurred())
		}
	})

	When("A DSPA object is present in the namespace then create the ds-pipeline-config secret", func() {
		const (
			name      = "dspa"
			namespace = "default"
		)

		// It("should create the DSPA and ensure ds-pipeline-config secret is generated", func() {
		// 	dspa := &unstructured.Unstructured{
		// 		Object: map[string]interface{}{
		// 			"apiVersion": "datasciencepipelinesapplications.opendatahub.io/v1alpha1",
		// 			"kind":       "DataSciencePipelinesApplication",
		// 			"metadata": map[string]interface{}{
		// 				"name":      name,
		// 				"namespace": namespace,
		// 			},
		// 			"spec": map[string]interface{}{
		// 				"objectStorage": map[string]interface{}{
		// 					"externalStorage": map[string]interface{}{
		// 						"host":   "s3.amazonaws.com",
		// 						"bucket": "ods-ci-ds-pipelines",
		// 						"scheme": "https",
		// 						"s3CredentialsSecret": map[string]interface{}{
		// 							"secretName": "dashboard-dspa-secret",
		// 							"accessKey":  "AWS_ACCESS_KEY_ID",
		// 							"secretKey":  "AWS_SECRET_ACCESS_KEY",
		// 						},
		// 					},
		// 				},
		// 			},
		// 			"status": map[string]interface{}{
		// 				"components": map[string]interface{}{
		// 					"apiServer": map[string]interface{}{
		// 						"externalUrl": "https://dspa.external.url",
		// 					},
		// 				},
		// 			},
		// 		},
		// 	}

		// 	dspa.SetGroupVersionKind(schema.GroupVersionKind{
		// 		Group:   "datasciencepipelinesapplications.opendatahub.io",
		// 		Version: "v1alpha1",
		// 		Kind:    "DataSciencePipelinesApplication",
		// 	})

		// 	By("Creating a fake DSPA object")
		// 	Expect(cli.Create(ctx, dspa)).To(Succeed())

		// 	By("Waiting for the ds-pipeline-config secret to be created")
		// 	Eventually(func() error {
		// 		secret := &corev1.Secret{}
		// 		return cli.Get(ctx, client.ObjectKey{Name: "ds-pipeline-config", Namespace: namespace}, secret)
		// 	}).Should(Succeed())

		// 	By("Verifying the secret contains expected keys")
		// 	createdSecret := &corev1.Secret{}
		// 	Expect(cli.Get(ctx, client.ObjectKey{Name: "ds-pipeline-config", Namespace: namespace}, createdSecret)).To(Succeed())

		// 	Expect(createdSecret.Data).To(HaveKey("ELYRA_AIRFLOW_URL"))
		// 	Expect(createdSecret.Data).To(HaveKey("ELYRA_AIRFLOW_USERNAME"))
		// 	Expect(createdSecret.Data).To(HaveKey("ELYRA_AIRFLOW_PASSWORD"))
		// })
	})
})
