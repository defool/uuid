package uuid

import (
	"math/rand"
	"net"
	"os"
	"sync/atomic"
	"time"
)

var (
	maxQPS int64 = 62 * 62 * 62 //238328
	lastID int64
	ipPid  []byte
)

func init() {
	rand.Seed(time.Now().Unix())
	lastID = rand.Int63() % maxQPS

	pid := os.Getpid()
	ip := getIP()
	var ipv int64
	for _, v := range ip {
		ipv += int64(v)
		ipv *= 256
	}
	ipv *= 256 * 256
	ipv += int64(pid)
	ipPid = make([]byte, 9)
	Base62Encode(ipv, ipPid)
}

// UUID returns unique string by timestamp / IP / PID / autoincrement ID
func UUID() string {
	bs := make([]byte, 18)
	t := time.Now().Unix()
	Base62Encode(t, bs[0:6])
	copy(bs[6:15], ipPid)

	id := atomic.AddInt64(&lastID, 1)
	atomic.CompareAndSwapInt64(&lastID, maxQPS, 0)
	Base62Encode(id, bs[15:18])
	return string(bs)
}

// RandID returns a string base on the timestamp / rand int64
func RandID() string {
	bs := make([]byte, 14)
	t := time.Now().Unix()
	Base62Encode(t, bs[0:6])
	Base62Encode(rand.Int63(), bs[6:10])
	Base62Encode(rand.Int63(), bs[10:14])
	return string(bs)
}

// Rand returns a given-size string base on the timestamp / rand int64
func Rand(size int) string {
	bs := make([]byte, size)
	for i := 0; i < size; i += 6 {
		end := i + 6
		if end > size {
			end = size
		}
		Base62Encode(rand.Int63(), bs[i:end])
	}
	return string(bs)
}

func getIP() []byte {
	conn, err := net.Dial("udp", "10.0.0.1:80")
	if err != nil {
		return make([]byte, 4)
	}
	defer conn.Close()

	addr := conn.LocalAddr().(*net.UDPAddr)
	if addr == nil {
		return make([]byte, 4)
	}
	return []byte(addr.IP.To4())
}
