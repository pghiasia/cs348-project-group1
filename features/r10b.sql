WITH GivenActor AS (
    SELECT pID
    FROM people
    WHERE name = 'John Doe'
),
FirstConnections AS (
    SELECT DISTINCT w2.pID AS first_pID
    FROM workedOn w1 JOIN workedOn w2 ON w1.tID = w2.tID
    WHERE w1.pID = (SELECT pID FROM GivenActor) AND w2.pID != (SELECT pID FROM GivenActor)
),
SecondConnections AS (
    SELECT DISTINCT w3.pID AS second_pID
    FROM workedOn w2 JOIN workedOn w3 ON w2.tID = w3.tID
    WHERE w2.pID IN (SELECT first_pID FROM FirstConnections) AND w3.pID NOT IN (
            SELECT pID FROM GivenActor
            UNION
            SELECT first_pID FROM FirstConnections
        )
)
SELECT fc.first_pID AS actorID, 'first' AS ConnectionLevel, p.name
FROM ((FirstConnections fc JOIN people p ON fc.first_pID = p.pID)
    UNION
    (SELECT sc.second_pID AS actorID, 'second' AS ConnectionLevel, p.nameFROM SecondConnections sc))
    JOIN people p ON sc.second_pID = p.pID;


