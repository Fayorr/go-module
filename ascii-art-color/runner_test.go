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
			name: "hello\\n\\nthere",
			expected: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               

 _     _                           
| |   | |                          
| |_  | |__     ___   _ __    ___  
| __| |  _ \   / _ \ | '__|  / _ \ 
\ |_  | | | | |  __/ | |    |  __/ 
 \__| |_| |_|  \___| |_|     \___| 
                                   
                                   
`,
		},
		{
			name: "{Hello & There #}",
			expected: `   __  _    _          _   _                                _______   _                                    _  _    __    
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
		t.Run(tt.name,  func(t *testing.T) {
			got := Runner(content, tt.name)
			if got != tt.expected {
				t.Errorf("Expected:\n%q\n\nGot:\n%q", tt.expected, got)
			}
		})
	}
}
