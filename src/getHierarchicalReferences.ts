import { getReferences, Reference } from './getReferences';
import { ReferencesByBook } from './main';

export function getHierarchicalReferences() {
  console.log('getting references');
  const data = getReferences().slice(13000, 14000);
  console.assert(!!data.length, 'Data has no entries');
  const groupedRefs: ReferencesByBook<Reference>[] = [];
  console.time('Sort References');
  for (const ref of data) {
    const book = ref.source.split('.')[0];
    let found = groupedRefs.find(y => y.name === book);
    if (!found) {
      groupedRefs.push({ name: book, children: [] });
    }
    const curBook = found ?? groupedRefs[groupedRefs.length - 1];
    curBook.children.push(ref);
  }
  console.timeEnd('Sort References');
  console.log('sorted references by book', groupedRefs);
  return groupedRefs;
}
