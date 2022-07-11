/*
 * Copyright (c) 2017-2022 Carl Kittelberger.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

#include <stdlib.h>
#include "stdint.h"

struct buffer_table_index
{
	int index;
};

int bsdiff_cgo(uint8_t* oldptr, int64_t oldsize,
	uint8_t* newptr, int64_t newsize,
	int bufferIndex);
	
int bspatch_cgo(uint8_t* oldptr, int64_t oldsize,
	uint8_t* newptr, int64_t newsize,
	int bufferIndex);
