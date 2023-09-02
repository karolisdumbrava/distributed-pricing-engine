import axios from "axios";

export const getProductsPricesByLocation = async (locationId) => {
    try {
        const response = await axios.get(
            `http://localhost:8080/api/products-prices-by-location?location_id=${locationId}`
        );
        return response.data;
    } catch (error) {
        console.log(error);
    }
};