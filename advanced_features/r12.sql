-- 8 is the rating being added, for movie tt0000001

--Display rating prior to update
BEGIN TRANSACTION;

SELECT * FROM 
((SELECT tID, primaryTitle, averageRating, numVotes
FROM movie) 
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
UPDATE movie
SET numVotes = numVotes + 1, 
    averageRating = ((averageRating * (numVotes - 1)) + 8) / numVotes
WHERE tID =  'tt0000001';

UPDATE short
SET numVotes = numVotes + 1, 
    averageRating = ((averageRating * (numVotes - 1)) + 8) / numVotes
WHERE tID =  'tt0000001';

UPDATE series
SET numVotes = numVotes + 1, 
    averageRating = ((averageRating * (numVotes - 1)) + 8) / numVotes
WHERE tID =  'tt0000001';

UPDATE episodes
SET numVotes = numVotes + 1, 
    averageRating = ((averageRating * (numVotes - 1)) + 8) / numVotes
WHERE tID =  'tt0000001';


--Display rating after update
SELECT * FROM 
((SELECT tID, primaryTitle, averageRating, numVotes
FROM movie) 
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

COMMIT;
