// from local
//LOAD CSV FROM 'file:///home/iox/dev/github.com/seblkma/justgo/virtue/neo4j/csv-import/products.csv' AS row
//RETURN count(row)

// my github csv
//https://raw.githubusercontent.com/SebLKMa/justgo/master/virtue/neo4j/csv-import/products.csv

//LOAD CSV FROM 'https://raw.githubusercontent.com/SebLKMa/justgo/master/virtue/neo4j/csv-import/products.csv' AS row
//RETURN row

//LOAD CSV FROM 'https://raw.githubusercontent.com/SebLKMa/justgo/master/virtue/neo4j/csv-import/products.csv' AS row
//WITH toInteger(row[0]) AS productId, row[1] AS productName, toFloat(row[2]) AS unitCost
//MERGE (p:Product {productId: productId})
//  SET p.productName = productName, p.unitCost = unitCost
//RETURN count(p)

// validate products loaded correctly
//MATCH (p:Product)
//RETURN p LIMIT 20

// import persons
//LOAD CSV FROM 'https://raw.githubusercontent.com/SebLKMa/justgo/master/virtue/neo4j/csv-import/persons.csv' AS row
//WITH row[0] AS name, row[1] AS email
//MERGE (p:Person {name: name})
//  SET p.name = name, p.email = email
//RETURN count(p)

// validate loaded correctly
//MATCH (p:Person)
//RETURN p LIMIT 20

// import emails
//LOAD CSV FROM 'https://raw.githubusercontent.com/SebLKMa/justgo/master/virtue/neo4j/csv-import/emails.csv' AS row
//WITH row[0] AS email, row[1] AS name
//MERGE (e:Emails {email: email})
//  SET e.email = email, e.name = name
//RETURN count(e)

//CREATE INDEX ON :Person(email);

// validate loaded correctly
//MATCH (e:Emails)
//RETURN e LIMIT 20

// working example from http://www.makedatauseful.com/graph-relations-in-neo4j-simple-load-example/

// import nodes from csv
LOAD CSV WITH HEADERS FROM "https://raw.githubusercontent.com/SebLKMa/justgo/master/virtue/neo4j/csv-import/persons.csv" AS nodes 
CREATE (p:Person { email: nodes.email });

// show nodes
MATCH (p:Person)
RETURN p LIMIT 20;

// import edges from csv
LOAD CSV WITH HEADERS FROM "https://raw.githubusercontent.com/SebLKMa/justgo/master/virtue/neo4j/csv-import/friendship.csv" AS edges
MATCH (a:Person { email: edges.source})
MATCH (b:Person { email: edges.target })
CREATE (a)-[:FRIEND_OF]->(b);

// show nodes edges
MATCH (p1:Person)-[rel:FRIEND_OF]->(p2:Person)
RETURN p1, rel, p2 LIMIT 50;

// delete all
//MATCH (n) DELETE n