package resourcesresults

import (
	"github.com/armosec/armoapi-go/armotypes"
	"github.com/kubescape/k8s-interface/workloadinterface"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	"github.com/kubescape/opa-utils/reporthandling/v1"
)

func MockResults() []Result {
	return []Result{
		*mockResultPassed(),
		*mockResultFailed(),
		*mockResultException(),
		*mockResultConfiguration(),
		*mockResultRequiresReview(),
		*mockResultManualReview(),
	}
}

func mockResultPassed() *Result {
	w := workloadinterface.NewWorkloadMock(nil)
	return &Result{
		ResourceID: w.GetID(),
		AssociatedControls: []ResourceAssociatedControl{
			*mockResourceAssociatedControl0089Passed(),
		},
	}
}

func mockResultException() *Result {
	w := workloadinterface.NewWorkloadMock(nil)
	return &Result{
		ResourceID: w.GetID(),
		AssociatedControls: []ResourceAssociatedControl{
			*mockResourceAssociatedControlException(),
		},
	}
}

func mockResultConfiguration() *Result {
	w := workloadinterface.NewWorkloadMock(nil)
	r := mockResourceAssociatedControlConfiguration()
	r.SetStatus(*mockControlWithActionRequiredConfiguration())
	return &Result{
		ResourceID: w.GetID(),
		AssociatedControls: []ResourceAssociatedControl{
			*r,
		},
	}
}

func mockResultRequiresReview() *Result {
	w := workloadinterface.NewWorkloadMock(nil)
	r := mockResourceAssociatedControlRequiresReview()
	r.SetStatus(*mockControlWithActionRequiredRequiresReview())
	return &Result{
		ResourceID: w.GetID(),
		AssociatedControls: []ResourceAssociatedControl{
			*r,
		},
	}
}

func mockResultManualReview() *Result {
	w := workloadinterface.NewWorkloadMock(nil)
	r := mockResourceAssociatedControlManualReview()
	r.SetStatus(*mockControlWithActionRequiredManualReview())
	return &Result{
		ResourceID: w.GetID(),
		AssociatedControls: []ResourceAssociatedControl{
			*r,
		},
	}
}

//	func mockResultSkipped() *Result {
//		return &Result{
//			ResourceID:         "resource/passed",
//			AssociatedControls: []ResourceAssociatedControl{},
//		}
//	}
func mockResultFailed() *Result {
	w := workloadinterface.NewWorkloadMock(nil)
	return &Result{
		ResourceID: w.GetID(),
		AssociatedControls: []ResourceAssociatedControl{
			*mockResourceAssociatedControl0087Failed(),
			*mockResourceAssociatedControl0088Failed(),
			*mockResourceAssociatedControl0089Passed(),
		},
	}
}
func mockResourceAssociatedControlException() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0089",
		Name:      "0089",
		Status:    apis.StatusInfo{InnerStatus: apis.StatusPassed, SubStatus: apis.SubStatusException},
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleException(),
		},
	}
}
func mockResourceAssociatedControlConfiguration() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0089",
		Name:      "0089",
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleMissedConfiguration(),
		},
	}
}

func mockResourceAssociatedControlRequiresReview() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0089",
		Name:      "0089",
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleRequiresReview(),
		},
	}
}

func mockResourceAssociatedControlManualReview() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0089",
		Name:      "0089",
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleManualReview(),
		},
	}
}

func mockResourceAssociatedControl0087Failed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0087",
		Name:      "0087",
		Status:    apis.StatusInfo{InnerStatus: apis.StatusFailed, SubStatus: apis.SubStatusUnknown},
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleA(),
			*mockResourceAssociatedRuleB(),
		},
	}
}
func mockResourceAssociatedControlFailed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0087",
		Name:      "0087",
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleA(),
			*mockResourceAssociatedRuleB(),
		},
	}
}

func mockResourceAssociatedControl0088Failed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0088",
		Name:      "0088",
		Status:    apis.StatusInfo{InnerStatus: apis.StatusFailed, SubStatus: apis.SubStatusUnknown},
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRuleB(),
		},
	}
}

