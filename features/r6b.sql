WITH AllTitles AS (
		SELECT tID, isAdult, releaseYear, originalTitle, averageRating,
		numVotes, runtimeMinutes, primaryTitle, 'movie' AS titleType
		FROM movie
		UNION ALL
		SELECT tID, isAdult, releaseYear, originalTitle, averageRating,
		numVotes, runtimeMinutes, primaryTitle, 'series' AS titleType
		FROM series
		UNION ALL
		SELECT tID, isAdult, releaseYear, originalTitle, averageRating,
		numVotes, runtimeMinutes, primaryTitle, 'short' AS titleType
		FROM short
		UNION ALL
		SELECT tID, isAdult, releaseYear, originalTitle, averageRating,
		numVotes, runtimeMinutes, primaryTitle, 'episode' AS titleType
		FROM episodes
	)
SELECT DISTINCT a.tID, a.primaryTitle, a.releaseYear, a.averageRating, a.isAdult, a.titleType
FROM AllTitles a
WHERE
    tID is not null AND averageRating is not null
    AND EXISTS (FROM genres rg SELECT 1 WHERE rg.tid = a.tID AND rg.genre = 'Comedy')
    AND a.titleType = 'movie'
    AND EXISTS (SELECT 1 FROM workedOn w JOIN people p ON w.pID = p.pID WHERE w.tID = a.tID AND p.name = 'Marlon Brando')
    AND a.originalTitle LIKE '%a%'
    AND a.releaseYear BETWEEN 1000 AND 2025
    AND a.averageRating BETWEEN 1 AND 9
    AND a.isAdult = 0
LIMIT 1000;