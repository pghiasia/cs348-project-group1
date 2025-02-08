SELECT title, rating 
FROM movies
WHERE rating = (SELECT MAX(rating) FROM movies);
