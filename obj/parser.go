package obj

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/render"
)

type Parser struct {
	Root     *render.Group
	group    *render.Group
	vertices []alg.Vector
	normals  []alg.Vector
}

func NewParser() *Parser {
	group := render.NewGroup()
	return &Parser{
		Root:     group,
		group:    group,
		vertices: make([]alg.Vector, 0),
		normals:  make([]alg.Vector, 0),
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
	group := render.NewGroup()
	p.Root = group
	p.group = group
	p.vertices = make([]alg.Vector, 0)
	p.normals = make([]alg.Vector, 0)

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
	if len(line) > 2 {
		switch line[0:2] {
		case "v ":
			p.parseVertex(line)
		case "vn":
			p.parseVertexNormal(line)
		case "f ":
			p.parseFace(line)
		case "g ":
			p.parseGroup(line)
		}
	}
}

func (p *Parser) parseVertex(line string) {
	var v1 float64
	var v2 float64
	var v3 float64
	fmt.Sscanf(line, "v %f %f %f", &v1, &v2, &v3)
	vertex := alg.NewPoint(v1, v2, v3)
	p.vertices = append(p.vertices, vertex)
}

func (p *Parser) parseVertexNormal(line string) {
	var v1 float64
	var v2 float64
	var v3 float64
	fmt.Sscanf(line, "vn %f %f %f", &v1, &v2, &v3)
	normal := alg.NewVector3(v1, v2, v3)
	p.normals = append(p.normals, normal)
}

func (p *Parser) parseFace(line string) {
	// vs := parseIntList(line[2:])
	vs := parseFaceParams(line[2:])
	v1 := p.getVertex(vs[0].vertex)
	for i := 1; i < len(vs)-1; i++ {
		v2 := p.getVertex(vs[i].vertex)
		v3 := p.getVertex(vs[i+1].vertex)
		if vs[i].smooth {
			n1 := p.getVertexNormal(vs[0].normal)
			n2 := p.getVertexNormal(vs[i].normal)
			n3 := p.getVertexNormal(vs[i+1].normal)
			tr := render.NewSmoothTriangle(v1, v2, v3, n1, n2, n3)
			p.group.AddKid(tr)
		} else {
			tr := render.NewTriangle(v1, v2, v3)
			p.group.AddKid(tr)
		}
	}
}

// func parseIntList(list string) []int {
// 	parts := strings.Split(list, " ")
// 	ints := make([]int, len(parts))
// 	for i := 0; i < len(ints); i++ {
// 		if v, err := strconv.Atoi(parts[i]); err == nil {
// 			ints[i] = v
// 		} else {
// 			panic(err)
// 		}
// 	}
// 	return ints
// }

func parseFaceParams(line string) []*faceParam {
	parts := strings.Split(strings.TrimSpace(line), " ")
	params := make([]*faceParam, len(parts))
	for i := 0; i < len(parts); i++ {
		param := &faceParam{smooth: false, vertex: 0, normal: 0}
		subparts := strings.Split(parts[i], "/")
		switch len(subparts) {
		case 1:
			param.vertex = parseIntOrFail(subparts[0])
		case 2:
			param.smooth = true
			param.vertex = parseIntOrFail(subparts[0])
			param.normal = parseIntOrFail(subparts[1])
		case 3:
			param.smooth = true
			param.vertex = parseIntOrFail(subparts[0])
			param.normal = parseIntOrFail(subparts[2])
		}
		params[i] = param
	}
	return params
}

func parseIntOrFail(value string) int {
	if v, err := strconv.Atoi(value); err == nil {
		return v
	} else {
		panic(err)
	}
}

type faceParam struct {
	smooth bool
	vertex int
	normal int
}

func (p *Parser) parseGroup(line string) {
	p.group = render.NewGroup()
	p.Root.AddKid(p.group)
}

// Index is 1-based!
func (p *Parser) getVertex(idx int) alg.Vector {
	return p.vertices[idx-1]
}

// Index is 1-based!
func (p *Parser) getVertexNormal(idx int) alg.Vector {
	return p.normals[idx-1]
}

func (p *Parser) getTriangle(idx int) *render.Triangle {
	return p.Root.Kids[idx].(*render.Triangle)
}

func (p *Parser) getSmoothTriangle(idx int) *render.SmoothTriangle {
	return p.Root.Kids[idx].(*render.SmoothTriangle)
}
