import React from 'react';
import { Card, CardContent, Typography, Chip, Box } from '@mui/material';
import { getTitleTypeColor, getRatingColor } from './tagColors'

const MovieCard = ({ movie, onClick }) => {
  const handleMouseMove = (e) => {
    const card = e.currentTarget;
    const rect = card.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    const centerX = rect.width / 2;
    const centerY = rect.height / 2;
    const angleX = ((y - centerY) / centerY) * 10;
    const angleY = ((centerX - x) / centerX) * 10;

    card.style.transform = `perspective(1000px) rotateX(${angleX}deg) rotateY(${angleY}deg)`;
  };

  const handleMouseLeave = (e) => {
    e.currentTarget.style.transform = 'perspective(1000px) rotateX(0) rotateY(0)';
  };

  return (
    <Card
      style={{ 
        cursor: 'pointer', 
        margin: '10px',
        transition: 'transform 0.1s', 
        transformStyle: 'preserve-3d',
        boxShadow: "0 4px 4px 0 rgba(0, 0, 0, 0.1), 0 6px 10px 0 rgba(0, 0, 0, 0.1)"
      }}
      variant="outlined"
      onClick={() => onClick(movie)}
      onMouseMove={handleMouseMove}
      onMouseLeave={handleMouseLeave}
    >
      <CardContent>
        <Box mb={2}>
            <Typography variant="h5" style={{fontWeight: "bold"}}>{movie.PrimaryTitle}</Typography>
        </Box>
        <Box mb={2}>
            <Typography variant="body" 
            style={{fontWeight: "bold", color: getRatingColor(movie.AverageRating)}}>
                Rating: {movie.AverageRating}
            </Typography>
        </Box>
        <Chip
          key={movie.TitleType}
          label={movie.TitleType}
          variant="outlined"
          style={{
            borderColor: getTitleTypeColor(movie.TitleType),
            color: getTitleTypeColor(movie.TitleType),
            margin: '2px'
          }}
        />
      </CardContent>
    </Card>
  );
};

export default MovieCard;