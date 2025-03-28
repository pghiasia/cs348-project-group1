import React, { useState } from 'react';
import { 
  MenuItem, 
  Select, 
  FormControl, 
  InputLabel, 
  Box, 
  TextField, 
  Button,
  Grid,
  Checkbox,
  FormControlLabel
} from '@mui/material';

const MovieFilters = ({ 
  onGenreChange, 
  onTitleTypeChange, 
  onCrewSearch,
  onTitleKeywordSearch,
  onYearRangeChange,
  onRatingRangeChange,
  onAdultContentChange,
  selectedGenre, 
  selectedTitleType,
  isAdultContentIncluded
}) => {
  const [crewMember, setCrewMember] = useState('');
  const [titleKeyword, setTitleKeyword] = useState('');
  const [startYear, setStartYear] = useState('');
  const [endYear, setEndYear] = useState('');
  const [lowRating, setLowRating] = useState('');
  const [highRating, setHighRating] = useState('');

  const genres = [
    'Action', 'Game-Show', 'War', 'Biography', 'Adult', 'Music', 
    'Short', 'Film-Noir', 'Family', 'Drama', 'Fantasy', 'Crime', 
    'Romance', 'Reality-TV', 'Documentary', 'Western', 'News', 
    'Comedy', 'Animation', 'Thriller', 'Talk-Show', 'Adventure', 
    'History', 'Musical', 'Sci-Fi', 'Mystery', 'Horror', 'Sport'
  ];

  const titleTypes = ['movie', 'short', 'episode', 'series'];

  const handleCrewSearch = () => {
    onCrewSearch(crewMember);
  };

  const handleTitleKeywordSearch = () => {
    onTitleKeywordSearch(titleKeyword);
  };

  const handleYearRangeSearch = () => {
    onYearRangeChange({
      startYear: startYear || null, 
      endYear: endYear || null
    });
  };

  const handleRatingRangeSearch = () => {
    onRatingRangeChange({
      lowRating: lowRating || null, 
      highRating: highRating || null
    });
  };

  return (
    <Box 
      sx={{ 
        mb: 2,
        p: 2,
        backgroundColor: '#f5f5f5',
        borderRadius: 2,
        boxShadow: '0 2px 4px rgba(0,0,0,0.1)'
      }}
    >
      <Grid container spacing={2}>
        {/* Existing Genre and Title Type Filters */}
        <Grid item xs={12} sm={6}>
          <FormControl fullWidth size="small">
            <InputLabel>Genres</InputLabel>
            <Select
              value={selectedGenre}
              label="Genres"
              onChange={(e) => onGenreChange(e.target.value)}
            >
              <MenuItem value="">
                <em>All Genres</em>
              </MenuItem>
              {genres.sort().map((genre) => (
                <MenuItem key={genre} value={genre}>
                  {genre}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>

        <Grid item xs={12} sm={6}>
          <FormControl fullWidth size="small">
            <InputLabel>Title Type</InputLabel>
            <Select
              value={selectedTitleType}
              label="Title Type"
              onChange={(e) => onTitleTypeChange(e.target.value)}
            >
              <MenuItem value="">
                <em>All Types</em>
              </MenuItem>
              {titleTypes.map((type) => (
                <MenuItem key={type} value={type}>
                  {type.charAt(0).toUpperCase() + type.slice(1)}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>

        {/* Crew and Title Keyword Search */}
        <Grid item xs={12}>
          <Grid container spacing={2}>
            <Grid item xs={6}>
              <Box sx={{ display: 'flex', gap: 2 }}>
                <TextField
                  fullWidth
                  size="small"
                  label="Search Crew Member"
                  variant="outlined"
                  value={crewMember}
                  onChange={(e) => setCrewMember(e.target.value)}
                  placeholder="Enter crew member name"
                />
                <Button 
                  variant="contained" 
                  onClick={handleCrewSearch}
                  sx={{ height: '40px' }}
                >
                  Search
                </Button>
              </Box>
            </Grid>
            <Grid item xs={6}>
              <Box sx={{ display: 'flex', gap: 2 }}>
                <TextField
                  fullWidth
                  size="small"
                  label="Search Title Keyword"
                  variant="outlined"
                  value={titleKeyword}
                  onChange={(e) => setTitleKeyword(e.target.value)}
                  placeholder="Enter title keyword"
                />
                <Button 
                  variant="contained" 
                  onClick={handleTitleKeywordSearch}
                  sx={{ height: '40px' }}
                >
                  Search
                </Button>
              </Box>
            </Grid>
          </Grid>
        </Grid>

        {/* Year Range Filter */}
        <Grid item xs={12}>
          <Grid container spacing={2}>
            <Grid item xs={6}>
              <Box sx={{ display: 'flex', gap: 2 }}>
                <TextField
                  fullWidth
                  size="small"
                  label="Start Year"
                  variant="outlined"
                  type="number"
                  value={startYear}
                  onChange={(e) => setStartYear(e.target.value)}
                  placeholder="Start Year"
                  InputProps={{
                    inputProps: { 
                      min: 1900, 
                      max: new Date().getFullYear() 
                    }
                  }}
                />
                <TextField
                  fullWidth
                  size="small"
                  label="End Year"
                  variant="outlined"
                  type="number"
                  value={endYear}
                  onChange={(e) => setEndYear(e.target.value)}
                  placeholder="End Year"
                  InputProps={{
                    inputProps: { 
                      min: 1900, 
                      max: new Date().getFullYear() 
                    }
                  }}
                />
                <Button 
                  variant="contained" 
                  onClick={handleYearRangeSearch}
                  sx={{ height: '40px' }}
                >
                  Filter
                </Button>
              </Box>
            </Grid>

            {/* Rating Range Filter */}
            <Grid item xs={6}>
              <Box sx={{ display: 'flex', gap: 2 }}>
                <TextField
                  fullWidth
                  size="small"
                  label="Low Rating"
                  variant="outlined"
                  type="number"
                  value={lowRating}
                  onChange={(e) => setLowRating(e.target.value)}
                  placeholder="Low Rating"
                  InputProps={{
                    inputProps: { 
                      min: 0, 
                      max: 10,
                      step: 0.1
                    }
                  }}
                />
                <TextField
                  fullWidth
                  size="small"
                  label="High Rating"
                  variant="outlined"
                  type="number"
                  value={highRating}
                  onChange={(e) => setHighRating(e.target.value)}
                  placeholder="High Rating"
                  InputProps={{
                    inputProps: { 
                      min: 0, 
                      max: 10,
                      step: 0.1
                    }
                  }}
                />
                <Button 
                  variant="contained" 
                  onClick={handleRatingRangeSearch}
                  sx={{ height: '40px' }}
                >
                  Filter
                </Button>
              </Box>
            </Grid>
          </Grid>
        </Grid>

        {/* Adult Content Filter */}
        <Grid item xs={12}>
          <FormControlLabel
            control={
              <Checkbox
                checked={isAdultContentIncluded}
                onChange={(e) => onAdultContentChange(e.target.checked)}
              />
            }
            label="Adult Content"
          />
        </Grid>
      </Grid>
    </Box>
  );
};

export default MovieFilters;