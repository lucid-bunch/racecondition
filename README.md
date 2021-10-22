# racecondition

```bash
➜  racecondition git:(master) go run main.go 
Running version 1
0  SD: success in 74ms
0 SB: success in 141ms
0 SC: success in 253ms
0  SA: success in 298ms
Duration: 298ms
=======================
0  SD: success in 19ms
0 SC:                    RECURSE ERROR - recurse
0 SB: success in 140ms
1  SD: success in 20ms
1 SB:                     NORMAL ERROR - return
Duration: 217ms
=======================
1 SC: success in 118ms
0  SA: success in 72ms
0  SA: success in 298ms
0  SD: success in 178ms
0 SC: success in 195ms
1  SA: success in 282ms
0 SB: success in 249ms
Duration: 254ms
=======================
0 SC: success in 159ms
0  SA: success in 202ms
0 SB: success in 252ms
0  SD: success in 292ms
Duration: 297ms
=======================
0  SA: success in 41ms
0 SC: success in 96ms
0 SB: success in 133ms
0  SD: success in 174ms
Duration: 175ms
=======================
0 SC: success in 1ms
0  SD:                    RECURSE ERROR - recurse
0  SA: success in 72ms
1 SB: success in 31ms
1 SC: success in 81ms
1  SD: success in 110ms
0 SB: success in 219ms
1  SA:                     NORMAL ERROR - return
Duration: 282ms
=======================
0  SD:                     NORMAL ERROR - return
Duration: 34ms
=======================
0 SC: success in 34ms
0 SB: success in 64ms
0 SB: success in 124ms
0 SC: success in 153ms
0  SA: success in 224ms
0  SA: success in 262ms
0  SD:                     NORMAL ERROR - return
Duration: 276ms
=======================
0 SC: success in 38ms
0  SA: success in 70ms
0 SB: success in 112ms
0  SD: success in 194ms
Duration: 199ms
=======================
0 SB: success in 32ms
0  SD: success in 94ms
0  SA: success in 141ms
0 SC: success in 182ms
Duration: 186ms
=======================
0 SC: success in 24ms
0  SA:                    RECURSE ERROR - recurse
0 SB: success in 89ms
1  SA: success in 75ms
1 SB: success in 90ms
1 SC: success in 134ms
1  SD: success in 298ms
Duration: 362ms
=======================
panic: send on closed channel

goroutine 47 [running]:
main.fetchVersion1.func4()
	/github.com/lucid-bunch/racecondition/main.go:80 +0x126
created by main.fetchVersion1
	/github.com/lucid-bunch/racecondition/main.go:75 +0x325
exit status 2
```
