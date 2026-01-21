#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "process.h"

// 计算 nums 的和、平均，并返回 Result*
Result* process(int id, double score, const char* name, int* nums, int n) {
    if (!nums || n <= 0) return NULL;

    Result* r = (Result*)malloc(sizeof(Result));
    if (!r) return NULL;

    r->sum = 0;
    for (int i = 0; i < n; i++) {
        r->sum += nums[i];
    }

    r->avg = r->sum / (double)n;

    char buf[256];
    snprintf(buf, sizeof(buf), "id=%d name=%s score=%.2f", id, name, score);
    r->message = (char*)malloc(strlen(buf) + 1);
    strcpy(r->message, buf);

    return r;
}

// 释放 Result 内存
void free_result(Result* r) {
    if (!r) return;
    if (r->message) free(r->message);
    free(r);
}
