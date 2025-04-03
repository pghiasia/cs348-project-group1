SELECT genre FROM genres WHERE tID = 'tt0000001';

SELECT name FROM workedOn w JOIN people p ON w.pID = p.pID WHERE w.tID = 'tt0000001';