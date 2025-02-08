-- 8 is the rating being added, for movie tt0000001
SELECT * FROM movies WHERE mid = 'tt0000001';

UPDATE movies
SET numVotes = (SELECT numVotes + 1 FROM movies WHERE mid = 'tt0000001')
WHERE mid =  'tt0000001';

UPDATE movies
SET rating = (SELECT  (rating * (numVotes - 1) + 8) / numVotes  FROM movies WHERE mid = 'tt0000001')
WHERE mid =  'tt0000001';

SELECT * FROM movies WHERE mid = 'tt0000001';
