import { css } from 'styled-components'

const CardStyle = css`
    background-color: #f0f0f0;
    border: 1px solid #ddd;
    border-radius: 5%;
    padding: 10px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);

    /* Custom styles for the WeatherCard */
    width: 220px;
    height: 210px;

    /* Center content within the card */
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    /* If it's the only card in the grid, position it in the upper left corner */
    &:only-child {
        grid-column: 1;
        grid-row: 1;
    }
`

export default CardStyle