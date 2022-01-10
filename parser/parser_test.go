package parser

import (
	"testing"
)

func TestParsesLineToEntry(t *testing.T) {
	lines := []string{
		"vígmálugr|adj|vígmálugr, adj. tilbøielig til at tale omKamp, Drab; hann var maðr - víg-málugr ok íllmálugr Háv. 219.",
		"gambansumbl|n|gambansumbl, n. stort Gjestebud. Lok. 8.",
	}
	result := ParseLines(lines)

	if result[0].headword != "vígmálugr" {
		t.Error("Did not return expected content. Received", result[0].headword, "expected ", "vígmálugr")
	}

	if result[0].partOfSpeech != "adj" {
		t.Error("Did not return expected content. Received", result[0].partOfSpeech, "expected ", "adj")
	}

	if result[0].definition != "vígmálugr, adj. tilbøielig til at tale omKamp, Drab; hann var maðr - víg-málugr ok íllmálugr Háv. 219." {
		t.Error("Did not return expected definition.")
	}
}
