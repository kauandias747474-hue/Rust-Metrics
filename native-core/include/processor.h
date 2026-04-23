#ifndef PROCESSOR_H
#define PROCESSOR_H

#ifdef __cplusplus
extern "C" {
#endif


typedef struct {
    float std_dev;
    float mean;
    int count;
} ProcessedResult;


ProcessedResult process_telemetry_data(const float* data, int size);

#ifdef __cplusplus
}
#endif

#endif
