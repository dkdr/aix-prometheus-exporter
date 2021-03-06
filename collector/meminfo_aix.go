package collector

/*
#cgo LDFLAGS: -lperfstat
#include <stdio.h>
#include <libperfstat.h>
#include <stdlib.h>

int getMemInfo(perfstat_memory_total_t *mem_now) {

	int	rc;
	rc = perfstat_memory_total(NULL, mem_now, sizeof(perfstat_memory_total_t), 1);
	if (rc <= 0 ) {
		return rc;
	}
	return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func (c *meminfoCollector) getMemInfo() (map[string]float64, error) {

	var memnow C.perfstat_memory_total_t

	if _, err := C.getMemInfo(&memnow); err != nil {
		return nil, fmt.Errorf("could not collect memory from getMemInfo: %v", err)
	}
	defer C.free(unsafe.Pointer(&memnow))

	// perfstat_memory_total returns data in number of 4k pages
	ps := float64(4000)
	return map[string]float64{
		"real_total_bytes":   ps * float64(memnow.real_total),
		"real_free_bytes":    ps * float64(memnow.real_free),
		"real_pinned_bytes":  ps * float64(memnow.real_pinned),
		"real_inuse_bytes":   ps * float64(memnow.real_inuse),
		"real_system_bytes":  ps * float64(memnow.real_system),
		"real_user_bytes":    ps * float64(memnow.real_user),
		"real_process_bytes": ps * float64(memnow.real_process),
		"real_avail_bytes":   ps * float64(memnow.real_avail),
		"virt_total_bytes":   ps * float64(memnow.virt_total),
		"virt_active_bytes":  ps * float64(memnow.virt_active),
	}, nil

}
