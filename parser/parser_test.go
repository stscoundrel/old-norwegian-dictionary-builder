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

	if result[0].Headword != "vígmálugr" {
		t.Error("Did not return expected content. Received", result[0].Headword, "expected ", "vígmálugr")
	}

	if result[0].PartOfSpeech != "adj" {
		t.Error("Did not return expected content. Received", result[0].PartOfSpeech, "expected ", "adj")
	}

	if result[0].Definition != "vígmálugr, adj. tilbøielig til at tale omKamp, Drab; hann var maðr - víg-málugr ok íllmálugr Háv. 219." {
		t.Error("Did not return expected definition.")
	}
}
