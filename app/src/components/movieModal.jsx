import React from 'react';
import { Modal, Box, Typography } from '@mui/material';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  boxShadow: 24,
  p: 4,
};

const MovieModal = ({ open, handleClose, movie }) => (
  <Modal open={open} onClose={handleClose}>
    <Box sx={style}>
      <Typography variant="h6">{movie.Title}</Typography>
      <Typography>Genres: {movie.Genres}</Typography>
      <Typography>Release: {movie.Release}</Typography>
      <Typography>Rating: {movie.Rating}</Typography>
      <Typography>NumVotes: {movie.NumVotes}</Typography>
    </Box>
  </Modal>
);

export default MovieModal;