# traceutils

This repository contains code for decoding and encoding runtime/trace files as well as useful functionality implemented on top.

```
go install github.com/felixge/traceutils/cmd/traceutils@latest
```

## anonymize

The anonymize command can be used to remove all file paths, function names and user logs from a trace file. The go stdlib is not anonymized, but all other packages are. This is useful for sharing traces that may contain sensitive information.

```
traceutils anonymize <input> <output>
```

Example output:

![screenshot of go tool trace showing an anonymized trace](./images/anonymize.png)

## breakdown

The breakdown command can be used to analyze the contents of a trace.

### bytes

```
traceutils breakdown bytes <input>
```

Example output:

```
+------------------------+----------+---------+
|       EVENT TYPE       |  BYTES   |    %    |
+------------------------+----------+---------+
| EventGoSched           | 235.9 kB | 35.88%  |
| EventGoStartLocal      | 193.8 kB | 29.48%  |
| EventHeapAlloc         | 81.6 kB  | 12.41%  |
| EventStack             | 44.1 kB  | 6.70%   |
| EventString            | 27.7 kB  | 4.22%   |
...
+------------------------+----------+---------+
|         TOTAL          | 657.5 KB | 100.00% |
+------------------------+----------+---------+
```

### count

```
traceutils breakdown count <input>
```

Example output:

```
+------------------------+--------+---------+
|       EVENT TYPE       | COUNT  |    %    |
+------------------------+--------+---------+
| EventGoStartLocal      |  61612 | 41.88%  |
| EventGoSched           |  58859 | 40.01%  |
| EventHeapAlloc         |  11801 | 8.02%   |
| EventGoPreempt         |   2354 | 1.60%   |
| EventGoSysCall         |   1791 | 1.22%   |
...
+------------------------+--------+---------+
|         TOTAL          | 147111 | 100.00% |
+------------------------+--------+---------+
```

### csv

```
traceutils breakdown csv <input>
```

Example output:

```
Event Type,Count,Bytes
EventGoStartLabel,596,4007
EventHeapAlloc,11801,81561
EventProcStart,1198,5381
EventGoStartLocal,61612,193841
EventGoSysCall,1791,7069
...
```

# License

MIT