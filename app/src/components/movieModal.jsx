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
      <Typography variant="h6">{movie.PrimaryTitle}</Typography>
      <Typography>Title Type: {movie.TitleType}</Typography>
      <Typography>Release: {movie.ReleaseYear}</Typography>
      <Typography>Rating: {movie.AverageRating}</Typography>
      <Typography>Adult: {movie.IsAdult.toString()}</Typography>
      <Typography>Movie ID: {movie.Tid}</Typography>
    </Box>
  </Modal>
);

export default MovieModal;