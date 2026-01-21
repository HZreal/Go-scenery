#ifndef PROCESS_H
#define PROCESS_H

typedef struct {
    int sum;        // nums 的和
    double avg;     // nums 的平均
    char* message;  // 信息字符串
} Result;

// 核心函数
Result* process(int id, double score, const char* name, int* nums, int n);

// 释放 Result 内存
void free_result(Result* r);

#endif
