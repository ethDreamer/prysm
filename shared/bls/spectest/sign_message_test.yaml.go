// Code generated by yaml_to_go. DO NOT EDIT.
// source: sign_msg.yaml

package spectest

type SignMessageTest struct {
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	ForksTimeline string   `json:"forks_timeline"`
	Forks         []string `json:"forks"`
	Config        string   `json:"config"`
	Runner        string   `json:"runner"`
	Handler       string   `json:"handler"`
	TestCases     []struct {
		Input struct {
			Privkey []byte `json:"privkey" ssz:"size=32"`
			Message []byte `json:"message" ssz:"size=32"`
			Domain  uint64 `json:"domain"`
		} `json:"input"`
		Output []byte `json:"output" ssz:"size=96"`
	} `json:"test_cases"`
}
