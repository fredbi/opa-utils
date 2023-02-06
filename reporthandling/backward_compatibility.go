package reporthandling

import (
	"github.com/kubescape/opa-utils/reporthandling/v1"
)

type (
	// Backward-compatible aliases to type definitions corresponding to v1 models

	// Source represents metadata about a File source.
	//
	// Deprecated: this alias will eventually be desupported. Use an explicit reference to the v1 or the v2 model: reporthandling/v{1|2}.Source
	Source = reporthandling.Source

	// Deprecated: this alias will eventually be desupported. Use an explicit reference to the v1 or the v2 model: reporthandling/v{1|2}.Resource
	Resource = reporthandling.Resource

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.ResourceIDs
	ResourcesIDs = reporthandling.ResourcesIDs

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.Framework
	Framework = reporthandling.Framework

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.FrameworkReport
	FrameworkReport = reporthandling.FrameworkReport

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.Control
	Control = reporthandling.Control

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.ControlReport
	ControlReport = reporthandling.ControlReport

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.PolicyRule
	PolicyRule = reporthandling.PolicyRule

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.RuleResponse
	RuleResponse = reporthandling.RuleResponse

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.RuleReport
	RuleReport = reporthandling.RuleReport

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.RuleStatus
	RuleStatus = reporthandling.RuleStatus

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.RuleMatchObjects
	RuleMatchObjects = reporthandling.RuleMatchObjects

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.PostureReport
	PostureReport = reporthandling.PostureReport

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.AlertObject
	AlertObject = reporthandling.AlertObject

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.AlertScore
	AlertScore = reporthandling.AlertScore

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 or the v2 model: reporthandling/v{1|2}.LastCommit
	LastCommit = reporthandling.LastCommit
)

var (
	// Backward-compatible aliases to package-level function

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.SetUniqueResourcesCounter
	SetUniqueResourcesCounter = reporthandling.SetUniqueResourcesCounter

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.NewResourceIMetadata
	NewResourceIMetadata = reporthandling.NewResourceIMetadata

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.ParseRegoResult
	ParseRegoResult = reporthandling.ParseRegoResult

	// Deprecated: this alias will eventually be desuppported. Use an explicit reference to the v1 model: reporthandling/v1.RegoSourceAggregator
	RegoResourcesAggregator = reporthandling.RegoResourcesAggregator
)

const (
	// Backward-compatible aliases to enum constant definitions

	// RegoLanguage refers to the Open Policy Language v1.
	//
	// Deprecated: this alias will eventually be desupported. Use enums.RegoLanguage instead.
	RegoLanguage = reporthandling.RegoLanguage

	// RegoLanguage2 refers to the Open Policy Language v2.
	//
	// Deprecated: this alias will eventually be desupported. Use enums.RegoLanguage2 instead.
	RegoLanguage2 = reporthandling.RegoLanguage2

	// Deprecated: this alias will eventually be desupported. Use enums.SourceTypeJson instead.
	SourceTypeJson = reporthandling.SourceTypeJson

	// Deprecated: this alias will eventually be desupported. Use enums.SourceTypeYaml instead.
	SourceTypeYaml = reporthandling.SourceTypeYaml

	// Deprecated: this alias will eventually be desupported. Use enums.SourceTypeHelmChart instead.
	SourceTypeHelmChart = reporthandling.SourceTypeHelmChart

	// Deprecated: this alias will eventually be desupported. Use enums.SourceTypeKustomizeDirectory instead.
	SourceTypeKustomizeDirectory = reporthandling.SourceTypeKustomizeDirectory
)
