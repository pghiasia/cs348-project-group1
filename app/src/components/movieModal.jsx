import React, { useState, useEffect } from 'react';
import { Modal, Box, Typography, Chip, Stack } from '@mui/material';

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

  useEffect(() => {
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

    fetchMovieDetails();
  }, [open, movie]);

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
            <Typography variant="subtitle1">Rating: {movie.AverageRating}</Typography>
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