package reporthandling

type ScanningTarget uint16

const (
	Cluster   ScanningTarget = 0
	File      ScanningTarget = 1
	Repo      ScanningTarget = 2
	GitLocal  ScanningTarget = 3
	Directory ScanningTarget = 4
)

func (st *ScanningTarget) String() string {
	switch *st {
	case 0:
		return "Cluster"
	case 1:
		return "File"
	case 2:
		return "Repo"
	case 3:
		return "GitLocal"
	case 4:
		return "Directory"
	default:
		return ""
	}
}
