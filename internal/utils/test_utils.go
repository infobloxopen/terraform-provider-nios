package utils

import "os"

func GetNIOSCACert1Ref() string {
	return os.Getenv("NIOS_CA_CERT1_REF")
}

func GetNIOSCACert2Ref() string {
	return os.Getenv("NIOS_CA_CERT2_REF")
}

func GetNIOSCACert1Serial() string {
	return os.Getenv("NIOS_CA_CERT1_SERIAL")
}

func GetNIOSCACert2Serial() string {
	return os.Getenv("NIOS_CA_CERT2_SERIAL")
}
