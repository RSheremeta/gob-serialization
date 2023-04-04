## Analysis of results run 

This is my extended analysis of all runs.
Feel free to reach out if you spotted an inconsistency or a mistake.

#### Spoiler
I've omitted highlighting ``Medium`` and ``Big`` structs' results, as the most interesting is in the smallest and the biggest stuff here, at least IMHO.

### Single struct

1. ++Encoding++:
    **Tiny**:
        - Gob loses JSON here, as it takes ~3x more time (3421 ns/op vs 1304 ns/op in JSON), ~6x more memory (1136 B/op vs 176 B/op in JSON), ~20x more memory allocations (20 allocs/op vs 1 allocs/op in JSON)
        - the *slowest* performance is shown by YAML: 6817 ns/op, 6776 B/op, 23 allocs/op
    **Huge**:
        - Gob is the *fastest* one here:
         it's ~3x faster than JSON (1692592 ns/op vs 5292438 ns/op),
         ~2x less memory (2449044 B/op vs 5156133 B/op in JSON),
         ~2.5x less memory allocations (29982 allocs/op vs JSON's 373306 allocs/op)
        - *YAML* shown the worst results here: 100121717 ns/op, 272096372 B/op, 373306 allocs/op 
2. ++Decoding++:
    **Tiny**:
        - Gob is the *slowest* one among all. JSON is the best here:
         it decodes ~9x slower than JSON (12662 ns/op vs 1354 ns/op),
         ~27x more memory needed (6760 B/op vs 248 B/op in JSON),
         ~30x more allocations (179 allocs/op vs 6 allocs/op)
        - XML takes 2nd place - 3317 ns/op, 1481 B/op, 33 allocs/op
    **Huge**:
        - Gob  is the best while YAML is the worst here.
         Gob's processing time is 1645925 ns/op vs 12986658 ns/op in JSON, which is ~8x faster,
         YAML results of a huge single struct decoding: 88956475 ns/op, 37407429 B/op, 745485 allocs/op


### Slice of structs


1. ++Encoding++:
    **Tiny**:
        - Gob holds 2nd place, while JSON is on the top by all aspects here:
            - Gob's results: 7862 ns/op, 1392 B/op, 22 allocs/op,
            - JSON's results: 3721 ns/op,  562 B/op, 1 allocs/op.
    **Huge**:
        - Gob has shown the best results for encoding huge struct, let's compare it with YAML's last place: 
            - Gob: 75574500 ns/op, 109713207 B/op, 690111 allocs/op,
            - YAML: 4233214871 ns/op, 12425343834 B/op, 18662808 allocs/op.
            - ++summary++: Gob is ~56x faster, ~113x less memory consuming and requires ~27x less allocations per operation comparing to YAML.


2. ++Decoding++:
    **Tiny**:  
        - XML took 1st place here with such results: 5858 ns/op, 1489 B/op, 32 allocs/op.
        - the entire ranking looks like: XML, JSON, Gob, YAML.
    **Huge**:
        - XML also leads a decoding of huge struct with brilliant results:
            - XML: 3062 ns/op, 920 B/op, 15 allocs/op
            - for instance, 2nd best result goes to Gob with 67565967 ns/op, 63096762 B/op, 1107478 allocs/op.
            - just compare the numbers XML vs Gob: XML is ~22k times faster, ~68k times less memory usage and ~73k less allocs per operation.


### Slice of pointers to structs

The results are mostly quite similar to the slice of plain structs, so I avoided hightlighting it.
Feel free to analyze it on your own.


### Single complex map type
**Reminder: comparing Gob vs JSON only!**

++Encoding++:
    Gob is better than JSON in terms of speed as it is ~3x faster while the other aspects are pretty same (109720182 B/op vs 127879850 B/op) and (690216 allocs/op vs 611148 allocs/op)

++Decoding++:
    Gob requires ~10x less time, ~2x memury usage, ~1.5x less allocations to decode a complex map rather than JSON.


### Slice of a complex map type
**Reminder: comparing Gob vs JSON only!**

++Encoding++:
    Gob is *quite better* than JSON in terms of time needed - ie - ~3.5x faster, while memory usage is pretty same.
    However, the amount of allocations per operation is quite higher: 6901129 for Gob vs 6111405 for JSON.
++Decoding++:
    JSON shows *way worse* results here: ~10x slower, needs ~2.5x more bytes and ~1.5x more allocations per operation rather than Gob.


### Slice of pointers to a complex map type
**Reminder: comparing Gob vs JSON only!**


++Encoding++: 
    Gob is ~3.5x faster, needs 20% less memory, but got slightly higher amount of allocations per operation needed: 6901129 vs 6111407 for JSON.
++Decoding++:
    Gob has shown *way better* results comparing to JSON here: ~10x faster, ~2.5x less memory and ~1.5x less allocs per operation needed.
