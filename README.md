# Port-Scanner

How to use exmaple:

```
func main() {
	var scanner *API.Scanner
	scanner = new(API.Scanner)
	scanner.Ip = "172.217.13.206"
	scanner.MinPort = 0
	scanner.MaxPort = 2048
	ports := scanner.Scan()
	fmt.Println(ports)
}
```
