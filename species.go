/*


species.go implementation of species of genomes.

@licstart   The following is the entire license notice for
the Go code in this page.

Copyright (C) 2016 jin yeom, whitewolf.studio

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

As additional permission under GNU GPL version 3 section 7, you
may distribute non-source (e.g., minimized or compacted) forms of
that code without the copy of the GNU GPL normally required by
section 4, provided you include this license notice and a URL
through which recipients can access the Corresponding Source.

@licend    The above is the entire license notice
for the Go code in this page.


*/

package neat

// Species is an implementation of species of genomes in NEAT, which
// are separated by measuring compatibility distance among genomes
// within a population.
type Species struct {
	sid            int       // species ID
	numGenerations int       // species age (number of generations)
	representative *Genome   // species representative
	genomes        []*Genome // genomes in this species
}

// NewSpecies creates a new species given a species ID, and the genome
// that first populates the new species.
func NewSpecies(sid int, g *Genome) *Species {
	return &Species{
		sid:            sid,
		age:            0,
		representative: g,
		genomes:        []*Genome{g},
	}
}

// SID returns this species' species ID.
func (s *Species) SID() int {
	return s.sid
}

// Age returns this species' age.
func (s *Species) NumGenerations() int {
	return s.numGenerations
}

// Representative returns this species' representative.
func (s *Species) Representative() *Genome {
	return s.representative
}

// Genomes returns this species' member genomes.
func (s *Species) Genomes() []*Genome {
	return s.genomes
}

// AddGenome adds a new genome to this species.
func (s *Species) AddGenome(g *Genome) {
	s.genomes = append(s.genomes, g)
}
