# Project Notes

## D3 Example

Hierarchical edge bundling in D3 is illustrated by showing some library's import/export relationships between the classes. Each text element around the circumference is a leaf of the tree. Each line between each node/leaf is an import statement.

The text is like a book of the Bible (for the inter-book references) or each verse (for the inter-verse references).

That sets up a similar structure to the class imports: leaf/node, inbound/outbound references.

I need to format the data similarly.

Before trying to format the data, I should start with trivially simple data and learn better what the algo is doing to render the graphic. After that, I can more reliably scale up.

Node
    Input
    Output
    Text
    Id?

### Dummy data examples

| Source   | Target    |
| -------- | --------- |
| Gen.22.2 | John.3.16 |
| Ps.36.7  | John.3.16 |
| Ps.89.11 | Gen.1.1   |
| Dan.6.10 | Rev.2.10  |
