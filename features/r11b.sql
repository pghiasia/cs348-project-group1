WITH RECURSIVE ActorNetwork (pID, depth, path) AS (
    -- Base case: start with the given actor
    SELECT p.pID, 0 AS depth, CAST(p.pID AS CHAR(200)) AS path, CAST(p.name AS CHAR(200)) AS name_path
    FROM people p
    WHERE p.name = 'Fred Astaire' -- !!! PUT ACTOR NAME HERE !!!
    UNION ALL
    -- Recursive case: find new connections via workedOn table
    SELECT w2.pID, an.depth + 1, CONCAT(an.path, '->', w2.pID), CONCAT(an.name_path, '->', (SELECT name FROM people WHERE pid = w2.pID))
    FROM ActorNetwork an
        JOIN workedOn w1 ON an.pID = w1.pID
        JOIN workedOn w2 ON w1.tID = w2.tID
    -- Prevent cycles by ensuring the actor is not already in the path
    WHERE an.path NOT LIKE '%' || CAST(w2.pID AS VARCHAR) || '%' -- manually check if it has appeared in path or not, cus there are cycles and regular equality doesn't work
        AND an.depth < 4 -- !!! PUT NUMBER OF CONNECTIONS HERE !!!
) 
SELECT DISTINCT p.name AS last_connection_name, an.pID AS last_connection_pID, an.depth, an.path, an.name_path 
FROM ActorNetwork an 
    JOIN people p ON an.pID = p.pID 
ORDER BY an.depth;
