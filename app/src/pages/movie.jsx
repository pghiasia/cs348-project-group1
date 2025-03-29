import React, { useState, useEffect } from 'react';
import { 
  Container, 
  Grid, 
  CircularProgress, 
  Typography, 
  Box 
} from '@mui/material';
import MovieCard from '../components/movieCard';
import MovieModal from '../components/movieModal';
import MovieFilters from '../components/movieFilter';
import MovieActorFilter from '../components/movieActorFilter';

const Movie = () => { 
  //base states
  const [movies, setMovies] = useState([]);
  const [selectedMovie, setSelectedMovie] = useState(null);
  const [open, setOpen] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  
  //filter states
  const [selectedGenre, setSelectedGenre] = useState('');
  const [selectedTitleType, setSelectedTitleType] = useState('');
  const [crewMember, setCrewMember] = useState('');
  const [titleKeyword, setTitleKeyword] = useState('');
  const [startYear, setStartYear] = useState(null);
  const [endYear, setEndYear] = useState(null);
  const [lowRating, setLowRating] = useState(null);
  const [highRating, setHighRating] = useState(null);
  const [isAdultContentIncluded, setIsAdultContentIncluded] = useState(false);
  
  //actor filter states
  const [actorName, setActorName] = useState('');
  const [isActorFilterActive, setIsActorFilterActive] = useState(false);

  const fetchMovies = async () => {
    try {
      setIsLoading(true);
      setError(null);

      const params = new URLSearchParams();
      
      if (selectedGenre) params.append('genre', selectedGenre);
      if (selectedTitleType) params.append('titleType', selectedTitleType);
      if (crewMember) params.append('crewMember', crewMember);
      if (titleKeyword) params.append('titleKeyword', titleKeyword);      
      if (startYear) params.append('startYear', startYear);
      if (endYear) params.append('endYear', endYear);
      if (lowRating) params.append('lowRating', lowRating);
      if (highRating) params.append('highRating', highRating);
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
      setError('Failed to load movies. Please try again.');
    }
    finally {
      setIsLoading(false);
    }
  }

  const fetchMoviesByActor = async (actorNameToSearch) => {
    if (!actorNameToSearch.trim()) return;
    
    try {
      setIsLoading(true);
      setError(null);
      
      const actorResponse = await fetch(`http://localhost:9888/highestRatingWithActor?actor=${actorNameToSearch.trim()}`, {
        method: "GET",
        headers: {
          'Content-Type': 'application/json',
        }
      });
      
      if (!actorResponse.ok) {
        throw new Error('Fetch movies by actor failed');
      }
      
      const actorData = await actorResponse.json();
      setMovies(actorData);
    }
    catch (error) {
      console.error(error);
      setError('Failed to load movies. Please try again.');
    }
    finally {
      setIsLoading(false);
    }
  }

  useEffect(() => {
    if (!isActorFilterActive) {
      fetchMovies();
    }
  }, [
    selectedGenre, 
    selectedTitleType, 
    crewMember, 
    titleKeyword, 
    startYear, 
    endYear, 
    lowRating, 
    highRating,
    isAdultContentIncluded,
    isActorFilterActive
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
    setIsActorFilterActive(false);
  };

  const handleTitleTypeChange = (titleType) => {
    setSelectedTitleType(titleType);
    setIsActorFilterActive(false);
  };

  const handleCrewSearch = (crew) => {
    setCrewMember(crew);
    setIsActorFilterActive(false);
  };

  const handleTitleKeywordSearch = (keyword) => {
    setTitleKeyword(keyword);
    setIsActorFilterActive(false);
  };

  const handleYearRangeChange = ({ startYear, endYear }) => {
    setStartYear(startYear);
    setEndYear(endYear);
    setIsActorFilterActive(false);
  };

  const handleRatingRangeChange = ({ lowRating, highRating }) => {
    setLowRating(lowRating);
    setHighRating(highRating);
    setIsActorFilterActive(false);
  };

  const handleAdultContentChange = (isIncluded) => {
    setIsAdultContentIncluded(isIncluded);
    setIsActorFilterActive(false);
  };

  const handleActorSearch = (name) => {
    fetchMoviesByActor(name);
    setActorName(name);
    setIsActorFilterActive(true);
  };

  const clearActorFilter = () => {
    setActorName('');
    setIsActorFilterActive(false);
  };

  const renderContent = () => {
    if (isLoading) {
      return (
        <Box 
          display="flex" 
          justifyContent="center" 
          alignItems="center" 
          height="60vh"
        >
          <CircularProgress />
        </Box>
      );
    }

    if (movies.length === 0) {
      return (
        <Box 
          display="flex" 
          justifyContent="center" 
          alignItems="center" 
          height="60vh"
        >
          <Typography variant="h6" color="textSecondary">
            No movies found
          </Typography>
        </Box>
      );
    }

    return (
      <Grid container spacing={2}>
        {movies.map(movie => (
          <Grid item xs={12} sm={6} md={4} key={movie.Mid}>
            <MovieCard movie={movie} onClick={handleOpen} />
          </Grid>
        ))}
      </Grid>
    );
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

      <MovieActorFilter 
        onActorSearch={handleActorSearch}
        onClearActorFilter={clearActorFilter}
        isActorFilterActive={isActorFilterActive}
        actorName={actorName}
      />

      {renderContent()}

      {selectedMovie && (
        <MovieModal open={open} handleClose={handleClose} movie={selectedMovie} />
      )}
    </Container>
  );
}

export default Movie;