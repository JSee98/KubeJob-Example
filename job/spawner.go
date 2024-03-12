package job

import (
	kubernetesinternal "blogs/kube-jobs/kubernetes-internal"
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func SpawnBasicJob(concurrency int) error {
	clientSet, err :=  kubernetesinternal.GetKubernetesClient()
	if err!= nil{
		return err
	}

	for i:= 0; i<concurrency; i++{
		jobsInterface := clientSet.BatchV1().Jobs("default")
		job, err := getBasicJob(fmt.Sprintf("job-%d", i), "hello-world")
		if err != nil {
			return err
		}
		createdJob, err := jobsInterface.Create(context.TODO(), job, metav1.CreateOptions{})
		if err != nil {
			return err
		}

		fmt.Printf("Job %s created successfully\n", createdJob.Name)
	}
	return nil
}

func getBasicJob(jobName, image string) (*batchv1.Job, error) {
    job := &batchv1.Job{
        ObjectMeta: metav1.ObjectMeta{
            Name: jobName,
        },
        Spec: batchv1.JobSpec{
            Template: corev1.PodTemplateSpec{
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{
                        {
                            Name:    fmt.Sprintf("job-%s", jobName),
                            Image:   image,
                            Command: []string{}, 
                        },
                    },
                    RestartPolicy: corev1.RestartPolicyNever,
                },
            },
        },
    }

    return job, nil
}
