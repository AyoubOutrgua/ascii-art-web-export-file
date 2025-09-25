package functions

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func HashFiles() error {
	hashShadow := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	hashStandard := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	hashThinkertoy := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	standardData, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		return err
	}
	standard := fmt.Sprintf("%x", sha256.Sum256(standardData))
	shadowData, err := os.ReadFile("banners/shadow.txt")
	if err != nil {
		return err
	}
	shadow := fmt.Sprintf("%x", sha256.Sum256(shadowData))
	thinkertoyData, err := os.ReadFile("banners/thinkertoy.txt")
	if err != nil {
		return err
	}
	thinkertoy := fmt.Sprintf("%x", sha256.Sum256(thinkertoyData))

	if standard != hashStandard {
		return fmt.Errorf("invalid hash for standard.txt")
	}
	if shadow != hashShadow {
		return fmt.Errorf("invalid hash for shadow.txt")
	}
	if thinkertoy != hashThinkertoy {
		return fmt.Errorf("invalid hash for thinkertoy.txt")
	}
	return nil
}