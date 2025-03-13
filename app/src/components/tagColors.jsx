export const getGenreColor = (genre) => {
    const genreColors = {
        documentary: '#4169E1',
        short: '#8e1544',
        animation: '#0fd835',
        comedy: '#f0682b',
        romance: '#c236e3',
        sport: '#B87333',   
    };

    return genreColors[genre.toLowerCase()] || '#7f7f7f'
}

export const getRatingColor = (rating) => {
    if (rating < 3.0) {
        return '#899499'
    }
    else if (rating >= 3.0 && rating < 5) {
        return '#0047AB'
    }
    else if (rating >= 5.0 && rating < 8) {
        return '#008000'
    }
    else {
        return '#D70040'
    }
}