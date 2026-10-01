package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"wg.dde1ta/files"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
     // strings.Builder is faster and more memory-efficient than string concatenation (+)
     var sb strings.Builder
     sb.Grow(length) // Pre-allocate the required memory

     for i := 0; i < length; i++ {
             // Pick a random index from the charset and write it to the builder
             randomIndex := rand.Intn(len(charset))
             sb.WriteByte(charset[randomIndex])
     }

     return sb.String()
}

func main() {
	filePath := "data/test.txt"
	var wg sync.WaitGroup

	randomIdentifer := generateRandomString(5);

	// Let's spawn 5 writers and 10 readers
	numWriters := 10
	numReaders := 50

	// Launch Writers
	for i := 1; i <= numWriters; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// Each worker gets its own FileManager
			fm := files.NewFileManager(filePath)
			
			// Sleep for a random time (0-50ms) to scramble the start order
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

			content := fmt.Sprintf("Data written by Writer %d | ID: %s", workerID, randomIdentifer)
			
			err := fm.Write(content)
			if err != nil {
				fmt.Printf("❌ [Writer %d] Error: %v\n", workerID, err)
				return
			}
			fmt.Printf("✅ [Writer %d] Successfully wrote | ID: %s \n", workerID, randomIdentifer)
		}(i)
	}

	// Launch Readers
	for i := 1; i <= numReaders; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			fm := files.NewFileManager(filePath)
			
			// Sleep randomly to interleave with the writers
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			
			stuff, err := fm.Read()
			if err != nil {
				fmt.Printf("❌  Error: %v\n", err)
				return
			}
			fmt.Printf("✅ Read content: '%s'\n", stuff)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}