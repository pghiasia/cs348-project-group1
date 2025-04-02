WITH RECURSIVE 
temp AS (
    SELECT pID FROM people WHERE name = 'Fred Astaire' LIMIT 1 -- !!! PUT ACTOR NAME HERE !!!
),
ActorNetwork (anc, des, depth) AS (
    -- Base case: start with the given actor
    SELECT W1.pID AS anc, W2.pID AS des, 1 AS depth
    FROM workedOn AS W1 JOIN workedOn AS W2 ON (W1.tid = W2.tid AND W1.pID != W2.pID)
    WHERE W1.pID = (SELECT * FROM temp) 
    UNION
    -- Recursive case: find new connections via workedOn table
    SELECT AN.des AS anc, W2.pID AS des, depth + 1
    FROM ActorNetwork AN, workedOn W1, workedOn W2
    WHERE AN.des = W1.pID AND W1.tID = W2.tID AND W1.pID != W2.pID
    AND AN.depth < 2 -- !!! PUT NUMBER OF CONNECTIONS HERE !!!
) 

SELECT P1.name AS anc, P2.name AS des 
FROM (SELECT DISTINCT anc, des FROM ActorNetwork) AS AN
JOIN people P1 ON AN.anc = P1.pID 
JOIN people P2 ON AN.des = P2.pID;
