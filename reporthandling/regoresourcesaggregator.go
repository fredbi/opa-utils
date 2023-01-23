package reporthandling

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/kubescape/k8s-interface/k8sinterface"
	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/objectsenvelopes"
)

var aggregatorAttribute = "resourcesAggregator"

func RegoResourcesAggregator(rule *PolicyRule, k8sObjects []workloadinterface.IMetadata) ([]workloadinterface.IMetadata, error) {
	if aggregateBy, ok := rule.Attributes[aggregatorAttribute]; ok {
		switch aggregateBy {
		case "subject-role-rolebinding":
			return AggregateResourcesBySubjects(k8sObjects)
		case "apiserver-pod":
			if obj := AggregateResourcesByAPIServerPod(k8sObjects); obj != nil {
				return []workloadinterface.IMetadata{obj}, nil
			}
			return []workloadinterface.IMetadata{}, nil
		default:
			return k8sObjects, nil
		}
	}
	return k8sObjects, nil
}

func AggregateResourcesBySubjects(k8sObjects []workloadinterface.IMetadata) ([]workloadinterface.IMetadata, error) {
	aggregatedK8sObjects := []workloadinterface.IMetadata{}

	for _, bindingWorkload := range k8sObjects {
		if !strings.HasSuffix(bindingWorkload.GetKind(), "Binding") { // types.Role
			continue
		}

		for _, roleWorkload := range k8sObjects {
			if !strings.HasSuffix(roleWorkload.GetKind(), "Role") {
				continue
			}

			bindingWorkloadObj := bindingWorkload.GetObject()
			kindVal, isKindRoleRef := workloadinterface.InspectMap(bindingWorkloadObj, "roleRef", "kind")
			if !isKindRoleRef {
				continue
			}

			nameVal, isNameRoleRef := workloadinterface.InspectMap(bindingWorkloadObj, "roleRef", "name")
			if !isNameRoleRef {
				continue
			}

			kind, isString := kindVal.(string)
			if !isString { // NOTE: now returns an error - used to panic
				return nil, fmt.Errorf("expected a Kind as string, but got %T", kindVal)
			}

			name, isString := nameVal.(string)
			if !isString { // NOTE: now returns an error - used to panic
				return nil, fmt.Errorf("expected a Kind as string, but got %T", kindVal)
			}

			if kind != roleWorkload.GetKind() || name != roleWorkload.GetName() {
				continue
			}

			subjects, isSubjects := workloadinterface.InspectMap(bindingWorkloadObj, "subjects")
			if !isSubjects {
				continue
			}

			data, isInterfaces := subjects.([]interface{})
			if !isInterfaces {
				continue
			}

			for _, subject := range data {
				// deep copy subject - don't change original subject in rolebinding
				subjectCopy, err := DeepCopyMap(subject.(map[string]interface{}))
				if err != nil {
					return aggregatedK8sObjects, err
				}

				// subjectCopy["apiVersion"] = fmt.Sprintf("%s/%s", objectsenvelopes.RegoAttackVectorGroup, objectsenvelopes.RegoAttackVectorVersion)
				subjectCopy[objectsenvelopes.RelatedObjectsKey] = []map[string]interface{}{
					bindingWorkload.GetObject(),
					roleWorkload.GetObject(),
				}
				newObj := objectsenvelopes.NewRegoResponseVectorObject(subjectCopy)
				aggregatedK8sObjects = append(aggregatedK8sObjects, newObj)
			}
		}
	}

	return aggregatedK8sObjects, nil
}

// Create custom object of apiserver pod. Has required fields + cmdline
func AggregateResourcesByAPIServerPod(k8sObjects []workloadinterface.IMetadata) workloadinterface.IMetadata {
	for _, obj := range k8sObjects {
		if !k8sinterface.IsTypeWorkload(obj.GetObject()) {
			continue
		}
		workload := workloadinterface.NewWorkloadObj(obj.GetObject())
		if workload.GetKind() == "Pod" && workload.GetNamespace() == "kube-system" {
			if strings.Contains(workload.GetName(), "apiserver") || strings.Contains(workload.GetName(), "api-server") {
				return workload
			}
		}
	}
	return nil
}

// DeepCopyMap performs a deep copy of the given map m.
func DeepCopyMap(m map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	var copy map[string]interface{}
	err = dec.Decode(&copy)
	if err != nil {
		return nil, err
	}
	return copy, nil
}
