package in_memory

// Create a var in order to add all ips from config
var ipsToCheck []string

// StoreIps will store all ip in a var
func StoreIps(getOutput []string) {
	for _, value := range getOutput {
		ipsToCheck = append(ipsToCheck, value)
	}
}

// CheckIfIpExistsInMemory is used to search a certain ip into memory
func CheckIfIpExistsInMemory(ip string) bool {
	existsIP := false

	for _, checkIp := range ipsToCheck {
		if checkIp == ip {
			existsIP = true
		}
	}
	return existsIP
}
