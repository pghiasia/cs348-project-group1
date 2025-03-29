import React, { useState, useEffect } from 'react';
import { Modal, Box, Typography, Chip, Stack, TextField, Button } from '@mui/material';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  maxHeight: '80vh',
  overflowY: 'auto',
  bgcolor: 'background.paper',
  boxShadow: 24,
  p: 4,
};

const MovieModal = ({ open, handleClose, movie }) => {
  const [movieDetails, setMovieDetails] = useState({
    Genres: [],
    Crew: []
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [userRating, setUserRating] = useState('');
  const [ratingSubmitted, setRatingSubmitted] = useState(false);

  const fetchMovieDetails = async () => {
    if (!open || !movie) return;

    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch(`http://localhost:9888/movie?tid=${movie.Tid}`, {
        method: "GET",
        headers: {
          'Content-Type': 'application/json',
        }
      });

      if (!response.ok) {
        throw new Error('Failed to fetch movie details');
      }

      const data = await response.json();
      setMovieDetails(data);
    } catch (err) {
      console.error('Error fetching movie details:', err);
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchMovieDetails();
    setUserRating('');
    setRatingSubmitted(false);
  }, [open, movie]);

  const handleRatingChange = (e) => {
    const value = e.target.value;
    if (value === '' || (parseFloat(value) >= 0 && parseFloat(value) <= 10)) {
      setUserRating(value);
    }
  };

  const handleRatingSubmit = async () => {
    if (userRating === '' || !movie) return;
    
    setIsLoading(true);
    setError(null);
    
    try {
      const response = await fetch(`http://localhost:9888/rating`, {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          tid: movie.Tid,
          rating: parseFloat(userRating)
        })
      });

      if (!response.ok) {
        throw new Error('Failed to submit rating');
      }

      setRatingSubmitted(true);
      
      await fetchMovieDetails();
    } catch (err) {
      console.error('Error submitting rating:', err);
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const uniqueCrew = Array.from(new Set(movieDetails.Crew));

  return (
    <Modal open={open} onClose={handleClose}>
      <Box sx={style}>
        <Typography variant="h6" gutterBottom>
          {movie.PrimaryTitle}
        </Typography>

        {isLoading ? (
          <Typography>Loading details...</Typography>
        ) : error ? (
          <Typography color="error">Error: {error}</Typography>
        ) : (
          <>
            <Typography variant="subtitle1">Title Type: {movie.TitleType}</Typography>
            <Typography variant="subtitle1">Release: {movie.ReleaseYear}</Typography>
            
            <Box sx={{ display: 'flex', alignItems: 'center', mt: 1, mb: 1 }}>
              <Typography variant="subtitle1" sx={{ mr: 2 }}>
                Rating: {Number(movie.AverageRating).toFixed(2)}
              </Typography>
              
              <TextField
                size="small"
                label="Your Rating (0-10)"
                variant="outlined"
                value={userRating}
                onChange={handleRatingChange}
                type="number"
                InputProps={{
                  inputProps: { 
                    min: 0, 
                    max: 10,
                    step: 0.1
                  },
                  sx: { width: '120px' }
                }}
              />
              
              <Button 
                variant="contained" 
                color="primary"
                size="small"
                onClick={handleRatingSubmit}
                sx={{ ml: 1 }}
                disabled={userRating === ''}
              >
                Rate
              </Button>
            </Box>
            
            {ratingSubmitted && (
              <Typography variant="body2" color="success.main" sx={{ mb: 1 }}>
                Rating submitted successfully!
              </Typography>
            )}
            
            <Typography variant="subtitle1">Adult: {movie.IsAdult.toString()}</Typography>
            <Typography variant="subtitle1">Movie ID: {movie.Tid}</Typography>

            {movieDetails.Genres.length > 0 && (
              <Box mt={2}>
                <Typography variant="subtitle1" gutterBottom>
                  Genres:
                </Typography>
                <Stack direction="row" spacing={1} useFlexGap flexWrap="wrap">
                  {movieDetails.Genres.map((genre, index) => (
                    <Chip 
                      key={index} 
                      label={genre} 
                      variant="outlined" 
                      color="primary" 
                    />
                  ))}
                </Stack>
              </Box>
            )}

            {uniqueCrew.length > 0 && (
              <Box mt={2}>
                <Typography variant="subtitle1" gutterBottom>
                  Crew Members:
                </Typography>
                <Stack direction="row" spacing={1} useFlexGap flexWrap="wrap">
                  {uniqueCrew.map((crewMember, index) => (
                    <Chip 
                      key={index} 
                      label={crewMember} 
                      variant="outlined" 
                      color="secondary" 
                    />
                  ))}
                </Stack>
              </Box>
            )}
          </>
        )}
      </Box>
    </Modal>
  );
};

export default MovieModal;