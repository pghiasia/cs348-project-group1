SELECT DISTINCT a.tID, a.primaryTitle, a.releaseYear, a.averageRating, a.isAdult, a.titleType
FROM
((((SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes, 'movie' AS titleType
FROM movie) 
UNION
(SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes, 'series' AS titleType
FROM series) 
UNION
(SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes, 'short' AS titleType
FROM short)
UNION
(SELECT tID, primaryTitle, OriginalTitle, isAdult, releaseYear, averageRating, numVotes, runtimeMinutes, 'episode' AS titleType
FROM episodes))
NATURAL JOIN 
workedOn)
NATURAL JOIN 
people) AS a

WHERE a.name = 'Fred Astaire' AND a.averageRating is not null ORDER BY a.averageRating DESC LIMIT 20;