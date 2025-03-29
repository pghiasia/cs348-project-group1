import React, { useState } from 'react';
import { 
  Paper, 
  Grid, 
  Typography, 
  TextField, 
  Button, 
  Box 
} from '@mui/material';

const MovieActorFilter = ({ 
  onActorSearch, 
  onClearActorFilter, 
  isActorFilterActive, 
  actorName 
}) => {
  const [inputActorName, setInputActorName] = useState(actorName || '');

  const handleActorNameChange = (e) => {
    setInputActorName(e.target.value);
  };

  const handleSearch = () => {
    if (inputActorName.trim()) {
      onActorSearch(inputActorName);
    }
  };

  return (
    <Paper 
      elevation={0} 
      sx={{ 
        mb: 3, 
        p: 2, 
        backgroundColor: '#f5f5f5', 
        borderRadius: 2,
        boxShadow: '0 2px 4px rgba(0,0,0,0.1)'
      }}
    >
      <Grid container spacing={2} alignItems="center">
        <Grid item xs={12}>
          <Typography variant="h6" sx={{ mb: 1 }}>
            Top 20 Movie by Actor
          </Typography>
        </Grid>
        <Grid item xs={12} sm={6}>
          <TextField
            fullWidth
            size="small"
            label="Actor's Name"
            variant="outlined"
            value={inputActorName}
            onChange={handleActorNameChange}
            placeholder="Enter actor's name"
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <Box sx={{ display: 'flex', gap: 2 }}>
            <Button 
              variant="contained" 
              color="primary"
              onClick={handleSearch}
              disabled={!inputActorName.trim()}
              sx={{ height: '40px' }}
            >
              Search Movies with Actor
            </Button>
            {isActorFilterActive && (
              <Button 
                variant="outlined" 
                color="secondary"
                onClick={onClearActorFilter}
                sx={{ height: '40px' }}
              >
                Clear Actor Filter
              </Button>
            )}
          </Box>
        </Grid>
      </Grid>
      {isActorFilterActive && (
        <Box sx={{ mt: 2 }}>
          <Typography variant="body2" color="primary">
            Currently showing movies with actor: <strong>{actorName}</strong>
          </Typography>
        </Box>
      )}
    </Paper>
  );
};

export default MovieActorFilter;