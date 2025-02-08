SELECT *
FROM movies m NATURAL JOIN movie_to_actor ma NATURAL JOIN actors a
WHERE a.name = 'Brigitte Bardot' AND m.genres LIKE '%Romance%';
