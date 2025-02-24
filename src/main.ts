import './style.css'
import * as d3 from 'd3';
import { getHierarchicalReferences } from './getHierarchicalReferences';

const width = 900;
const radius = width / 2;

export interface ReferencesByBook<T> {
  name: string // book
  children: T[] // references
}

const tree = d3.cluster()
  .size([2 * Math.PI, radius - 100]);

const hierRefs = getHierarchicalReferences();
const root = tree(d3.hierarchy<ReferencesByBook<Reference>>(hierRefs))

const line = d3.lineRadial()
  .curve(d3.curveBundle.beta(0.85))
  .radius(([, y]) => y)
  .angle(([x]) => x / 180 * Math.PI);

const svg = d3.create('svg')
  .attr('width', width)
  .attr('height', width)
  .attr('viewBox', [-width / 2, -width / 2, width, width])
  .attr('style', 'max-width: 100%; height: auto; font: 10px sans-serif;');

const link = svg.append("g").selectAll(".link");
const node = svg.append("g").selectAll(".node");

document.body.append(svg.node() ?? '');

// console.log(data)

// const root = tree(d3.hierarchy<Reference[]>(data))
// console.log(data.leaves());

// function getReferences() {
//   import('../crossReferences.json').then(data => data.default?.references ?? []).then(console.log)
// }

// https://observablehq.com/@d3/hierarchical-edge-bundling
// https://d3js.org/getting-started#installing-from-npm

// maybe group by testament
// then group by book
// verses are leaves

// maybe sort by bookIndex (e.g. Genesis == 0, Exodus == 1, etc.)

// starting data should only be books - { title: string; inx: number; }
// this will get me a circle with names rotating around the outside