### V 0.1.0
I took the standard simple approach for this: loopping throught the file using a buffer scanner. 
I stored had a station struct and an array of these structs, I then compared the values for each processed line. This gave me a runtime of under 5 minutes(which is quite impressive).

### V 0.2.0
CPU profiling showed that a lot of time was spent on the string split function. Since we only need to split it at one point in instead chose to find the index of the ';' character and use a slice to reference the values. This reduced the execution time by ~50 seconds, bringing us down to ~3m,50s. 

### V 0.3.0
Implemented hashmaps because most of the time was spent looking through the array for a matching station. This brought execution time by ~30s. But I feel this introduced another problem.

