package reverse

import "testing"

func TestValidateForManagementAllowsDisabledLocalConfig(t *testing.T) {
	cfg := &Config{DBFilePath: "reverse.db"}
	if err := cfg.ValidateForManagement(); err != nil {
		t.Fatalf("disabled config should be valid for management: %v", err)
	}
	if err := cfg.ValidateForStart(); err == nil || err.Error() != "enable HTTP or DNS server first" {
		t.Fatalf("ValidateForStart() error = %v, want disabled-server error", err)
	}
}

func TestValidateForStartAllowsDNSOnly(t *testing.T) {
	cfg := &Config{
		DBFilePath:      "reverse.db",
		DNSServerConfig: DNSServerConfig{Enabled: true},
	}
	if err := cfg.ValidateForStart(); err != nil {
		t.Fatalf("DNS-only config should be startable: %v", err)
	}
}

func TestValidateForStartAllowsRemoteWithoutLocalServers(t *testing.T) {
	cfg := &Config{
		Token:        "token",
		ClientConfig: ClientConfig{RemoteServer: true, HTTPBaseURL: "http://127.0.0.1:88"},
	}
	if err := cfg.ValidateForStart(); err != nil {
		t.Fatalf("remote config should be startable: %v", err)
	}
}
