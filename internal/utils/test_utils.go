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

func GetNIOSDtcCertRef() string {
	return os.Getenv("NIOS_DTC_CERT1_REF")
}

func GetNIOSDtcCert2Ref() string {
	return os.Getenv("NIOS_DTC_CERT2_REF")
}

func GetNIOSADAuthServiceRef() string {
	return os.Getenv("NIOS_AD_AUTH_SERVICE_ACTIVE_DIR_REF")
}

func GetNIOSADAuthServiceRef2() string {
	return os.Getenv("NIOS_AD_AUTH_SERVICE_ACTIVE_DIR_TEST_REF")
}

func GetNIOSGridMasterHostName() string {
	return os.Getenv("NIOS_GRID_MASTER_HOSTNAME")
}

func GetNIOSGridMemberHostName() string {
	return os.Getenv("NIOS_GRID_MEMBER_HOSTNAME")
}

func GetNIOSNotificationRestEndpointRef() string {
	return os.Getenv("NIOS_NOTIFICATION_REST_ENDPOINT_REF")
}

func GetNIOSPxgridEndpointRef() string {
	return os.Getenv("NIOS_PXGRID_ENDPOINT_REF")
}

func GetGSSTSIGCertRef() string {
	return os.Getenv("NIOS_GSS_TSIG_CERT_REF")
}

func GetSyslogEndpointRef() string {
	return os.Getenv("NIOS_SYSLOG_ENDPOINT_REF")
}
