SELECT *
FROM
(((((SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes
FROM movie) 
UNION
(SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes
FROM series) 
UNION
(SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes
FROM short)
UNION
(SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes
FROM episodes))
NATURAL JOIN 
genres)
NATURAL JOIN 
workedOn)
NATURAL JOIN 
people) AS a

WHERE a.name = 'Fred Astaire' AND a.genre LIKE '%Romance%';
