package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"

	"git.sr.ht/~mendelmaleh/notally"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	// dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	var doc notally.ExportedNotes
	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	// pretty.Print(doc)

	var b strings.Builder

	b.WriteString("> notes\n\n")
	notes(&b, doc.Notes)

	b.WriteString("> archived notes\n\n")
	notes(&b, doc.ArchivedNotes)

	b.WriteString("> deleted notes\n\n")
	notes(&b, doc.DeletedNotes)

	// replace left angle brackets because they break formatting
	out := strings.ReplaceAll(b.String(), "<", "`<`")

	fmt.Print(out)
}

func notes(b *strings.Builder, n notally.Notes) {
	for _, v := range n.Note {
		note(b, v)
		b.WriteString("\n")
	}

	for _, v := range n.List {
		list(b, v)
		b.WriteString("\n")
	}
}

func common(b *strings.Builder, c notally.Common) {
	b.WriteString("# ")

	if c.Title == "" {
		c.Title = "(unnamed)"
	}

	b.WriteString(c.Title)
	b.WriteString(c.DateCreated.Format(" (2006-01-02)\n"))
}

func note(b *strings.Builder, n notally.Note) {
	common(b, n.Common)

	b.WriteString(n.Body)
	b.WriteString("\n")
}

func list(b *strings.Builder, n notally.List) {
	common(b, n.Common)

	for _, i := range n.Item {
		if i.Checked == "true" {
			b.WriteString("+ ")
		} else {
			b.WriteString("- ")
		}

		b.WriteString(i.Text)
		b.WriteString("\n")
	}
}
