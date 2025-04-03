SELECT title, rating 
FROM movie
WHERE rating = (SELECT MAX(rating) FROM movie);
