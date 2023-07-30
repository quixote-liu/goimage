package image

import (
	"log"
	"testing"
)

func TestImageRotate90(t *testing.T) {
	instance := New("./test/test_rotate.png")
	err := instance.CWRotate90("./test/test_rotate_90.png")
	if err != nil {
		log.Fatalf("rotate 90 failed: %v", err)
	}
}

func TestImageRotate180(t *testing.T) {
	instance := New("./test/test_rotate.png")
	err := instance.CWRotate180("./test/test_rotate_180.png")
	if err != nil {
		log.Fatalf("rotate 180 failed: %v", err)
	}
}

func TestImageRotate270(t *testing.T) {
	instance := New("./test/test_rotate.png")
	err := instance.CWRotate270("./test/test_rotate_270.png")
	if err != nil {
		log.Fatalf("rotate 270 failed: %v", err)
	}
}
