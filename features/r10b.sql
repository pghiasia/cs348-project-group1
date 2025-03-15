WITH GivenActor AS (
    SELECT pID
    FROM people
    WHERE name = 'James Cagney'
),
FirstConnections AS (
    SELECT DISTINCT w2.pID
    FROM workedOn w1 JOIN workedOn w2 ON w1.tID = w2.tID
    WHERE w1.pID IN (SELECT pID FROM GivenActor) AND w2.pID NOT IN (SELECT pID FROM GivenActor)
)
SELECT name FROM FirstConnections NATURAL JOIN people;
