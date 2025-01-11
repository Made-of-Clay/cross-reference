// I normally chunk these out individually per model, but this might be a small enough project to do 1 file.

/*
[{
    title: 'Old Testament',
    books: [{ // Book[]
        title: 'Genesis',
        abbv: 'Gen',
        inOT: true, // might not use this
        refs: [{  // Reference[]
            source: 'Gen.1.1',
            target: 'Eccl.12.1',
            book: 'Gen'
        }],
    }],
}]
*/

export type Testament = string // 

export interface Book {
    abbv: string
    title: string
    inOT: boolean // Old Testament (OT) New Testament (NT)
    refs: Reference[]
}

export interface Reference {
    source: string
    target: string // may be single verse or range of verses
    book: Book
}
