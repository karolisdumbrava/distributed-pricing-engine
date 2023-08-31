import axios from "axios";

export const getRandomLocation = async () => {
    try {
        const response = await axios.get(
            "http://localhost:8081/api/random-location"
        );
        return response.data;
    } catch (error) {
        console.log(error);
    }
};