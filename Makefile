run-all: 
		go test -bench=. -benchmem -benchtime=10x > results_run/all.csv

run-single:
		go test -bench=BenchmarkEncodeSingle -benchmem -benchtime=10x > results_run/single_encode.csv
		go test -bench=BenchmarkDecodeSingle -benchmem -benchtime=10x > results_run/single_decode.csv

run-slice:
		go test -bench=BenchmarkEncodeSlice -benchmem -benchtime=10x > results_run/slice_encode.csv
		go test -bench=BenchmarkDecodeSlice -benchmem -benchtime=10x > results_run/slice_decode.csv

run-ptr-slice:
		go test -bench=BenchmarkEncodePtrSlice -benchmem -benchtime=10x > results_run/slice_ptr_encode.csv
		go test -bench=BenchmarkDecodePtrSlice -benchmem -benchtime=10x > results_run/slice_ptr_decode.csv

.PHONY: run-all run-single run-slice run-ptr-slice