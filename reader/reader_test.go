package reader

import (
	"testing"
)

func TestReadsAllLines(t *testing.T) {
	result := ReadLinesFromDictionary()
	expectedLength := 42022

	if len(result) != expectedLength {
		t.Error("Did not return expected content. Received", len(result), "expected ", expectedLength)
	}
}

func TestReadingResultHasExpectedContent(t *testing.T) {
	result := ReadLinesFromDictionary()

	expected1 := "-æringr|uten ordklasse|-æringr (af ár dvs. Aare) i sexæringr, átt-æringr, teinæringr."
	expected2 := "hjarn|n|hjarn, n. Sne som ved Tøveir og paafølgende Frost er kommen til at danne en haard ogfast Masse, Folkesprogets hjaarn, kjaadn(se Aasen 293 a47; H. Ross i Nyt norskTidskrift 1, 67; maaske af et Verbum hjarra); hjarn at ríða bæði á vötnumok mýrum Sturl. II, 12722; bæði áþá ok hjarni Flat. I, 437; svelli okhinu harðasta hjarni var steypt yfirjörð Flat. I, 43834."
	expected3 := "vígmálugr|adj|vígmálugr, adj. tilbøielig til at tale omKamp, Drab; hann var maðr - víg-málugr ok íllmálugr Háv. 219."

	if result[1] != expected1 {
		t.Error("Did not return expected content. Received", result[1], "expected ", expected1)
	}

	if result[15000] != expected2 {
		t.Error("Did not return expected content. Received", result[15000], "expected ", expected2)
	}

	if result[40000] != expected3 {
		t.Error("Did not return expected content. Received", result[40000], "expected ", expected3)
	}
}
