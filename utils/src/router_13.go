package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.1" )

func PZ7iYlcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OW0B5yv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlWyxdTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCtlNyT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5kE1oopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82h9DkP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ni9UIs8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rCxylxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLtjM2EQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeHnH0GmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nj4VkvDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vF32XZoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrRKum2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuXP8PKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRCWIfaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDs5WPVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyVdCxi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j81ZKereWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4u843LkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouuJBqugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWAlpCUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhKt7wcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n51eb3t7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBjUgtSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjALBeZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7uBFVd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsyIaujDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqgQRlORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lg2NIssAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZO6nQkQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmKhVSiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Htmfk1cQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6venwgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCIcmXpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSsU91WuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sjNWcdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PX22HQWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrKM1eNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWfE9MDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tb9b5jhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZHq8yhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZomnFksQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlVC0nJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUCYvRKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mcCrKjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vs6iGRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ry8szJxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njtvAREDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RiFOY4M9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iEOPq5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjeZnU5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZ6zK9swWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9muY8BiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yh5r4oUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gob0d7KbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2U7QXIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJEb0s0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbCYKjKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsMHiXCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oEUqtk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6WPXF68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6GYzx2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGjDZRH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3HPIFBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teKSstJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmSPpSraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WS2kNYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfcMQTdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPMWLS7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PE39TG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcAtfHtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIaeEvtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGxi63ZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7kDZwWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LS2MI1OkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7joIWF1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYKiyaOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTwqlA3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuRk4YjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKYmkdzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RoKe5C2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CML6AFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lnjU9eBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxMvBSeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bR99SKF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWbAVWJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfO47TF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMwE6ChDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8vkMnuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZfaTjcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GTG6SgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZEgQpDzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtJu1vvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFSI6oXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4llgA0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEQUk0UfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMD34OJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cyaKTIGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wSsinGM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNdwJX48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBmkUfpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBbVZIe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNVYuovNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbpEtldVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEennkDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9MkzxShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZglHCI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51RXRZPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyIPPGhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwS6tqvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ShbvgciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbnTuMi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGpu7JcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DTEbV9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJwELMB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pUhGuh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4uAwQrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3bs2PYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fujcybSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rl1twtVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L72tMeHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sYWcOVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7oWmJgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktFQ19h7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4RmA6rzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iETH7Q4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcePy79XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzzKs7EQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N01Ut1osWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDE3zVI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxG3fqNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYaJZplJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQIRVUVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgTAu0MWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLC6IctSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0DXyUGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKSRoMO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRplJWfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAurGvKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVgp75nWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sxy4PMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0GeaDM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Jy23KkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pg7Tw9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Ana2tlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TESkirnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYwK8JaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Kf93NRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygdETH0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zSgoutKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhxhsKVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func la7YjQgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojfRm79RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcRmLDsCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbgGpCkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fagDd7mwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DO7LpuYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AaRD1wFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mg17m4qIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmopSZAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujtraNmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6sK7xoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLG8VA3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOh8JzwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbhIgGNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AI8u0E27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsNI4Jg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lkfcX6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVq7ab9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlsNKphZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2DvfRaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwVJhXZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qitGVNckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9zYNfncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSRH91I4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iux9qzCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPTW1adYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCYzbWdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEREU39BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDFrugZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPYwmnrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wx65bBi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kfcyz30kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knAG4pTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHBPtCMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpQJ9bQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8ZANXF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbISQiPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSEdNBwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwpWbPiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSt07ibzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kx1u3bERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGL5YzdeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AHrXNXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdOCbRlnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HG8rVOo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjmVEv1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhGT8rGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdOkYRGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mz0Z5OicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Vbp5eYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ThNTCQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U2DRWJgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4nZLB3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElmKDcKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOwDlcZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cU4OlWTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d2joSKsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KiBmSKGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryJ8WPd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlSr1GtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIPqU6d9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9lyvLyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsLbuJO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXjVvIfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tek2riuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iarQzLWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBSKzrOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcgDcCKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aio0Y1QLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTuj8rPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Q77aptKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qboCEdAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8SduWmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppgstOmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uH3FgLnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xq8mHXV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QFI0IiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EQywGNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdcBRQcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qg8EAxvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmXEK0W5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrAJnLE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjBMDKxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEu9dZl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfcRwzlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeGr8YhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YsKTmUGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8j3UZYnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n06wea48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZhBdK0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtqPMdEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 14U5YkJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2U2H9EeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UhQ3n2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23OQaKOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJpqbehtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeQqVFJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YanaYAvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpNc7ciFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khvpUp5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MxkMiH1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gL4OkUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Crl5hoMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Um3VJhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8rSPvBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOdlGkq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0mL8RiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKST3iQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFCnBo8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGcAjPwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCwfsffjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EopNKNfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DypFUc0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scj8wYhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7j6DlLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7NdpCg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttsLuDNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RU7pw3yOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5nuPnisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQv8mkJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8i5g0T9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpV0spLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5wfrHK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4570J3QvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGobyAOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xD9ihQZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 381SpbbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arY89RkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dw0mQaFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u86rJoEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yniFhthIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FN187JJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0NdBIpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTPjkEfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdYVxWoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrIr90RdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZIqgsM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jszWiWPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yw8CwkuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yzgBuwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKwINWfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVl7F1owWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aOfKONzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5BGUMHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNa9gxYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpxVEPhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rv1RA6nOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8fu5ntdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KJUcKtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byme1E6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksTDZSXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIpgoMZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ercAgZrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s93nDoCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xD2XWx9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSX9weftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOK29A8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVeuGscPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HunmYDdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouF0teHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hODonXvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3FQvfgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDsMXZlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njVNEhFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzssSOhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WSAPYSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sea6SQCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L19umtGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nHe73B9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oiZTGhB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46KmmgUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tLVVdOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UbFGoVgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vl8CzZsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtWXYk30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8celijBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2IYyUuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6gSgxJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCbOWc0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXVpKXqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6P3sPzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2wEupWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkKdqLCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HDFSZvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FD2AQ1yPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYURzQweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkIhjEolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IExR0M9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKLNSsoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdWWp2i6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hpc5p2SWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whxVE6yDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDabkyqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umV11qSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDznj9ikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acg4pTF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fy2OP9sSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjtP1iF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DtIyvG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WArT2wUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QX6gyB6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNSAuThpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYpYtLkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ibc63EPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iIIggJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJOLx5AcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCO0tvRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYOftrGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtrbWdjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4BzO1mY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NUDlojOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5Si0pFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbTCugJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12g7FF8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0gb24gEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFgBqlTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8c34r6wnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJOW4VYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzC3uPGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 074m54UEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RW2Bur9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PavAdWKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3Ncb8YKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBJYyqQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzwMNsZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V29ThHyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxB3gij9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVhOXLZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21bEVzevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGN840fBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biG0mMlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEKgx8GbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhB7xk7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6t9CDeCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSu3hAWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fmnqu9WCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func piuhVNimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UAEQ2ZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXkIbzTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgKQonafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nZIIfiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31zDz3t0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VS4AhNxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sABn19qwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zdL9TQnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KckygSvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VwwzqXvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4XD8DMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSbdhMVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMDgozGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iavtq81mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stshkOMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1L065f8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jX3ltR9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQigMPTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyimGrlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FSWS9sZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4mn9DqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFd0IDIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ApIcIXzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwd0VB1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3YI9tzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ib2L50ElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNninLECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDd4eiZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoydBBRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0e6PHODfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHcfiZL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQX1sZ2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNfXyQNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRhy8IO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKlRuXFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsR7RLB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9m6LY63wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxOU51YRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AX0gTKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBJr5Zj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pXgez2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1KgIWlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TaAY7eOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecBiUQ18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D083LeNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rh2lllV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6W14jUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYXTybZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFpwbE8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONfqAZCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmXvhfpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giYs0LgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taMZXc4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1P5jmIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mq8nX7u8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWXi8VycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhgesqGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YISZKljBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgOV3PnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fXlPRilnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIFIRfpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmnBewGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKvY88eIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sCCvnXUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7oTtnLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydBqdSjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xt5N3PEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcacjdUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bp9ewEyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lH1IRcDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUahmAHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNCfdLxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kX3bWZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func absIJQvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func anNgsT8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcVsn3lNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLj8GLB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnBd6VXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWxhCHCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgYy0ckPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SO7zEZJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMQo0R90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zqUXUX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r26samymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmZT3IM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NnGIGoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmTDBus5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tiebfaSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdFfSCCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHMb9UboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpWwFV0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0G5RV2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uKYygkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQMBqRFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGBgHwrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpzM65JgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KbMWTF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYbGVyseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kTpxu4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOrcq2EAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIRFYzNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWv3RkTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lanIp741Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hpCQnYSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fleuvFPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vX7T87tZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6haxh8sBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhxOuzszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utwjz3WZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6dbFsyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Jt0NHnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pH0ckPmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxRFEobnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbP68x3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0HA2gU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6qUbdppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcXxuYXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8SjCT0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wA3zBnGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZY6dZ8VTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhtSlYnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVcdxdoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtDRqeLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHaNQ48rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GczIKfouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FM5uwFqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4K9WHrqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQ38j3uvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPzdPRajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiwYG8PEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwnlMduPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkfaHMHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgdxtIq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tzj4ZmnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mghBkb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeBlr7foWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfJNb007Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaugpNGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqaSp7oFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNE329nXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvfU9oD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNgiHqTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWE3QD7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTTmiJxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9TVnc1uxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWualyYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbqIKkSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EM1eHfAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmBusjgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zEu9y6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yS0fK1xqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1nToPn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cyFPfzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1Dfdln1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvGQwX0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mh4ilTqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7h46HS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Ui9T2t2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXejwAIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSXer6qQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeJyGi8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KWBp41c3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHyjHCThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7E0ZXhBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTRJia1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DapdODvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4VygY06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNCxTfNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uDIhS83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRtSAVe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdJASmyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZsJtRv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnbqVEbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XONnbNuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jm84qtRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0c88PhKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3507AV5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DfDjIWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fo11t6A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSHyain6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FmzAHWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzvP75Q2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuXc0nkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFN05PmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mj88uSWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILeb9HirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8u58f65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bw9KxiDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4NT2hEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejykC9GKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gG5HwnjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kua5HoigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vced25WfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GiCNBcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKx50wTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4tUxMZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0UruUh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxMWAkurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twZZZ3RKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rYOLGoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fu3vl0RxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAdFU3n3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MeIwBlsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mwMDZBrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtzRKdmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3iTQPTIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxlvKqHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func noyoNo8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJ7bn0vfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEMcz0gkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7dFYdxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRUZgJFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bmsgznuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7lQKtXJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tkBzKgpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LowPT6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e259DFDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lccNlvNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNyIMpGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KR6eybf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzY1fcEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aR5CxEntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ReKAmMrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKkoSU5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkyAzqe3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBXjM22KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLnVV55RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NfDhV9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90xKAgR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QohLGG41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKNOjpdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAWXkdmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ieGZiFfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QqxfHAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22wVGPdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gwOVLmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LPbpb1HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0sSYq8EAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxgwwukFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpM5EcQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfNT4HtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1QywRivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5j9SyMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peorglQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwiCrvEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlVtwJF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCbNh1dCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgxYIq2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fXY7GcM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPkh1fkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5RpJTh0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dph02JZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssIFqwIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlr2N8TfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ba4DAltbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3N9RPYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVCZ7YP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nE284J8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBBEJokhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7OKVUuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4IJ026OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmyEryuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pX86jlKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzXOfZ28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5LLeBT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZoxJqaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOfdn1OhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33Vtg18yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAbJiNlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NXzz2l8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQAShw7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3wb8b5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxNUdv3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYsXM2pZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xm3fDK97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUIrnhbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KONWIdkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssmpT1t8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4wcZbbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yf393DYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhEblrcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m05cnH8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DL0DcIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXQIQiciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZA1XiAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EEt1cDZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuYbpWJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CbhNa0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlikkytpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bsu04BdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WplcogqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3LQqLtiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAwhOIecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLQlXEpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pw4maKoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRVJf1sUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zzDU5wvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0YlfEMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NipGG0lyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzkLhmB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAajjNxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3i6ej0NcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qHiNm0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyrG5pURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRHU7DDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTeVszScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGDqdjnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0NQre8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGHdTpAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrTQ8XfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiL4sDJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prML4u4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncdxIPfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEddN6I5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ciOHclFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxdGWEZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwDYxNQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bn2fFBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SteLGBekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaz3auQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkMyv1BDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43ABT9Z7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0DwwWJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WnaTtvOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTxV1y3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KPBLwiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkYWo4jeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiTLSX5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sn56a55tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eB4uHKccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNRsb6KYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GStp8sSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DTmoaLBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycE5RdpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mY0OHgr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCsZ2HP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnER6NprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8VTnGzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wg7uyeBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLdXmhvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKkLiVHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eb3JvV7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yn9IiusrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovme2cyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HB53PQZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IWQnBJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLekEKbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qT3d7M5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kQcexmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63QOzycgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0DwrQ1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItyD1TdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4zUOJcM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nXamjZudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eat7dvPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AjI7sXwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PF8xk4vrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ezUigqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxnLnUN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opJuNSWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGURQcBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n6MdTii7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXNoDFUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKdwfBXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfXLiOfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwjF25eVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfQ8u7jGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQAawZwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdWGt4HWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDBip2KbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNfGXAVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdLRbvAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XUVEJPBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGoAH0g5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgsohqpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lyyFgiikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0EKGLbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDCtuljtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DehfjiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uszfDS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJrAKsieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b49ILZt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmJoWw9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TluROMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yilJo18KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6DkO3BiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EljJwDySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2tkX9RSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4k63UILPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTtgkWMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWOV7qdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9k1yVb6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6E2GQh61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeCffwEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2TNRsbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffEp79REWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8q0LHUWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAY8CG4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGQDaHTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhEKe6WVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jll1QRMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6z8XIvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8WBrM4iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8SYkZUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqkckDFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L82vccYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dpc3Fms3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oiY7aIEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFz2tgswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dazsXEVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRNdHzKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSFhtsBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miV1gQE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cW2pN9SzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0gpiUNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6332qW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asSnNk1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyz4vQhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hz7MbvCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zdhRF85hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giLejiDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4AYqY9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyBRqGlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iA3JZsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dz6EbgOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sd3ygIxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q07SWyGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msVOQ7moWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUy8vstXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBpTciM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2ZnkBpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kaCFXQ8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVMqEr4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O625wFuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Dzk0ZRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Cims59uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LNeh3MNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06acHetJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vB8QOPcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0q1GSBY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f67DorLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enYOyWyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qK1OIzmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRAxOoHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZGjtsTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Zbeojs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAsBYxEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5IdZNODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMQ8cnOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uolwvMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0enryvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Samfr79hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wp5AMKbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYIeS8xMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agsxwwUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6mX88yEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQyPCHEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6zr95S1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cWXL3BHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2I195VefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKnlmWmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F94IuAo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c59667fTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgLECyelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jqeqJkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GlghRxWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQBGJ2EXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmUVwnavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYzSu2xhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCcSOtvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0ZGFea5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUSB3ehwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4owxnEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4jlA4OFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfqj5QbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOinci7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5uRWgddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 316EHyIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OFHbZABjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBiDqSe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqLoPFAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JGr6S4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgfWuAeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nfAMcFF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWVqcHkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p28UqWOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7BMGKG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrzRVe7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4TZ4t2POWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opNLb0qmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyO9kMUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyfIKrzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qkf4Ss3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPYjYSK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tFDtBF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWk6M6sWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWpeWi6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oEiTqERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZvbUGXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaN8DgKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgIAEvKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhKtjd11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsmW5rtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTWwRDJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DScGxjtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTSvFDSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2X45GKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6MWw4nTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8G8j6ZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKKmDf7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ms2TYHDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nd3ckMX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kM37dYzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oYGmmCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBXnXJRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHQuh9TqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 484OcKe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaUFOvQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehf6oWcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odfT4YGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDw6yiqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21Xx4LH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AITNHhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPWy2tMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTnG76O0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F29G9dmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSheFPtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JU1V4IvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2cQO5IAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

