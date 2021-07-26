# Port-Scanner

How to use exmaple:

```
func main() {
	var scanner *PortScanner.Scanner
	scanner = new(PortScanner.Scanner)
	scanner.Ip = "172.217.13.206"
	scanner.MinPort = 0
	scanner.MaxPort = 2048
	ports := scanner.Scan()
	fmt.Println(ports)
}
```
