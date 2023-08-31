import { useEffect, useState } from "react";
import { getRandomLocation } from "./api/getRandomLocation";

function App() {
  const [location, setLocation] = useState(null);

  useEffect(() => {
    const fetchLocation = async () => {
      try {
        const location = await getRandomLocation();
        setLocation(location);
      } catch (error) {
        console.log(error);

    }}
    fetchLocation();
  }, []);

  return (
    <div>
      {location ? (
        <div>
          Current location: 
          <br />
          {location.city}, {location.country}
        </div>
      ) : (
        <div>Loading...</div>
      )}
    </div>
  );
}

export default App;