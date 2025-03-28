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
    -- Only movies are desired.
    a.titleType = 'movie'
    -- Adult flag set to non-adult.
    AND a.isAdult = 0
    -- Title keyword filter (searching in original title).
    AND a.originalTitle LIKE '%Miss%'
    -- Release year range filter.
    AND a.releaseYear BETWEEN 1800 AND 2020
    -- Average rating range filter.
    AND a.averageRating BETWEEN 1.0 AND 9.0
    -- Genre filter: ensure title exists in relatedGenres with a genre containing "Comedy".
    AND EXISTS (
        SELECT 1
        FROM genres rg
        WHERE rg.tid = a.tID
        AND rg.genre = 'Romance'
    )
    -- Crew member filter: ensure a crew member with a name containing "Tom Hanks" worked on this title.
    AND EXISTS (
        SELECT 1
        FROM workedOn w
        JOIN people p ON w.pID = p.pID
        WHERE w.tID = a.tID
        AND p.name = 'Marlon Brando'
    );
