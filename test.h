#ifndef _TEST_H_
#define _TEST_H_

#include "b.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    GoString Imei;
    GoString Idfa;
    GoInt64 AppId;
    GoSlice RetargetingTags;
    GoInt64 Rit;
} ReqContext;

typedef struct {
    GoInt64 Aid;
    GoInt64 AdvId;
    GoString Title;
    GoSlice Cids;
} AdEngineModel;


#ifdef __cplusplus
}
#endif

#endif
