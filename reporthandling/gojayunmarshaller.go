package reporthandling

import (
	"time"

	"github.com/francoispqt/gojay"
)

/*
  responsible on fast unmarshaling of various COMMON containerscan structures and substructures

*/
// UnmarshalJSONObject - File inside a pkg
func (postureReport *PostureReport) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {

	switch key {
	case "customerGUID":
		err = dec.String(&(postureReport.CustomerGUID))

	case "clusterName":
		err = dec.String(&(postureReport.ClusterName))

	case "reportID":
		err = dec.String(&(postureReport.ReportID))
	case "jobID":
		err = dec.String(&(postureReport.JobID))
	case "generationTime":
		err = dec.Time(&(postureReport.ReportGenerationTime), time.RFC3339)
		postureReport.ReportGenerationTime = postureReport.ReportGenerationTime.Local()
	}
	return err

}

// func (files *PkgFiles) UnmarshalJSONArray(dec *gojay.Decoder) error {
// 	lae := PackageFile{}
// 	if err := dec.Object(&lae); err != nil {
// 		return err
// 	}

// 	*files = append(*files, lae)
// 	return nil
// }

func (postureReport *PostureReport) NKeys() int {
	return 0
}

//------------------------
