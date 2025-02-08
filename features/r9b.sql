SELECT *
FROM movies m NATURAL JOIN movie_to_actor ma NATURAL JOIN actors a
WHERE a.name = "Fred Astaire"
ORDER BY m.rating DESC
LIMIT 1;
