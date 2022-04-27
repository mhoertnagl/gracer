package obj

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/render"
)

type Parser struct {
	Vertices []*vertexContainer
	Root     *render.Group
}

type vertexContainer struct {
	Vertex alg.Vector
}

func NewParser() *Parser {
	return &Parser{
		Vertices: make([]*vertexContainer, 0),
		Root:     render.NewGroup(),
	}
}

func (p *Parser) ParseFile(file string) {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()
	p.ParseReader(f)
}

func (p *Parser) ParseString(text string) {
	p.ParseReader(strings.NewReader(text))
}

func (p *Parser) ParseReader(r io.Reader) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) > 0 {
			p.parseLine(line)
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}

func (p *Parser) parseLine(line string) {
	switch line[0:1] {
	case "v":
		p.parseVertex(line)
	case "f":
		p.parseFace(line)
	}
}

func (p *Parser) parseVertex(line string) {
	var v1 float64
	var v2 float64
	var v3 float64
	fmt.Sscanf(line, "v %f %f %f", &v1, &v2, &v3)
	vertex := alg.NewPoint(v1, v2, v3)
	container := &vertexContainer{Vertex: vertex}
	p.Vertices = append(p.Vertices, container)
}

func (p *Parser) parseFace(line string) {
	var f1 int
	var f2 int
	var f3 int
	fmt.Sscanf(line, "f %d %d %d", &f1, &f2, &f3)
	v1 := p.getVertex(f1)
	v2 := p.getVertex(f2)
	v3 := p.getVertex(f3)
	tr := render.NewTriangle(v1, v2, v3)
	p.Root.AddKid(tr)
}

// Index is 1-based!
func (p *Parser) getVertex(idx int) alg.Vector {
	return p.Vertices[idx-1].Vertex
}

func (p *Parser) getTriangle(idx int) *render.Triangle {
	return p.Root.Kids[idx].(*render.Triangle)
}
