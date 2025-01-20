import data from '../crossReferences.json';

export interface Reference {
    source: string
    target: string
}

export function getReferences(): Reference[] {
    return (data as { references: Reference[] }).references;
}