package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var addresses = []int64{
	0x1F08D,
	0x1F091,
	0x1F48D,
	0x1F491,
}

var snAddreses = []int64{
	0x0001F023,
	0x0001F423,
}

func main() {
	// Ask for user input
	var filePath = "MEMORY_G3.bin"
	var speed int
	var changeRegion string

	fmt.Print("Enter the speed (1 to 255): ")
	fmt.Scanln(&speed)
	if speed < 1 || speed > 255 {
		log.Fatalf("Invalid speed value: %d (should be a number between 1 and 255)", speed)
	}

	// Ask if user wants to change the region
	fmt.Print("Do you want to change the region (Y/N)? ")
	fmt.Scanln(&changeRegion)

	if changeRegion != "Y" && changeRegion != "N" {
		log.Fatalf("Invalid input for region change: %s (should be 'Y' or 'N')", changeRegion)
	}

	copyPath := filePath + ".patched.bin"

	if _, err := os.Stat(copyPath); err == nil {
		fmt.Printf("Deleting existing copy: %s\n", copyPath)
		if err := os.Remove(copyPath); err != nil {
			log.Fatalf("Failed to delete old copy: %v", err)
		}
	}

	err := copyFile(filePath, copyPath)
	if err != nil {
		log.Fatalf("Error copying file: %v", err)
	}
	fmt.Printf("Created copy: %s\n", copyPath)

	file, err := os.OpenFile(copyPath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error opening copy: %v", err)
	}
	defer file.Close()

	for _, addr := range addresses {
		_, err := file.Seek(addr, 0)
		if err != nil {
			log.Fatalf("Error seeking to address 0x%X: %v", addr, err)
		}
		_, err = file.Write([]byte{byte(speed)})
		if err != nil {
			log.Fatalf("Error writing to address 0x%X: %v", addr, err)
		}
		fmt.Printf("Written value %d to address 0x%X\n", speed, addr)
	}

	if changeRegion == "Y" {
		for _, addr := range snAddreses {
			_, err := file.Seek(addr, 0)
			if err != nil {
				log.Fatalf("Error seeking to address 0x%X: %v", addr, err)
			}
			_, err = file.Write([]byte{0x43})
			if err != nil {
				log.Fatalf("Error writing to address 0x%X: %v", addr, err)
			}
			fmt.Printf("Written value 0x43 to address 0x%X (US SN)\n", addr)
		}
	} else {
		fmt.Println("Region change skipped.")
	}

	fmt.Println("Done.")
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
