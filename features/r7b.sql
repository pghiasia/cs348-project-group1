SELECT *

(((((SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes
FROM movies) 
UNION
(SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes
FROM series) 
UNION
(SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes
FROM short))

NATURAL JOIN 

genres)

NATURAL JOIN 

workedOn)

NATURAL JOIN 

people) AS a

WHERE a.name = 'Brigitte Bardot' AND a.genres LIKE '%Romance%';
