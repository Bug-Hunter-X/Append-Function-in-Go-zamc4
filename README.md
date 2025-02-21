# Go Slice Append Function and Capacity

This example demonstrates a common misconception when using the `append` function with slices in Go.  The capacity of a slice plays a crucial role in how `append` behaves, particularly when exceeding the initial capacity.

The `bug.go` file shows the behavior.  The solution, found in `bugSolution.go`, provides insights into managing slice capacity and avoiding unexpected results.

## Understanding the Issue

Go slices have a length (number of elements) and a capacity (total allocated space).  When appending beyond the capacity, Go internally reallocates the slice, potentially impacting performance and potentially leading to unintended memory behavior.  The example highlights how this reallocation affects the underlying array of the slice, which may not be obvious.

## Solution and Best Practices

The solution explains how to pre-allocate sufficient capacity to avoid repeated allocations when appending multiple elements.  Pre-allocation enhances performance and makes the code more predictable.
