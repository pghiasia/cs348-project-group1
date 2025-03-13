import React, { useState, useEffect } from 'react';
import MovieCard from '../components/movieCard';
import MovieModal from '../components/movieModal';
import { Container, Grid, Box, Button } from '@mui/material';

const Movie = () => {
  const [movies, setMovies] = useState([]);
  const [selectedMovie, setSelectedMovie] = useState(null);
  const [open, setOpen] = useState(false);

  const handleFecthMovies = async () => {
    try{
        const response = await fetch('http://localhost:9888/movies', {
            method: "POST",
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              "sortby": "title"
            })
        })
    
        if (!response.ok) {
            throw new Error('Fetch movie failed');
        }
    
        const data = await response.json();
        setMovies(data)
    }
    catch (error) {
        console.error(error)
    }
}

  useEffect(() => {
    handleFecthMovies()
  }, []);

  const handleOpen = (movie) => {
    setSelectedMovie(movie);
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
    setSelectedMovie(null);
  };

  const sortByRating = async () => {
    try {
      const response = await fetch('http://localhost:9888/movies', {
          method: "POST",
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            "sortby": "rating"
          })
      })
  
      if (!response.ok) {
          throw new Error('Fetch movie failed');
      }
  
      const data = await response.json();
      setMovies(data)
    }
    catch (error) {
      console.error(error)
    }
  };

  const sortByName = async () => {
    try {
      const response = await fetch('http://localhost:9888/movies', {
          method: "POST",
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            "sortby": "title"
          })
      })
  
      if (!response.ok) {
          throw new Error('Fetch movie failed');
      }
  
      const data = await response.json();
      setMovies(data)
    }
    catch (error) {
      console.error(error)
    }
  };

  const sortByGenre = async () => {
    try {
      const response = await fetch('http://localhost:9888/movies', {
          method: "POST",
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            "sortby": "genres"
          })
      })
  
      if (!response.ok) {
          throw new Error('Fetch movie failed');
      }
  
      const data = await response.json();
      setMovies(data)
    }
    catch (error) {
      console.error(error)
    }
  };

  return (
    <Container>
      <Box display="flex" justifyContent="center" mb={4} mt={4}>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={sortByRating}
          sx={{ margin: '0 10px', transition: 'background-color 0.3s', '&:hover': { backgroundColor: '#1e88e5' } }}
        >
          Sort by Rating
        </Button>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={sortByName}
          sx={{ margin: '0 10px', transition: 'background-color 0.3s', '&:hover': { backgroundColor: '#1e88e5' } }}
        >
          Sort by Name
        </Button>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={sortByGenre}
          sx={{ margin: '0 10px', transition: 'background-color 0.3s', '&:hover': { backgroundColor: '#1e88e5' } }}
        >
          Sort by Genres
        </Button>
      </Box>
      <Grid container spacing={2}>
        {movies.map(movie => (
          <Grid item xs={12} sm={6} md={4} key={movie.Mid}>
            <MovieCard movie={movie} onClick={handleOpen} />
          </Grid>
        ))}
      </Grid>
      {selectedMovie && (
        <MovieModal open={open} handleClose={handleClose} movie={selectedMovie} />
      )}
    </Container>
  );
}

export default Movie;