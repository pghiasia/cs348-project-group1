-- 8 is the rating being added, for movie tt0000001

--Display rating prior to update
SELECT * FROM 
((SELECT tID, primaryTitle, averageRating, numVotes
FROM movies) 
UNION
(SELECT tID, primaryTitle, averageRating, numVotes
FROM series) 
UNION
(SELECT tID, primaryTitle, averageRating, numVotes
FROM short)
UNION
(SELECT tID, primaryTitle, averageRating, numVotes
FROM episodes)) AS a
WHERE tID = 'tt0000001';

-- Update rating
UPDATE movies
SET numVotes = COALESCE((SELECT numVotes + 1 FROM movies WHERE tID = 'tt0000001'), numVotes), 
    averageRating = COALESCE((SELECT  (averageRating * (numVotes - 1) + 8) / numVotes  FROM movies WHERE tID = 'tt0000001'), averageRating)
WHERE tID =  'tt0000001';

UPDATE short
SET numVotes = COALESCE((SELECT numVotes + 1 FROM movies WHERE tID = 'tt0000001'), numVotes), 
    averageRating = COALESCE((SELECT  (averageRating * (numVotes - 1) + 8) / numVotes  FROM movies WHERE tID = 'tt0000001'), averageRating)
WHERE tID =  'tt0000001';

UPDATE series
SET numVotes = COALESCE((SELECT numVotes + 1 FROM movies WHERE tID = 'tt0000001'), numVotes), 
    averageRating = COALESCE((SELECT  (averageRating * (numVotes - 1) + 8) / numVotes  FROM movies WHERE tID = 'tt0000001'), averageRating)
WHERE tID =  'tt0000001';

UPDATE episodes
SET numVotes = COALESCE((SELECT numVotes + 1 FROM movies WHERE tID = 'tt0000001'), numVotes), 
    averageRating = COALESCE((SELECT  (averageRating * (numVotes - 1) + 8) / numVotes  FROM movies WHERE tID = 'tt0000001'), averageRating)
WHERE tID =  'tt0000001';


--Display rating after update
SELECT * FROM 
((SELECT tID, primaryTitle, averageRating, numVotes
FROM movies) 
UNION
(SELECT tID, primaryTitle, averageRating, numVotes
FROM series) 
UNION
(SELECT tID, primaryTitle, averageRating, numVotes
FROM short)
UNION
(SELECT tID, primaryTitle, averageRating, numVotes
FROM episodes)) AS a
WHERE tID = 'tt0000001';
