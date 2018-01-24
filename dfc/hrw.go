/*
 * Copyright (c) 2017, NVIDIA CORPORATION. All rights reserved.
 *
 */
package dfc

import (
	"github.com/OneOfOne/xxhash"
)

const LCG32 = 1103515245

// A variant of consistent hash based on rendezvous algorithm by Thaler and Ravishankar,
// aka highest random weight (HRW)

func hrwTarget(name string, smap *Smap) (si *daemonInfo) {
	// NOTE: commented out on purpose - trading off read access to unlocked map
	// 	 (that changes very rarely)
	//       vs locking zillion times - use sync.Map otherwise
	// smap.lock.Lock()
	// defer smap.lock.Unlock()
	var max uint32
	for id, sinfo := range smap.Smap {
		cs := xxhash.ChecksumString32S(id+":"+name, LCG32)
		if cs > max {
			max = cs
			si = sinfo
		}
	}
	return
}

func hrwMpath(name string) (mpath string) {
	var max uint32
	for path, _ := range ctx.mountpaths {
		cs := xxhash.ChecksumString32S(path+":"+name, LCG32)
		if cs > max {
			max = cs
			mpath = path
		}
	}
	return
}
