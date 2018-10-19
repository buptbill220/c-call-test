#include <stdio.h>
#include <stdlib.h>
#include "test.h"

#include "b.h"

int main() {
    GoInt a = 12;
    GoInt b = 99;
    printf("awesome.Add(12,99) = %d\n", Add(a, b));
    printf("awesome.Cosine(1) = %f\n", (float)(Cosine(1.0)));
    GoInt data[6] = {77, 12, 5, 99, 28, 23};
    GoSlice nums = {data, 6, 6};
    Sort(nums);
    for (int i = 0; i < 6; i++){
        printf("%d,", ((GoInt *)nums.data)[i]);
    }
    GoString msg = {"Hello from C!", 13};
    Log(msg);
    printf("print log===========\n");
	GoInt64 *data1 = (GoInt64*)malloc(sizeof(GoInt64) * 10);
	GoString *str = (GoString*)malloc(sizeof(GoString) * 10);
	for (int i = 0; i < 10; i++) {
		data1[i] = i;
		str[i] = msg;
	}
	GoSlice gs = {data1, 10, 10};
	GoSlice gs1 = {str, 10, 10};
	SetMap(gs1, gs);
	SetData(gs);
	Print();
    free(data1);
    free(str);
    return 0;
}
