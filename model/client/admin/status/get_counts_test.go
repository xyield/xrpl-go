package status

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestGetCountsResponse(t *testing.T) {
	s := GetCountsResponse{
		"AL_hit_rate":          float64(48.36725616455078),
		"HashRouterEntry":      float64(3048),
		"Ledger":               float64(46),
		"NodeObject":           float64(10417),
		"SLE_hit_rate":         float64(64.62035369873047),
		"STArray":              float64(1299),
		"STLedgerEntry":        float64(646),
		"STObject":             float64(6987),
		"STTx":                 float64(4104),
		"STValidation":         float64(610),
		"Transaction":          float64(4069),
		"dbKBLedger":           float64(10733),
		"dbKBTotal":            float64(39069),
		"dbKBTransaction":      float64(26982),
		"fullbelow_size":       float64(0),
		"historical_perminute": float64(0),
		"ledger_hit_rate":      float64(71.0565185546875),
		"node_hit_rate":        float64(3.808214902877808),
		"node_read_bytes":      float64(393611911),
		"node_reads_hit":       float64(1283098),
		"node_reads_total":     float64(679410),
		"node_writes":          float64(1744285),
		"node_written_bytes":   float64(794368909),
		"status":               "success",
		"treenode_cache_size":  float64(6650),
		"treenode_track_size":  float64(598631),
		"uptime":               "3 hours, 50 minutes, 27 seconds",
		"write_load":           float64(0),
	}

	j := `{
	"AL_hit_rate": 48.36725616455078,
	"HashRouterEntry": 3048,
	"Ledger": 46,
	"NodeObject": 10417,
	"SLE_hit_rate": 64.62035369873047,
	"STArray": 1299,
	"STLedgerEntry": 646,
	"STObject": 6987,
	"STTx": 4104,
	"STValidation": 610,
	"Transaction": 4069,
	"dbKBLedger": 10733,
	"dbKBTotal": 39069,
	"dbKBTransaction": 26982,
	"fullbelow_size": 0,
	"historical_perminute": 0,
	"ledger_hit_rate": 71.0565185546875,
	"node_hit_rate": 3.808214902877808,
	"node_read_bytes": 393611911,
	"node_reads_hit": 1283098,
	"node_reads_total": 679410,
	"node_writes": 1744285,
	"node_written_bytes": 794368909,
	"status": "success",
	"treenode_cache_size": 6650,
	"treenode_track_size": 598631,
	"uptime": "3 hours, 50 minutes, 27 seconds",
	"write_load": 0
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
