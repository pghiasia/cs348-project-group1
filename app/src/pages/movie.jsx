import React, { useState, useEffect } from 'react';
import { Container, Grid } from '@mui/material';
import MovieCard from '../components/movieCard';
import MovieModal from '../components/movieModal';
import MovieFilters from '../components/movieFilter';

const Movie = () => { 
  const [movies, setMovies] = useState([]);
  const [selectedMovie, setSelectedMovie] = useState(null);
  const [open, setOpen] = useState(false);
  
  // New state for dropdowns and search
  const [selectedGenre, setSelectedGenre] = useState('');
  const [selectedTitleType, setSelectedTitleType] = useState('');
  const [crewMember, setCrewMember] = useState('');
  const [titleKeyword, setTitleKeyword] = useState('');
  const [startYear, setStartYear] = useState(null);
  const [endYear, setEndYear] = useState(null);
  const [lowRating, setLowRating] = useState(null);
  const [highRating, setHighRating] = useState(null);
  const [isAdultContentIncluded, setIsAdultContentIncluded] = useState(false);

  const fetchMovies = async () => {
    try {
      const params = new URLSearchParams();
      
      if (selectedGenre) params.append('genre', selectedGenre);
      if (selectedTitleType) params.append('titleType', selectedTitleType);
      if (crewMember) params.append('crewMember', crewMember);
      if (titleKeyword) params.append('titleKeyword', titleKeyword);
      
      // Add year range parameters
      if (startYear) params.append('startYear', startYear);
      if (endYear) params.append('endYear', endYear);
      
      // Add rating range parameters
      if (lowRating) params.append('lowRating', lowRating);
      if (highRating) params.append('highRating', highRating);
      
      // Add adult content filter
      params.append('isAdult', isAdultContentIncluded);
      
      const response = await fetch(`http://localhost:9888/movies?${params.toString()}`, {
        method: "GET",
        headers: {
          'Content-Type': 'application/json',
        }
      });
    
      if (!response.ok) {
        throw new Error('Fetch movie failed');
      }
    
      const data = await response.json();
      setMovies(data);
    }
    catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    fetchMovies();
  }, [
    selectedGenre, 
    selectedTitleType, 
    crewMember, 
    titleKeyword, 
    startYear, 
    endYear, 
    lowRating, 
    highRating,
    isAdultContentIncluded
  ]);

  const handleOpen = (movie) => {
    setSelectedMovie(movie);
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
    setSelectedMovie(null);
  };

  const handleGenreChange = (genre) => {
    setSelectedGenre(genre);
  };

  const handleTitleTypeChange = (titleType) => {
    setSelectedTitleType(titleType);
  };

  const handleCrewSearch = (crew) => {
    setCrewMember(crew);
  };

  const handleTitleKeywordSearch = (keyword) => {
    setTitleKeyword(keyword);
  };

  const handleYearRangeChange = ({ startYear, endYear }) => {
    setStartYear(startYear);
    setEndYear(endYear);
  };

  const handleRatingRangeChange = ({ lowRating, highRating }) => {
    setLowRating(lowRating);
    setHighRating(highRating);
  };

  const handleAdultContentChange = (isIncluded) => {
    setIsAdultContentIncluded(isIncluded);
  };

  return (
    <Container>
      <MovieFilters 
        onGenreChange={handleGenreChange}
        onTitleTypeChange={handleTitleTypeChange}
        onCrewSearch={handleCrewSearch}
        onTitleKeywordSearch={handleTitleKeywordSearch}
        onYearRangeChange={handleYearRangeChange}
        onRatingRangeChange={handleRatingRangeChange}
        onAdultContentChange={handleAdultContentChange}
        selectedGenre={selectedGenre}
        selectedTitleType={selectedTitleType}
        isAdultContentIncluded={isAdultContentIncluded}
      />

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