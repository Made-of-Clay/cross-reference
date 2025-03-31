export function exportSvg() {
    const svgXml = '<?xml version="1.0" encoding="UTF-8"?>';
    const svg = document.body.querySelector('svg');
    const doc = document.createDocumentFragment();
    doc.append(svg.cloneNode(true));
    doc.querySelector('svg').prepend(document.head.querySelector('style').cloneNode(true));
  
    const svgStr = new XMLSerializer().serializeToString(doc);
    const blob = new Blob([svgXml + svgStr], { type: "image/svg+xml" });
    const url = URL.createObjectURL(blob);
  
    const link = document.createElement('a');
    link.href = url;
    link.download = 'graphic.svg';
    link.click();
}
