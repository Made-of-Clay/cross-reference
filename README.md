# Cross Reference

This repo is meant to provide data visualizations of biblical cross references between books. This could take many shapes.

The inspiration for this comes from work done by [Chris Harrison](https://www.chrisharrison.net/index.php/visualizations/BibleViz) as referenced by [Jordan Peterson](https://youtu.be/f-wWBGo6a2w?si=ANEP2CA4G2Zac8s5&t=4362) in his Bible Lecture Series. Peterson asserts the Bible is the world's first hyperlinked text.

While I love Harrison's data viz, I want to create my own with the references using a circular layout rather than the arc graph Harrison generated. [d3's hierarchical edge bundling](https://observablehq.com/@d3/hierarchical-edge-bundling?intent=fork) appears to be the best use case for this.

## Paused Progress

After working on this through Jan 2025, I've found my skills or comprehension lacking. (#skillissues) I need to learn more about the graph concept and/or the d3 tools available to render such a graph. In addition, the data I'm using is very large. I knew this would be the case, and I was originally aiming to make a static render SVG (or export as PDF); no interactivity. There would be 200k+ SVG nodes alone for the verses and linking reference arcs.

I've considered more stripped-down UI options for better interactivity. One version would list the available verses in a sidebar that, when hovered, would render the linking arcs (incoming and outgoing). Those reference arcs would be dynamically added to the DOM, otherwise, the problem of SVG nodes wouldn't change. That's still 100k+ in verses alone.

I decided to pause this project, because I'm only working on this as a side (passion) project. I've given a month to work through some of the issues and do research. While that's resulted in the current state of the repo, it's also shown me there's more work to do. Rather than grind it out right now, I'm pivoting to the next project until I return and take another crack at it.

More to come...
