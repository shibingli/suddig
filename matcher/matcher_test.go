package matcher_test

//EXPENSIVE

// import (
// 	"io"
// 	"net/http"
// 	"strings"
// 	"testing"
// 	"time"
//
// 	"github.com/VincentBrodin/suddig/configs"
// 	"github.com/VincentBrodin/suddig/matcher"
// )
//
// func TestParalelleMatcher(t *testing.T) {
// 	m := matcher.New(configs.Defualt())
// 	url := "https://sample-files.com/downloads/documents/txt/long-doc.txt"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		t.FailNow()
// 	}
// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.FailNow()
// 	}
// 	str := string(b)
// 	in := strings.Split(str, " ")
// 	for range 10 {
// 		in = append(in, in...)
// 	}
//
// 	t.Logf("Size: %d\n", len(in))
//
// 	start := time.Now()
// 	out := m.ParallelMatch("lorem", in)
// 	t.Logf("Found %d matches\n", len(out))
// 	elapsed := time.Since(start)
// 	t.Logf("Paralelle took %s\n", elapsed)
//
// 	start = time.Now()
// 	out = m.FindMatches("lorem", in)
// 	t.Logf("Found %d matches\n", len(out))
// 	elapsed = time.Since(start)
// 	t.Logf("Normal took %s\n", elapsed)
// }
//
//
// func TestParalelleRanker(t *testing.T) {
// 	m := matcher.New(configs.Defualt())
// 	url := "https://sample-files.com/downloads/documents/txt/long-doc.txt"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		t.FailNow()
// 	}
// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.FailNow()
// 	}
// 	str := string(b)
// 	in := strings.Split(str, " ")
// 	for range 10 {
// 		in = append(in, in...)
// 	}
//
// 	t.Logf("Size: %d\n", len(in))
//
// 	start := time.Now()
// 	_ = m.ParallelRank("lorem", in)
// 	elapsed := time.Since(start)
// 	t.Logf("Paralelle took %s\n", elapsed)
//
// 	start = time.Now()
// 	_ = m.RankMatches("lorem", in)
// 	elapsed = time.Since(start)
// 	t.Logf("Normal took %s\n", elapsed)
// }