func mockResourceAssociatedControl0089Passed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0089",
		Name:      "0089",
		Status:    apis.StatusInfo{InnerStatus: apis.StatusPassed, SubStatus: apis.SubStatusUnknown},
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRulePassed(),
		},
	}
}
func mockResourceAssociatedControlPassed() *ResourceAssociatedControl {
	return &ResourceAssociatedControl{
		ControlID: "C-0089",
		Name:      "0089",
		ResourceAssociatedRules: []ResourceAssociatedRule{
			*mockResourceAssociatedRulePassed(),
		},
	}
}
func mockResourceAssociatedRuleA() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:                  "ruleA",
		Status:                apis.StatusFailed,
		SubStatus:             apis.SubStatusUnknown,
		Paths:                 []armotypes.PosturePaths{{FailedPath: "path/to/fail/A"}},
		Exception:             []armotypes.PostureExceptionPolicy{},
		ControlConfigurations: nil,
	}
}

func mockResourceAssociatedRuleB() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:                  "ruleB",
		Status:                apis.StatusFailed,
		SubStatus:             apis.SubStatusUnknown,
		Paths:                 []armotypes.PosturePaths{{FailedPath: "path/to/fail/B"}},
		Exception:             []armotypes.PostureExceptionPolicy{},
		ControlConfigurations: nil,
	}
}

func mockResourceAssociatedRulePassed() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:      "rulePassed",
		Status:    apis.StatusPassed,
		SubStatus: apis.SubStatusUnknown,
	}
}

func mockResourceAssociatedRuleException() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:      "ruleException",
		Status:    apis.StatusPassed,
		SubStatus: apis.SubStatusException,
	}
}

func mockResourceAssociatedRuleMissedConfiguration() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:                  "ruleMissedConfiguration",
		Status:                apis.StatusFailed,
		SubStatus:             apis.SubStatusUnknown,
		Paths:                 []armotypes.PosturePaths{{FailedPath: "path/to/fail/B"}},
		Exception:             []armotypes.PostureExceptionPolicy{},
		ControlConfigurations: nil,
	}
}

func mockResourceAssociatedRuleRequiresReview() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:                  "ruleRequiresReview",
		Status:                apis.StatusFailed,
		SubStatus:             apis.SubStatusUnknown,
		Paths:                 []armotypes.PosturePaths{{FailedPath: "path/to/fail/B"}},
		Exception:             []armotypes.PostureExceptionPolicy{},
		ControlConfigurations: nil,
	}
}

func mockResourceAssociatedRuleManualReview() *ResourceAssociatedRule {
	return &ResourceAssociatedRule{
		Name:                  "ruleManualReview",
		Status:                apis.StatusFailed,
		SubStatus:             apis.SubStatusUnknown,
		Paths:                 []armotypes.PosturePaths{{FailedPath: "path/to/fail/B"}},
		Exception:             []armotypes.PostureExceptionPolicy{},
		ControlConfigurations: nil,
	}
}

func mockControlWithActionRequiredConfiguration() *reporthandling.Control {
	return &reporthandling.Control{
		PortalBase: armotypes.PortalBase{Attributes: map[string]interface{}{"actionRequired": "configuration"}},
	}
}

func mockControlWithActionRequiredRequiresReview() *reporthandling.Control {
	return &reporthandling.Control{
		PortalBase: armotypes.PortalBase{Attributes: map[string]interface{}{"actionRequired": "requires review"}},
	}
}

func mockControlWithActionRequiredManualReview() *reporthandling.Control {
	return &reporthandling.Control{
		PortalBase: armotypes.PortalBase{Attributes: map[string]interface{}{"actionRequired": "manual review"}},
	}
}

// func mockResourceAssociatedRuleWithFWException() *ResourceAssociatedRule {
// 	return &ResourceAssociatedRule{
// 		Name:        "ruleB",
// 		FailedPaths: []string{"path/to/fail/B"},
// 		Exception:   []armotypes.PostureExceptionPolicy{},
// 	}
// }
