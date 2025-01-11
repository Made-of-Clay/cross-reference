import './style.css'
import * as d3 from 'd3';

const width = 900;
const radius = width / 2;

const svg = d3.create('svg')
  .attr('width', width)
  .attr('height', width)
  .attr('viewBox', [-width / 2, -width / 2, width, width])
  .attr('style', 'max-width: 100%; height: auto; font: 10px sans-serif;');

document.body.append(svg.node() ?? '');

// https://observablehq.com/@d3/hierarchical-edge-bundling
// https://d3js.org/getting-started#installing-from-npm

// maybe group by testament
// then group by book
// verses are leaves

// maybe sort by bookIndex (e.g. Genesis == 0, Exodus == 1, etc.)

// starting data should only be books - { title: string; inx: number; }
// this will get me a circle with names rotating around the outside