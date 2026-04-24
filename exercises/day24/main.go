package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func hashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func hashMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func generateSecureKey(length int) string {
	key := ""
	for i := 0; i < length; i++ {
		key += fmt.Sprintf("%x", i%16)
	}
	return key
}

func verifyHash(original, hash string) bool {
	return hashSHA256(original) == hash
}

func main() {
	fmt.Println("=== Day 24: Cryptography and Security ===")

	fmt.Println("\n--- SHA256 Hashing ---")
	data := "hello world"
	hash := hashSHA256(data)
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("SHA256: %s\n", hash)

	fmt.Println("\n--- MD5 Hashing ---")
	md5Hash := hashMD5(data)
	fmt.Printf("MD5: %s\n", md5Hash)

	fmt.Println("\n--- Hash Verification ---")
	if verifyHash(data, hash) {
		fmt.Println("Hash verification: PASSED")
	} else {
		fmt.Println("Hash verification: FAILED")
	}

	fmt.Println("\n--- Secure Key Generation ---")
	key := generateSecureKey(32)
	fmt.Printf("Generated key (32 bytes): %s\n", key)

	fmt.Println("\n--- Multiple Hashes ---")
	passwords := []string{"password123", "secure_pass", "test_pwd"}
	for _, pwd := range passwords {
		h := hashSHA256(pwd)
		fmt.Printf("Password: %s -> Hash: %s\n", pwd, h[:16]+"...")
	}

	fmt.Println("\n=== Day 24 Complete ===")
	fmt.Println("Next: Learn about memory management on Day 25.")
}
