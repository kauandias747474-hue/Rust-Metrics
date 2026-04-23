#include "../include/processor.h"
#include <vector>
#include <cmath>
#include <numeric>

extern "C"
   ProcessedResult process_telemetry_data(const float* data, int size) {
    ProcessedResult result = {0.0f, 0.0f, size};

    try {

        if (data == nullptr || size <= 1) {
            return result;
        }
    

        float sum = 0.0f;
        for (int i = 0; i < size; ++i) {
            sum += data[i];
        }
        result.mean = sum / size;

        float variance_sum = 0.0f
        for (int i = 0; i < size; i++) { 
            variance_sum += std::pow(data[i] - result.mean, 2);
        }


        result.std_dev = std::sqrt(variance_sum / size);


   } catch (...) {
    
    result.std_dev = - 1.0f
   
   }
}

