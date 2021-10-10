db = new Mongo().getDB("url_shortener");
db.createCollection('links', { capped: false });
db.links.createIndex( { "id": 1 }, { unique: true } )
db.links.insert({"id": NumberLong(14776336), "link": "https://example.com"})