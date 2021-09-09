package sample

import (
	"golang-pc-book/pb"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringfromSet("Intel", "AMD")
}

func randomStringfromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringfromSet("Xeon E-22286M", "Core i9-9980HK", "Core i7-9750H", "Core i5-9400F")
	}
	return randomStringfromSet("Ryzen 7 PRO 2700U", "Ryzen 5 PRO 3500U", "Ryzen 3 PRO 3200GE")
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomGPUBrand() string {
	return randomStringfromSet("NVDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVDIA" {
		return randomStringfromSet("RTX 2060", "RTX 2070", "GTX 1070")
	}

	return randomStringfromSet("RX 590", "RX 580", "RX 5700XT")
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringfromSet("Apple", "LENOVO", "DELL")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringfromSet("Macbook Air", "Macbook Pro", "Macboook")
	case "DELL":
		return randomStringfromSet("Latitude", "Vastro", "XPS")
	default:
		return randomStringfromSet("Thinkpad 1", "Thinkpad 2", "Thinkpad 3")
	}
}
