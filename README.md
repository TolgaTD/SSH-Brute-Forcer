SSH Brute Forcer

Overview
This is a simple SSH brute force tool written in Go. It attempts to connect to a specified SSH server using a list of provided usernames and passwords.

Features
•	Multithreaded password guessing
•	Ability to specify username, target IP address, and port
•	Uses a password list stored in a text file

Requirements
•	Go installed on your system

Usage

1-Clone the repository:
•	git clone https://github.com/TolgaTD/SSH-Brute-Forcer.git

2- Navigate to the project directory:
•	cd SSH-Brute-Forcer

3- Build the executable:
•	go build -o ssh_brute_forcer main.go

4- Run the program:
•	./ssh_brute_forcer

5- Follow the on-screen prompts to enter the required information (username, target IP address, port, and password file path).
•	Input the Username: admin
•	Input the Target IP and Port (e.g., IP:PORT): 192.168.1.100:22
•	Input the Password text file path (e.g., C:/path/to/your/passwords.txt): passwords.txt


Contributing
Contributions are welcome! If you have any suggestions or improvements, feel free to open an issue or submit a pull request.

License
This project is licensed under the MIT License. See the LICENSE file for details.

