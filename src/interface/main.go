package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

const debug = false

type StringSlice []string

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	var buf io.Writer

	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		// ...
		fmt.Println(buf)
	}

	strs := StringSlice{"a", "v", "ba", "e", "g", "q", "c"}
	fmt.Println(strs)
	sort.Sort(StringSlice(strs))
	fmt.Println(strs)

	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Lenght")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
