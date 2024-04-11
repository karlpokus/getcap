package main

import (
	"log"

	"github.com/canonical/go-tpm2"
	"github.com/canonical/go-tpm2/linux"
)

// Replicating most of tpm2_getcap properties-variable

type Prop struct {
	Name string
	Prop tpm2.Property
}

func main() {
	device, err := linux.DefaultTPM2Device()
	if err != nil {
		log.Fatal(err)
	}
	tpm, err := tpm2.OpenTPMDevice(device)
	if err != nil {
		log.Fatal(err)
	}
	defer tpm.Close()
	props := []Prop{
		{"TPM2_PT_LOCKOUT_COUNTER", tpm2.PropertyLockoutCounter},
		{"TPM2_PT_MAX_AUTH_FAIL", tpm2.PropertyMaxAuthFail},
		{"TPM2_PT_LOCKOUT_INTERVAL", tpm2.PropertyLockoutInterval},
		{"TPM2_PT_LOCKOUT_RECOVERY", tpm2.PropertyLockoutRecovery},
		{"TPM_PT_ORDERLY_COUNT", tpm2.PropertyOrderlyCount},
	}
	for _, p := range props {
		i, err := tpm.GetCapabilityTPMProperty(p.Prop)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s: %d", p.Name, i)
	}
}
