package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		result   string
	}{
		{
			name:     "hello world",
			expected: `--color=red "hello world"`,
			result: ``,
		},
		{
			name: "{Hello & There #}",
			expected: "--color=red {Hello & There #}"
			result: `   __  _    _          _   _                                _______   _                                    _  _    __    
  / / | |  | |        | | | |                 ___          |__   __| | |                                 _| || |_  \ \   
 | |  | |__| |   ___  | | | |   ___          ( _ )            | |    | |__     ___   _ __    ___        |_  __  _|  | |  
/ /   |  __  |  / _ \ | | | |  / _ \         / _ \/\          | |    |  _ \   / _ \ | '__|  / _ \        _| || |_    \ \ 
\ \   | |  | | |  __/ | | | | | (_) |       | (_>  <          | |    | | | | |  __/ | |    |  __/       |_  __  _|   / / 
 | |  |_|  |_|  \___| |_| |_|  \___/         \___/\/          |_|    |_| |_|  \___| |_|     \___|         |_||_|    | |  
  \_\                                                                                                              /_/   
                                                                                                                         
`,
		},
	}

	content, _ := readFile()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Runner(content, tt.flag, tt.subString, tt.sentence)
			if got != tt.result {
				t.Errorf("Expected:\n%q\n\nGot:\n%q", tt.result, got)
			}
		})
	}
}
