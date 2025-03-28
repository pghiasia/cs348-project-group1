export const getTitleTypeColor = (titleType) => {
    const titleColors = {
        movie: '#3B82F6',
        series: '#9333EA',
        short: '#10B981',
        episode: '#F97316',
    };

    return titleColors[titleType.toLowerCase()] || '#6B7280'
}

export const getRatingColor = (rating) => {
    if (rating < 3.0) {
        return '#D4D4D4'
    }
    else if (rating >= 3.0 && rating < 5) {
        return '#A3A3A3'
    }
    else if (rating >= 5.0 && rating < 8) {
        return '#737373'
    }
    else {
        return '#525252'
    }
}