package reporthandling

import "github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"

type IBasicPostureReport interface {
	reportsummary.IBasicPostureReport
	GetScannigTarget() ScanningTarget
	GetContextMetadata() *ContextMetadata
	GetRepositoryHash() string
}
