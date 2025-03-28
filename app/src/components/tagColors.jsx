export const getTitleTypeColor = (titleType) => {
    const titleColors = {
        movie: '#4169E1',
        series: '#8e1544',
        short: '#0fd835',
        episode: '#f0682b',
    };

    return titleColors[titleType.toLowerCase()] || '#7f7f7f'
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